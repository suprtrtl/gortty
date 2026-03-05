package main 

import (
	"fmt"
//	tea "charm.land/bubbletea/v2"
)

type model struct {
	data []int
	graph ArrayGraph
}

func main() {
	bg := BarGraph{component: "▐▌"}
	dims := Dimension{
		width: 48,
		height: 30,
	}
	
	data := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 26, 32, 45, 60, 82}

	fmt.Println(len(data))
	fmt.Println(bg.Render(data, dims))
}
