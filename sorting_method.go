package main

import (
	"time"
)

type sortStep struct {
	sorted bool
}

type SortingMethod interface {
	sort(model)
}

type BubbleSort struct{}

func (bs BubbleSort) sort(m model) {
	for itr := range m.data {
		sorted := true
		for index := 0; index < len(m.data)-1-itr; index++ {
			if m.data[index] > m.data[index+1] {
				sorted = false
				m.data[index], m.data[index+1] = m.data[index+1], m.data[index]
				m.programPtr.Send(sortStep{false})
				time.Sleep(time.Millisecond * time.Duration(m.delay))
			}
		}

		if sorted {
			m.programPtr.Send(sortStep{true})
			return
		}
	}
}
