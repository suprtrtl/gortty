package main

import (
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	data       []int
	method     SortingMethod
	graph      ArrayGraph
	dims       Dimension
	delay      int
	programPtr *tea.Program
}

type programPtrMsg struct{ programPtr *tea.Program }
type startSort struct{}

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

	case startSort:
		go m.method.sort(m)
		return m, nil

	case programPtrMsg:
		m.programPtr = msg.programPtr
		return m, nil

	case sortStep:
		if msg.sorted {
			return m, tea.Quit
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

	go func() { // temporary setup, setup program ptr then start sorting after 2 secs
		time.Sleep(time.Second * 1)
		p.Send(programPtrMsg{p})
		time.Sleep(time.Second * 1)
		p.Send(startSort{})
	}()

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
