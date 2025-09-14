package main

import (
	"fmt"
	"strings"
)

func main() {

	lines := ReadLinesOrPanic("input/day6.txt")

	//var matrix [1000][1000]int

	for _, line := range lines {

		words := strings.Split(line, " ")
		
		if words[0] == "toggle" {

			startingCoordinates := strings.Split(words[1], ",")
			endingCoordinates := strings.Split(words[1], ",")

		} else if words[1] == "on" {


		} else if words[1] == "off" {


		}
	}

}
