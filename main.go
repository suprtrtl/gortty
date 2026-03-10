package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	data    []int
	queue   *SortingQueue
	method  SortingMethod
	graph   ArrayGraph
	dims    Dimension
	delay   int
	program *tea.Program
}

type ProgramPtrMsg *tea.Program
type StartSortMsg struct{}

func InitialModel() model {
	sq := NewSortingQueue()
	return model{
		data: []int{432, 17, 298, 401, 56, 73, 489, 120, 345, 210, 67, 154, 399, 278, 44, 311, 92, 407, 188, 265,
			134, 358, 21, 476, 303, 84, 250, 168, 392, 59, 147, 326, 415, 203, 12, 437, 290, 361, 75, 224},

		queue: &sq,
		method: nil,
		graph: BarGraph{component: "▐▌"},
		// graph: BarGraph{component: "▊"},
		dims:  Dimension{width: 0, height: 0, spacing: 2},
		delay: 25,
	}
}

func (m model) RandomizeData() {
	rand.Shuffle(len(m.data), func(i, j int) {
		m.data[i], m.data[j] = m.data[j], m.data[i]
	})
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
		method := m.queue.Next()
		m.method = method
		m.RandomizeData()
		time.Sleep(time.Millisecond * time.Duration(m.delay))
		go m.method.Sort(m)
		return m, tea.RequestWindowSize

	case RenderStepMsg:
		if msg { // future handling for when the algorithm completes sorting
			return m, func() tea.Msg {
				return StartSortMsg{}
			} // Sort again
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

	s := graph + "\n";

	switch m.method {
	case BubbleSort{}:
		s += "bubble sort";
	case SelectionSort{}:
		s += "selection sort";
	case MergeSort{}:
		s += "merge sort";
	}

	view := tea.NewView(s)

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
