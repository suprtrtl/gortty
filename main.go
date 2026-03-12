package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
)

func GenerateSteppedArray(n uint) []int {
	data := make([]int, n)

	for i := range data {
		data[i] = i
	}

	return data
}

type model struct {
	data        []int
	queue       *SortingQueue
	method      SortingMethod
	graph       ArrayGraph
	dims        Dimension
	delay       int
	highlighted highlightMap // map[int]struct{} where keys are used for O(1) lookup of highlighted indicies
	program     *tea.Program
}

type ProgramPtrMsg *tea.Program
type StartSortMsg struct{}

func InitialModel() model {
	sq := NewSortingQueue()
	return model{
		data:   GenerateSteppedArray(4),
		queue:  &sq,
		method: nil,
		graph:  BarGraph{component: "▉"},
		// graph: BarGraph{component: "▊"},
		dims:  Dimension{width: 0, height: 0, spacing: 2},
		delay: 50,
	}
}

func (m *model) SetData() {
	m.DataResize()
	m.RandomizeData()
}

func (m *model) DataResize() {
	numChars := m.dims.height - 2*m.dims.spacing
	width := float64(m.dims.width - 2*m.dims.spacing)

	switch graphType := m.graph.(type) {
	case BarGraph:
		width /= float64(graphType.GetComponentSize())
	}

	multiplier := math.Floor(width / float64(numChars))
	m.data = GenerateSteppedArray(uint(numChars * int(multiplier)))
}

func (m *model) RandomizeData() {
	rand.Shuffle(len(m.data), func(i, j int) {
		m.data[i], m.data[j] = m.data[j], m.data[i]
	})
}

// TODO: We can probably use this instead of the go routine running p.send() after program init
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "=", "+":
			// no upper bound for now
			m.delay += 5
			return m, nil

		case "-", "_":
			if m.delay > 5 {
				m.delay -= 5
				return m, nil
			}
		}

	case ProgramPtrMsg:
		m.program = msg
		return m, nil

	case StartSortMsg:
		method := m.queue.Next()
		m.method = method
		m.SetData()
		time.Sleep(time.Millisecond * time.Duration(m.delay))
		go m.method.Sort(m, true) // by default i'm enabling weights until configs are done
		return m, tea.RequestWindowSize

	case RenderStepMsg:
		m.highlighted = msg.highlighted
		if msg.isSorted { // future handling for when the algorithm completes sorting
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
	graph := m.graph.Render(m.data, m.dims, m.highlighted)

	s := graph + "\n"

	switch m.method.(type) {
	case BubbleSort:
		s += "bubble sort"
	case SelectionSort:
		s += "selection sort"
	case MergeSort:
		s += "merge sort"
	case CombSort:
		s += "comb sort"
	case QuickSort:
		s += "quick sort"
	}

	s += fmt.Sprintf(" | delay %d", m.delay)

	view := tea.NewView(s)

	view.AltScreen = true

	return view
}

func main() {
	p := tea.NewProgram(InitialModel())
	go func() {
		p.Send(ProgramPtrMsg(p)) // TODO: find a way to get p from within into and get rid of this!
		p.Send(StartSortMsg{})
	}()

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
