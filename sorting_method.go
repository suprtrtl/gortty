package main

type SortingSteps [][]int

type SortingMethod interface {
	// Takes in array and returns a list of all total arrays ever modified from sorting,
	// terrible space efficiency and should be replaced however better solution unclear
	GetSortingSteps([]int) SortingSteps
}

type BubbleSort struct{}

func (bs BubbleSort) GetSortingSteps(data []int) SortingSteps {
	n := len(data)
	// Also inefficient should get a size and capacity beforehand
	steps := SortingSteps{data}

	for i := range n {
		swapped := false
		for j := range n - i - 1 {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				swapped = true

				// Add array to steps list
				tmp := make([]int, len(data))
				copy(tmp, data)
				steps = append(steps, tmp)
			}
		}

		if (swapped == false) {
			break
		}
	}

	return steps
}
