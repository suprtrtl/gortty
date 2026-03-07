package main

import (
	"math/rand"
	"time"
)

// Holds Whether or not array is sorted
type RenderStepMsg bool

type SortingMethod interface {
	Sort(model)
}

type BogoSort struct{}
type BubbleSort struct{}
type QuickSort struct{}
type SelectionSort struct{}

func (bg BogoSort) Sort(m model) {
	for {
		sorted := true
		for index := 0; index < len(m.data)-1; index++ {
			if m.data[index] > m.data[index+1] {
				sorted = false
			}
		}

		if sorted {
			m.program.Send(RenderStepMsg(true))
			return
		} else {
			rand.Shuffle(len(m.data), func(i, j int) {
				m.data[i], m.data[j] = m.data[j], m.data[i]
			})
			m.program.Send(RenderStepMsg(false))
		}

		time.Sleep(time.Millisecond * time.Duration(m.delay))
	}
}

func (bs BubbleSort) Sort(m model) {
	for itr := range m.data {
		sorted := true
		for index := 0; index < len(m.data)-1-itr; index++ {
			if m.data[index] > m.data[index+1] {
				sorted = false
				m.data[index], m.data[index+1] = m.data[index+1], m.data[index]

				m.program.Send(RenderStepMsg(false))
				time.Sleep(time.Millisecond * time.Duration(m.delay))
			}
		}

		if sorted {
			m.program.Send(RenderStepMsg(true))
			return
		}
	}
}

func (ss SelectionSort) Sort(m model) {
	for itr := 0; itr < len(m.data)-1; itr++ {
		minIndex := itr
		for index := itr + 1; index < len(m.data); index++ {
			if m.data[minIndex] > m.data[index] {
				minIndex = index
			}
		}

		m.data[itr], m.data[minIndex] = m.data[minIndex], m.data[itr]
		m.program.Send(RenderStepMsg(false))
		time.Sleep(time.Millisecond * time.Duration(m.delay))
	}

	m.program.Send(RenderStepMsg(true))
}
