package main

import (
	"fmt"
	"time"
	// tea "charm.land/bubbletea/v2"
)

type model struct {
	data []int
	graph ArrayGraph
}

func main() {
	bg := BarGraph{component: "▐█▌"}
	dims := Dimension{
		width: 225,
		height: 50,
	}

	bs := BubbleSort{}
	
	data := []int{0, 172, 45, 298, 13, 207, 89, 154, 266, 31, 74, 219, 6, 143, 255, 92, 168, 24, 301, 57, 184, 76, 212, 39, 147, 260, 18, 95, 233, 64, 121, 275, 52, 199, 8, 136, 221, 70, 158, 247, 34, 110, 289, 41, 173, 59, 204, 97, 262, 15}

	fmt.Println(bg.Render(data, dims))

	steps := bs.GetSortingSteps(data)
	_ = steps
	
	for _, val := range steps {
		fmt.Println(bg.Render(val, dims))
		time.Sleep(time.Millisecond * 25)
	}
	
}
