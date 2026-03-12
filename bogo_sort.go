package main

import (
	"math/rand"
	"time"
)

type BogoSort struct{}

func (bg BogoSort) Sort(m model, useWeights bool) {
	delay := m.delay
	if useWeights {
		delay = int(float32(delay) * 0.25)
	}

	for {

		sorted := IsSorted(m.data)

		if sorted {
			m.program.Send(RenderStepMsg{true, highlightMap{}})
			return
		} else {
			rand.Shuffle(len(m.data), func(i, j int) {
				m.data[i], m.data[j] = m.data[j], m.data[i]
			})
			m.program.Send(RenderStepMsg{false, highlightMap{}})
		}

		time.Sleep(time.Millisecond * time.Duration(delay))
	}
}
