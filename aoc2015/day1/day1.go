package main

import (
	goUtils "AoC"
)

func main() { 
	
	file, err := goUtils.ReadFileToString("../input/day1.txt")
	if err != nil {

		return
	}

	var floor int
	var position int
	var firstTimeBasement []int

	for _, character := range file {

		if string(character) == "(" {

			floor++
			position ++

		} else if string(character) == ")" {

			floor--
			position++

			if floor == -1 {

				firstTimeBasement = append(firstTimeBasement, position)
			}
		}
	}

	p1Result := "Day 1 - Part 1: Santa's floor is " + goUtils.IntToString(floor)
	goUtils.WriteToFileAtLine("../results.txt", p1Result, 1)

	p2Result := "Day 1 - Part 2: Santa first time basement " + goUtils.IntToString(firstTimeBasement[0])
	goUtils.WriteToFileAtLine("../results.txt", p2Result, 2)
}