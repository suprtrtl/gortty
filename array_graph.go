package main

import (
	//	"fmt"
	//"math"
	"slices"
	"strings"
	"unicode/utf8"
)

type Dimension struct {
	width   int
	height  int
	spacing int
}

func NewDimension(w int, h int, spacing int) Dimension {
	return Dimension{
		width:   w,
		height:  h,
		spacing: spacing,
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

	componentSize := utf8.RuneCountInString(bg.component)

	// Calculates the max amount of space required to render
	maxDataWidth := len(data)*componentSize + 2*window.spacing

	if window.width < maxDataWidth {
		return "Window dimensions too small"
	}

	maxDataHeight := 0

	for _, val := range data {
		if val < maxDataHeight {
			continue
		}

		maxDataHeight = val
	}

	numChars := window.height - 2*window.spacing

	dataScale := float64(maxDataHeight) / float64(numChars)

	centerSpacing := (window.width - maxDataWidth) / 2

	strSlice := []string{}


	for i := range numChars {
		var tmpStr strings.Builder

		tmpStr.WriteString(strings.Repeat(" ", centerSpacing))

		for _, dataVal := range data {
			if float64(dataVal) >= (float64(i) * dataScale) {
				tmpStr.WriteString(bg.component)
			} else {
				tmpStr.WriteString(strings.Repeat(" ", componentSize))
			}
		}

		strSlice = append(strSlice, tmpStr.String()+"\n")
	}

	strSlice = append(strSlice, strings.Repeat("\n", window.spacing))

	slices.Reverse(strSlice)

	s := strings.Join(strSlice, "")

	return s
}
