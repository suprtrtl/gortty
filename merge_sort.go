package main

import "time"

type MergeSort struct{}

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
