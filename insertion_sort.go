package main

import (
	"time"
)

type InsertionSort struct{}

func (is InsertionSort) Sort(m model, _ bool) {
	n := len(m.data)

	for i := 1; i < n; i++ {
		key := m.data[i]
		j := i - 1

		for j >= 0 && m.data[j] > key {
			m.data[j+1] = m.data[j]
			j = j - 1

			m.program.Send(RenderStepMsg{false, highlightMap{
				i: {},
				j: {},
			}})
			time.Sleep(time.Millisecond * time.Duration(m.delay))
		}

		m.data[j+1] = key
	}

	time.Sleep(time.Millisecond * time.Duration(m.delay))
	m.program.Send(RenderStepMsg{true, highlightMap{}})
}
