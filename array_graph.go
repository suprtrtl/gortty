package main

import (
	//	"fmt"
	//"math"
	"slices"
	"strings"
)

type Dimension struct {
	width  int
	height int
}

func NewDimension(w int, h int) Dimension {
	return Dimension{
		width:  w,
		height: h,
	}
}

type ArrayGraph interface {
	// Returns a string fit for rendering with bubble tea
	Render([]int, Dimension) string
}

type BarGraph struct {
	component string
}

func (bg BarGraph) Render(data []int, window Dimension) string {

	if window.width < len(data) {
		return "Window dimenstions too small"
	}

	maxDataHeight := 0

	for _, val := range data {
		if val < maxDataHeight {
			continue
		}

		maxDataHeight = val
	}

	dataScale := float64(maxDataHeight) / float64(window.height)

	numChars := window.height

	strSlice := []string{}

	for i := range numChars {
		var tmpStr strings.Builder

		for _, dataVal := range data {
			if float64(dataVal) >= (float64(i) * dataScale) {
				tmpStr.WriteString(bg.component)
			} else {
				tmpStr.WriteString(" ")
			}
		}

		strSlice = append(strSlice, tmpStr.String()+"\n")
	}

	slices.Reverse(strSlice)

	s := strings.Join(strSlice, "")

	return s
}
