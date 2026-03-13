package main

import (
	"time"
)

type ShellSort struct{}

func (ss ShellSort) Sort(m model, _ bool) {
	n := len(m.data)

	for gap := n / 2; gap > 0; gap /= 2 {

		for i := gap; i < n; i++ {
			temp := m.data[i]
			j := i

			for j >= gap && m.data[j-gap] > temp {
				m.data[j] = m.data[j-gap]
				j -= gap

				m.program.Send(RenderStepMsg{false, highlightMap{
					i: {},
					j: {},
				}})
				time.Sleep(time.Millisecond * time.Duration(m.delay))

			}

			m.data[j] = temp

		}
	}

	m.program.Send(RenderStepMsg{true, highlightMap{}})
	time.Sleep(time.Millisecond * time.Duration(m.delay))
}
