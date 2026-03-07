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
type MergeSort struct{}

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
				time.Sleep(time.Millisecond * time.Duration(m.delay))
			}
		}

		m.data[itr], m.data[minIndex] = m.data[minIndex], m.data[itr]
		m.program.Send(RenderStepMsg(false))
	}

	m.program.Send(RenderStepMsg(true))
}

// Merge sort methods

func (ms MergeSort) merge(model model, data []int, l int, m int, r int) {
	// Size of 2 subarrays
	n1 := m - l + 1
	n2 := r - m

	lArr := make([]int, n1)
	rArr := make([]int, n2)

	for i := range n1 {
		lArr[i] = data[l+i]
	}
	for j := range n2 {
		rArr[j] = data[m+1+j]
	}

	// Merge

	i, j := 0, 0

	k := l

	for i < n1 && j < n2 {
		if lArr[i] <= rArr[j] {
			data[k] = lArr[i]
			i++
		} else {
			data[k] = rArr[j]
			j++
		}
		k++

		model.program.Send(RenderStepMsg(false))
		time.Sleep(time.Millisecond * time.Duration(model.delay))
	}

	for i < n1 {
		data[k] = lArr[i]
		i++
		k++
	}

	for j < n2 {
		data[k] = rArr[j]
		j++
		k++
	}

}

func (ms MergeSort) mergeSort(model model, data []int, l int, r int) {
	if l < r {
		m := l + (r-l)/2

		ms.mergeSort(model, data, l, m)
		ms.mergeSort(model, data, m+1, r)

		ms.merge(model, data, l, m, r)

	}
}

func (ms MergeSort) Sort(m model) {
	ms.mergeSort(m, m.data, 0, len(m.data)-1)
	m.program.Send(RenderStepMsg(true))
	m.program.Send(RenderStepMsg(false))
	time.Sleep(time.Millisecond * time.Duration(m.delay))
}
