package main

import "fmt"

func main() {

	inputString := ReadFileOrPanic("input/day1.txt")

	var floor int
	var position int
	var basementFirstPosition int

	for _, char := range inputString {

		if string(char) == "(" {

			floor++
			position++

			if floor == -1 && basementFirstPosition == 0 {

				basementFirstPosition = position
			}

		} else if string(char) == ")" {

			floor--
			position++

			if floor == -1 && basementFirstPosition == 0 {

				basementFirstPosition = position
			}
		}
	}

	fmt.Printf("Santa's flood: %v.\n", floor)
	fmt.Printf("Santa's first basement position: %v.\n", basementFirstPosition)
}
