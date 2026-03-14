package main

import (
	"math"
	"time"
)

type CombSort struct{}

func (cs CombSort) Sort(m model, _ bool) {
	shrink := 1.3
	gap := len(m.data)
	sorted := false

	for !sorted {
		gap = int(math.Floor(float64(gap) / shrink))
		if gap <= 1 {
			sorted = true
			gap = 1
		}

		for i := range len(m.data) - gap {
			next := gap + i
			if m.data[i] > m.data[next] {
				m.data[i], m.data[next] = m.data[next], m.data[i]
				sorted = false
				m.program.Send(RenderStepMsg{false, highlightMap{
					i: {},
				}})
				time.Sleep(time.Millisecond * time.Duration(m.delay))
			}
		}
	}

	time.Sleep(time.Millisecond * time.Duration(m.delay))
	m.program.Send(RenderStepMsg{true, highlightMap{}})
}
