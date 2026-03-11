package main

import (
	"math/rand"
)

// Holds sorting Methods
type SortingQueue struct {
	// A list of all sorting methods you want to use
	methods []SortingMethod
	index   int
}

func NewSortingQueue() SortingQueue {
	sq := SortingQueue{
		[]SortingMethod{
			// BogoSort{}, - Removing for testing purposes
			BubbleSort{},
			SelectionSort{},
			MergeSort{},
			CombSort{},
		},
		0,
	}

	sq.randomize()

	return sq
}

func (sq *SortingQueue) randomize() {
    rand.Shuffle(len(sq.methods), func(i, j int) {
        sq.methods[i], sq.methods[j] = sq.methods[j], sq.methods[i]
    })
}

func (sq *SortingQueue) Next() SortingMethod {
    if sq.index >= len(sq.methods) {
        sq.index = 0
        sq.randomize()
    }
    method := sq.methods[sq.index]
    sq.index++
    return method
}

