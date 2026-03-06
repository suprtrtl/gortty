package main

import (
	"time"
)

type RenderStepMsg struct {
	IsSorted bool
}

type SortingMethod interface {
	Sort(model)
}

type BubbleSort struct{}

func (bs BubbleSort) Sort(m model) {
	for itr := range m.data {
		sorted := true
		for index := 0; index < len(m.data)-1-itr; index++ {
			if m.data[index] > m.data[index+1] {
				sorted = false
				m.data[index], m.data[index+1] = m.data[index+1], m.data[index]
				m.program.Send(RenderStepMsg{false})
				time.Sleep(time.Millisecond * time.Duration(m.delay))
			}
		}

		if sorted {
			m.program.Send(RenderStepMsg{true})
			return
		}
	}
}
