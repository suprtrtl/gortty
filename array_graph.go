package main

import (
	color "github.com/fatih/color"
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
	Render([]int, Dimension, highlightMap) string
}

type BarGraph struct {
	component string
}

func (bg BarGraph) Render(data []int, window Dimension, hl highlightMap) string {

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

	for char := range numChars {
		row := bg.WriteRow(data, componentSize, numChars, dataScale, centerSpacing, char, hl)
		strSlice = append(strSlice, row)
	}

	strSlice = append(strSlice, strings.Repeat("\n", window.spacing))

	slices.Reverse(strSlice)

	s := strings.Join(strSlice, "")

	return s
}

func (bg BarGraph) WriteRow(data []int, componentSize int, numChars int, dataScale float64, centerSpacing int, char int, hl highlightMap) string {
	var builder strings.Builder

	builder.WriteString(strings.Repeat(" ", centerSpacing))

	for index, dataVal := range data {
		if float64(dataVal) >= (float64(char) * dataScale) {
			bg.SetComponentColor(&builder, index, hl)
		} else {
			builder.WriteString(strings.Repeat(" ", componentSize))
		}
	}
	return builder.String() + "\n"
}

func (bg BarGraph) SetComponentColor(builder *strings.Builder, index int, hl highlightMap) {
	if _, ok := hl[index]; ok {
		builder.WriteString(color.GreenString(bg.component))
	} else {
		builder.WriteString(bg.component)
	}
}
