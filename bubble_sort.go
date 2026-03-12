package main

import "time"

type BubbleSort struct{}

func (bs BubbleSort) Sort(m model, useWeights bool) {
	delay := m.delay
	if useWeights {
		delay = int(float32(delay) * 0.10)
	}

	for itr := range m.data {
		sorted := true
		for index := 0; index < len(m.data)-1-itr; index++ {
			if m.data[index] > m.data[index+1] {
				sorted = false
				m.data[index], m.data[index+1] = m.data[index+1], m.data[index]

				m.program.Send(RenderStepMsg{false, highlightMap{index: {}}})
				time.Sleep(time.Millisecond * time.Duration(delay))
			}
		}

		if sorted {
			m.program.Send(RenderStepMsg{true, highlightMap{}})
			return
		}
	}
}
