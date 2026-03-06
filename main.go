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
		data:   []int{6, 2, 7, 1, 4, 8, 3, 5, 6, 2, 7, 1, 4, 8, 3, 5, 6, 2, 7, 1, 4, 8, 3, 5},
		method: BubbleSort{},
		graph:  BarGraph{component: "▐▌"},
		dims:   Dimension{width: 48, height: 30},
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
		return m, nil

	case RenderStepMsg:
		if msg { // future handling for when the algorithm completes sorting
			return m, tea.Quit // for now, we just quit
		}
		return m, nil
	}

	return m, nil
}

func (m model) View() tea.View {
	return tea.NewView(m.graph.Render(m.data, m.dims))
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
