package main

import (
	"fmt"
	"sort"
	"strings"
)

func neededPaper(l, w, h int) int {

	sizes := []int{l, w, h}
	sort.Ints(sizes)

	return 2 * l * w + 2 * w * h + 2 * h * l + sizes[0] * sizes[1]
}

func neededRibbon(l, w, h int) int {

	sizes := []int{l, w, h}
	sort.Ints(sizes)

	return sizes[0] * 2 + sizes[1] * 2 + l * w * h
}

func main() {

	fileLines := ReadLinesOrPanic("input/day2.txt")

	var totalSquareFeetPaper int
	var totalFeetRibbon int

	for _, line := range fileLines {

		sizesString := strings.Split(line, "x")

		totalSquareFeetPaper += neededPaper(StringToInt(sizesString[0]), StringToInt(sizesString[1]), StringToInt(sizesString[2]))
		totalFeetRibbon += neededRibbon(StringToInt(sizesString[0]), StringToInt(sizesString[1]), StringToInt(sizesString[2]))
	}

	fmt.Printf("Total square feet of paper: %v.\n", totalSquareFeetPaper)
	fmt.Printf("Total feet of ribbon: %v.\n", totalFeetRibbon)
}
