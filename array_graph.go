package main

import (
	//	"fmt"
	//"math"
	"math"
	"slices"
	"strings"
	"unicode/utf8"
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

	componentSize := utf8.RuneCountInString(bg.component)

	if window.width < len(data) * componentSize {
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
				tmpStr.WriteString(strings.Repeat(" ", componentSize))
			}
		}

		strSlice = append(strSlice, tmpStr.String()+"\n")
	}

	slices.Reverse(strSlice)

	s := strings.Join(strSlice, "")

	return s
}

type BraileDots [2]uint8

var BrailleChars = map[BraileDots]string{
	{0, 0}: " ",
	{1, 0}: "⡀",
	{2, 0}: "⡄",
	{3, 0}: "⡆",
	{4, 0}: "⡇",
	{0, 1}: "⢀",
	{1, 1}: "⣀",
	{2, 1}: "⣄",
	{3, 1}: "⣆",
	{4, 1}: "⣇",
	{0, 2}: "⢠",
	{1, 2}: "⣠",
	{2, 2}: "⣤",
	{3, 2}: "⣦",
	{4, 2}: "⣧",
	{0, 3}: "⢰",
	{1, 3}: "⣰",
	{2, 3}: "⣴",
	{3, 3}: "⣶",
	{4, 3}: "⣷",
	{0, 4}: "⢸",
	{1, 4}: "⣸",
	{2, 4}: "⣼",
	{3, 4}: "⣾",
	{4, 4}: "⣿",
}

type BrailleGraph struct {}

func (bg BrailleGraph) Render(data []int, window Dimension) string {

	// Make it align with the braile characters and the for loop below
	if len(data) % 2 == 1 {
		data = append(data, 0)
	}

	if window.width * 2 < len(data) {
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

		for j := 0; j < len(data); j += 2 {
			dataVals := [2]int{data[j], data[j+1]}

			var (
				l,r uint8 = 0,0
			)

			if float64(dataVals[0]) >= (float64(i) * dataScale) {
				relative := float64(dataVals[0]) - (float64(i) * dataScale)
				// Divide by the height of each pip in braile chars
				l = uint8(math.Ceil(relative / 4))
			} 	

			if float64(dataVals[1]) >= (float64(i) * dataScale) {
				relative := float64(dataVals[1]) - (float64(i) * dataScale)
				r = uint8(math.Ceil(relative / 4))
			} 	

			dots := BraileDots{l,r}

			tmpStr.WriteString(BrailleChars[dots])
		}

		strSlice = append(strSlice, tmpStr.String()+"\n")
	}

	slices.Reverse(strSlice)

	s := strings.Join(strSlice, "")

	return s
}


