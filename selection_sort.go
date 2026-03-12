package main

import "time"

type SelectionSort struct{}

func (ss SelectionSort) Sort(m model, useWeights bool) {
	delay := m.delay
	if useWeights {
		delay = int(float32(delay) * 0.10)
	}

	for itr := 0; itr < len(m.data)-1; itr++ {
		minIndex := itr
		for index := itr + 1; index < len(m.data); index++ {
			if m.data[minIndex] > m.data[index] {
				minIndex = index
			}
			m.program.Send(RenderStepMsg{false, highlightMap{
				minIndex: {},
				index:    {},
			}})
			time.Sleep(time.Millisecond * time.Duration(delay))
		}

		m.data[itr], m.data[minIndex] = m.data[minIndex], m.data[itr]
	}

	m.program.Send(RenderStepMsg{true, highlightMap{}})
}
