package main

import (
	"math"
	"math/rand"
	// "slices"
	"time"
)

// Render instructions

// use keys as a list of elements to highlight (with O(1) lookup for faster rendering)
type highlightMap map[int]struct{}
type RenderStepMsg struct {
	isSorted    bool
	highlighted highlightMap
}

type SortingMethod interface {
	// model - bubble tea model
	// bool (useWeights) modifies delay for less efficient algorithms so they run faster
	Sort(model, bool)
}

func IsSorted(data []int) bool {
	sorted := true
	for index := 0; index < len(data)-1; index++ {
		if data[index] > data[index+1] {
			sorted = false
			break
		}
	}
	return sorted
}

type BogoSort struct{}
type BubbleSort struct{}
type QuickSort struct{}
type SelectionSort struct{}
type MergeSort struct{}
type CombSort struct{}

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
			m.program.Send(RenderStepMsg{false, map[int]struct{}{}})
		}

		time.Sleep(time.Millisecond * time.Duration(delay))
	}
}

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

func (qs QuickSort) Sort(m model, useWeights bool) {
	delay := m.delay
	if useWeights {
		delay = int(float32(m.delay) * 1)
	}

	qs.quickSort(m.data, m, delay, 0)

	m.program.Send(RenderStepMsg{true, highlightMap{}})
}

func (qs QuickSort) quickSort(data []int, m model, delay int, offset int) {
	if len(data) <= 1 {
		return
	}

	pivotIndex := qs.partition(data, m, delay, offset)

	qs.quickSort(data[:pivotIndex+1], m, delay, offset)
	qs.quickSort(data[pivotIndex+1:], m, delay, offset+pivotIndex+1)
}

func (qs QuickSort) partition(data []int, m model, delay int, offset int) int {
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
		time.Sleep(time.Millisecond * time.Duration(delay))
	}

	return rightPtr
}

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

	model.program.Send(RenderStepMsg{false, highlightMap{}}) // TODO: @suprtrtl, i dunno what to highlight here. from: andrei
	time.Sleep(time.Millisecond * time.Duration(model.delay))
}

func (ms MergeSort) mergeSort(model model, data []int, l int, r int) {
	if l < r {
		m := l + (r-l)/2

		ms.mergeSort(model, data, l, m)
		ms.mergeSort(model, data, m+1, r)

		ms.merge(model, data, l, m, r)

	}
}

func (ms MergeSort) Sort(m model, _ bool) {
	ms.mergeSort(m, m.data, 0, len(m.data)-1)
	m.program.Send(RenderStepMsg{true, highlightMap{}})
	time.Sleep(time.Millisecond * time.Duration(m.delay))
}

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

	m.program.Send(RenderStepMsg{true, highlightMap{}})
	time.Sleep(time.Millisecond * time.Duration(m.delay))
}
