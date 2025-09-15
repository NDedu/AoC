package main

import (
	"fmt"
	"strings"
)

func main() {

	lines := ReadLinesOrPanic("input/day6.txt")

	var matrixPart1 [1000][1000]int
	var matrixPart2 [1000][1000]int

	var openLightsPart1 int
	var brightnessPart2 int

	for _, line := range lines {

		words := strings.Split(line, " ")
		
		if words[0] == "toggle" {

			startingCoordinates := strings.Split(words[1], ",")
			endingCoordinates := strings.Split(words[3], ",")

			startingX := StringToInt(startingCoordinates[0])
			startingY := StringToInt(startingCoordinates[1])
			endingX := StringToInt(endingCoordinates[0])
			endingY := StringToInt(endingCoordinates[1])

			matrixPart1 = lightConfig(startingX, startingY, endingX, endingY, "togglePart1", matrixPart1)
			matrixPart2 = lightConfig(startingX, startingY, endingX, endingY, "togglePart2", matrixPart2)
		
		} else if words[1] == "on" {
			
			startingCoordinates := strings.Split(words[2], ",")
			endingCoordinates := strings.Split(words[4], ",")

			startingX := StringToInt(startingCoordinates[0])
			startingY := StringToInt(startingCoordinates[1])
			endingX := StringToInt(endingCoordinates[0])
			endingY := StringToInt(endingCoordinates[1])

			matrixPart1 = lightConfig(startingX, startingY, endingX, endingY, "turnOnPart1", matrixPart1)
			matrixPart2 = lightConfig(startingX, startingY, endingX, endingY, "turnOnPart2", matrixPart2)
		
		} else if words[1] == "off" {

			startingCoordinates := strings.Split(words[2], ",")
			endingCoordinates := strings.Split(words[4], ",")
			
			startingX := StringToInt(startingCoordinates[0])
			startingY := StringToInt(startingCoordinates[1])
			endingX := StringToInt(endingCoordinates[0])
			endingY := StringToInt(endingCoordinates[1])

			matrixPart1 = lightConfig(startingX, startingY, endingX, endingY, "turnOffPart1", matrixPart1)
			matrixPart2 = lightConfig(startingX, startingY, endingX, endingY, "turnOffPart2", matrixPart2)
		}
	}

	for i := 0; i < 1000; i++ {

		for j := 0; j < 1000; j++ {

			if matrixPart1[i][j] == 1 {

				openLightsPart1++
			}

			brightnessPart2 += matrixPart2[i][j]
		}
	}

	fmt.Printf("Total number of open lights in part 1: %v.\n", openLightsPart1)
	fmt.Printf("Total brightness of lights in part 2: %v.\n", brightnessPart2)
}

func lightConfig(startX, startY, endX, endY int, action string, matrix [1000][1000]int) [1000][1000]int {

	for i := startX; i <= endX; i++ {

		for j := startY; j <= endY; j++ {

			switch action {
			case "togglePart1":
				
				if matrix[i][j] == 1 {

					matrix[i][j] = 0

				} else if matrix[i][j] == 0 {

					matrix[i][j] = 1
				}

			case "turnOnPart1":

				matrix[i][j] = 1
			
			case "turnOffPart1":

				matrix[i][j] = 0

			case "togglePart2":

				matrix[i][j] += 2

			case "turnOnPart2":

				matrix[i][j]++
			
			case "turnOffPart2":

				if matrix[i][j] > 0 {

					matrix[i][j]--
				}
			}
		}
	}

	return matrix
} 
