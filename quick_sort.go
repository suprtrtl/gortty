package main

import (
	"math/rand"
	"time"
)

type QuickSort struct{}

func (qs QuickSort) Sort(m model, _ bool) {

	qs.quickSort(m.data, m, 0)

	time.Sleep(time.Millisecond * time.Duration(m.delay))
	m.program.Send(RenderStepMsg{true, highlightMap{}})
}

func (qs QuickSort) quickSort(data []int, m model, offset int) {
	if len(data) <= 1 {
		return
	}

	pivotIndex := qs.partition(data, m, offset)

	qs.quickSort(data[:pivotIndex+1], m, offset)
	qs.quickSort(data[pivotIndex+1:], m, offset+pivotIndex+1)
}

func (qs QuickSort) partition(data []int, m model, offset int) int {
	leftPtr := -1
	rightPtr := len(data)

	pivot := data[rand.Intn(rightPtr)]

	for { // main loop
		for { // find left element larger than pivot
			leftPtr++
			if data[leftPtr] >= pivot {
				break
			}
		}

		for { // find right element smaller than pivot
			rightPtr--
			if data[rightPtr] <= pivot {
				break
			}
		}

		if leftPtr >= rightPtr { // break loop on cross
			break
		}

		data[leftPtr], data[rightPtr] = data[rightPtr], data[leftPtr]
		m.program.Send(RenderStepMsg{false, highlightMap{
			leftPtr + offset:  {},
			rightPtr + offset: {},
		}})
		time.Sleep(time.Millisecond * time.Duration(m.delay))
	}

	return rightPtr
}
