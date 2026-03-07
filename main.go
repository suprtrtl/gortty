package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	data    []int
	method  SortingMethod
	graph   ArrayGraph
	dims    Dimension
	delay   int
	program *tea.Program
}

type ProgramPtrMsg *tea.Program
type StartSortMsg struct{}

func InitialModel() model {
	return model{
		data: []int{432, 17, 298, 401, 56, 73, 489, 120, 345, 210, 67, 154, 399, 278, 44, 311, 92, 407, 188, 265,
			134, 358, 21, 476, 303, 84, 250, 168, 392, 59, 147, 326, 415, 203, 12, 437, 290, 361, 75, 224,
			318, 49, 406, 177, 268, 95, 350, 421, 63, 284, 139, 472, 33, 199, 344, 118, 257, 381, 6, 463,
			170, 292, 54, 409, 146, 227, 368, 88, 315, 240, 402, 71, 183, 299, 52, 434, 266, 157, 321, 90,
			413, 205, 47, 372, 281, 104, 498, 36, 219, 356, 142, 307, 79, 451, 234, 165, 389, 97, 260, 341,
			58, 426, 173, 294, 111, 360, 24, 475, 198, 327, 62, 410, 285, 149, 374, 8, 439, 216, 301, 94,
			352, 131, 468, 41, 255, 383, 76, 420, 187, 309, 15, 447, 233, 169, 396, 102, 274, 337, 64, 458,
			192, 314, 83, 371, 145, 222, 404, 53, 287, 178, 349, 121, 466, 29, 207, 332, 98, 417, 270, 156,
			388, 45, 343, 116, 259, 362, 70, 448, 184, 296, 135, 379, 22, 461, 243, 174, 395, 109, 264, 328,
			50, 412, 286, 167, 354, 81, 430, 190, 305, 14, 442, 238, 176, 397, 101, 273, 336, 65, 459, 193},
		method: MergeSort{},
		graph:  BarGraph{component: "▊"},
		dims:   Dimension{width: 0, height: 0, spacing: 2},
		delay:  50,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case ProgramPtrMsg:
		m.program = msg
		return m, nil

	case StartSortMsg:
		go m.method.Sort(m)
		return m, tea.RequestWindowSize

	case RenderStepMsg:
		if msg { // future handling for when the algorithm completes sorting
			return m, tea.Quit // for now, we just quit
		}
		return m, tea.RequestWindowSize

	case tea.WindowSizeMsg:

		m.dims = Dimension{
			width:   msg.Width,
			height:  msg.Height,
			spacing: m.dims.spacing,
		}

		return m, nil
	}

	return m, nil
}

func (m model) View() tea.View {
	graph := m.graph.Render(m.data, m.dims)
	view := tea.NewView(graph)

	view.AltScreen = true

	return view
}

func main() {
	p := tea.NewProgram(InitialModel())
	go func() {
		p.Send(ProgramPtrMsg(p))
		p.Send(StartSortMsg{})
	}()

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
