package main

import "fmt"

func bubble_sort(ag ArrayGraph, data []int, dims Dimension) {
	for range data {
		sorted := true
		for index := 0; index < len(data)-1; index++ {
			if data[index] > data[index+1] {
				sorted = false

				tmp := data[index]
				data[index] = data[index+1]
				data[index+1] = tmp

				fmt.Println(ag.Render(data, dims))
			}

			if sorted {
				return
			}
		}
	}
}
