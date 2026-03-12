package main

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
