package main

import (
	goUtils "AoC"
	"strings"
)

func switchLights(matrix [1000][1000]int, startLine int, endLine int, startCol int, endCol int, instruction string) [1000][1000]int {

	for i := startLine; i <= endLine; i++ {

		for j := startCol; j <= endCol; j++ {

			if instruction == "on" {

				matrix[i][j] = 1

			} else if instruction == "off" {

				matrix[i][j] = 0

			} else if instruction == "toggle" {

				if matrix[i][j] == 0 {

					matrix[i][j] = 1

				} else if matrix[i][j] == 1 {

					matrix[i][j] = 0
				}
			}
		}
	}

	return matrix
}

func increaseBright(matrix [1000][1000]int, startLine int, endLine int, startCol int, endCol int, instruction string) [1000][1000]int {

	for i := startLine; i <= endLine; i++ {

		for j := startCol; j <= endCol; j++ {

			if instruction == "on" {

				matrix[i][j]++

			} else if instruction == "off" {

				if matrix[i][j] > 0 {

					matrix[i][j]--
				}

			} else if instruction == "toggle" {

				matrix[i][j] += 2
			}
		}
	}

	return matrix
}

func main() {

	lines := goUtils.ReadFileLines("../input/day6.txt")

	var gridP1 [1000][1000]int
	var gridP2 [1000][1000]int

	for _, line := range lines {

		words := strings.Fields(line)

		if words[0] == "turn" && words[1] == "on" {

			startCoordinates := strings.Split(words[2], ",")
			endCoordinates := strings.Split(words[4], ",")

			gridP1 = switchLights(
				gridP1, 
				goUtils.StringToInt(startCoordinates[0]), 
				goUtils.StringToInt(endCoordinates[0]), 
				goUtils.StringToInt(startCoordinates[1]), 
				goUtils.StringToInt(endCoordinates[1]), 
				words[1])

			gridP2 = increaseBright(
				gridP2, 
				goUtils.StringToInt(startCoordinates[0]), 
				goUtils.StringToInt(endCoordinates[0]), 
				goUtils.StringToInt(startCoordinates[1]), 
				goUtils.StringToInt(endCoordinates[1]), 
				words[1])

		} else if words[0] == "turn" && words[1] == "off" {

			startCoordinates := strings.Split(words[2], ",")
			endCoordinates := strings.Split(words[4], ",")

			gridP1 = switchLights(
				gridP1, 
				goUtils.StringToInt(startCoordinates[0]), 
				goUtils.StringToInt(endCoordinates[0]), 
				goUtils.StringToInt(startCoordinates[1]), 
				goUtils.StringToInt(endCoordinates[1]), 
				words[1])

			gridP2 = increaseBright(
				gridP2, 
				goUtils.StringToInt(startCoordinates[0]), 
				goUtils.StringToInt(endCoordinates[0]), 
				goUtils.StringToInt(startCoordinates[1]), 
				goUtils.StringToInt(endCoordinates[1]), 
				words[1])

		} else if words[0] == "toggle" {

			startCoordinates := strings.Split(words[1], ",")
			endCoordinates := strings.Split(words[3], ",")

			gridP1 = switchLights(
				gridP1,
				goUtils.StringToInt(startCoordinates[0]), 
				goUtils.StringToInt(endCoordinates[0]), 
				goUtils.StringToInt(startCoordinates[1]), 
				goUtils.StringToInt(endCoordinates[1]), 
				words[0])

			gridP2 = increaseBright(
				gridP2, 
				goUtils.StringToInt(startCoordinates[0]), 
				goUtils.StringToInt(endCoordinates[0]), 
				goUtils.StringToInt(startCoordinates[1]), 
				goUtils.StringToInt(endCoordinates[1]), 
				words[0])
		}
	}

	var openLights int = 0
	var totalBright int = 0

	for i := 0; i < 1000; i++ {

		for j := 0; j < 1000; j++ {

			if gridP1[i][j] == 1 {

				openLights++
			}

			totalBright += gridP2[i][j]
		}
	}

	p1Result := "Day 6 - Part 1: Number of open lights: " + goUtils.IntToString(openLights)
	goUtils.WriteToFileAtLine("../results.txt", p1Result, 11)

	p2Result := "Day 6 - Part 2: Total brightness: " + goUtils.IntToString(totalBright)
	goUtils.WriteToFileAtLine("../results.txt", p2Result, 12)
}
