package main

import (
// tea "charm.land/bubbletea/v2"
)

type model struct {
	data  []int
	graph ArrayGraph
}

func main() {
	bg := BarGraph{component: "▐▌"}
	dims := Dimension{
		width:  48,
		height: 30,
	}

	data := []int{6, 2, 7, 1, 4, 8, 3, 5, 6, 2, 7, 1, 4, 8, 3, 5, 6, 2, 7, 1, 4, 8, 3, 5}

	bubble_sort(bg, data, dims)
}
