package main

import (
	goUtils "AoC"
	"sort"
	"strings"
)

func boxTotalArea(length int, heigth int, width int) int {

	surfaceArea := 2*length*width + 2*width*heigth + 2*heigth*length

	sides := []int{length, heigth, width}
	sort.Ints(sides)

	return surfaceArea + sides[0]*sides[1]
}

func ribbonNeeded(length int, heigth int, width int) int {

	sides := []int{length, heigth, width}
	sort.Ints(sides)

	return length*heigth*width + sides[0]*2 + sides[1]*2
}

func main() {

	lines, err := goUtils.ReadFileLines("../input/day2.txt")
	if err != nil {
		return
	}

	var totalWrappingPaper int = 0
	var totalRibbon int = 0

	for _, line := range lines {

		var boxSizes []int

		sidesSize := strings.Split(line, "x")
		for _, sideSize := range sidesSize {

			sideSizeInt, err := goUtils.StringToInt(sideSize)
			if err != nil {
				return
			}

			boxSizes = append(boxSizes, sideSizeInt)
		}

		totalWrappingPaper += boxTotalArea(boxSizes[0], boxSizes[1], boxSizes[2])
		totalRibbon += ribbonNeeded(boxSizes[0], boxSizes[1], boxSizes[2])
	}

	p1Result := "Day 2 - Part 1: Total wrapping paper needed: " + goUtils.IntToString(totalWrappingPaper)
	goUtils.WriteToFileAtLine("../results.txt", p1Result, 3)

	p2Result := "Day 2 - Part 2: Total ribbon needed: " + goUtils.IntToString(totalRibbon)
	goUtils.WriteToFileAtLine("../results.txt", p2Result, 4)
}
