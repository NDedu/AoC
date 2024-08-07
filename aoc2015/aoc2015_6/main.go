package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertToInt(s string) int {

	number, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return number
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	var openLightsCount int = 0
	var totalBrightness int = 0

	matrix := [1000][1000]int{}

	for scanner.Scan() {

		input = scanner.Text()

		replacer := strings.NewReplacer(",", " ")
		input = replacer.Replace(string(input))
		instructions := strings.Fields(input)

		if string(instructions[0]) == "turn" &&
			string(instructions[1]) == "on" {

			x1 := convertToInt(string(instructions[2]))
			y1 := convertToInt(string(instructions[3]))
			x2 := convertToInt(string(instructions[5]))
			y2 := convertToInt(string(instructions[6]))

			for i := y1; i <= y2; i++ {

				for j := x1; j <= x2; j++ {

					// Part 1
					// matrix[i][j] = 1

					// Part 2
					matrix[i][j]++
				}
			}

		} else if string(instructions[0]) == "turn" &&
			string(instructions[1]) == "off" {

			x1 := convertToInt(string(instructions[2]))
			y1 := convertToInt(string(instructions[3]))
			x2 := convertToInt(string(instructions[5]))
			y2 := convertToInt(string(instructions[6]))

			for i := y1; i <= y2; i++ {

				for j := x1; j <= x2; j++ {

					// Part 1
					// matrix[i][j] = 0

					// Part 2
					if matrix[i][j] > 0 {

						matrix[i][j]--
					}
				}
			}

		} else if string(instructions[0]) == "toggle" {

			x1 := convertToInt(string(instructions[1]))
			y1 := convertToInt(string(instructions[2]))
			x2 := convertToInt(string(instructions[4]))
			y2 := convertToInt(string(instructions[5]))

			for i := y1; i <= y2; i++ {

				for j := x1; j <= x2; j++ {

					// Part 1
					/*
					   if matrix[i][j] == 0 {

					     matrix[i][j] = 1
					   } else if matrix[i][j] == 1 {

					     matrix[i][j] = 0
					   }
					*/

					// Part 2
					matrix[i][j] += 2
				}
			}
		}

	}

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix); j++ {

			// Part 1
			if matrix[i][j] == 1 {

				openLightsCount++
			}

			// Part 2
			totalBrightness += matrix[i][j]
		}
	}

	fmt.Println("Number of open lights: ", openLightsCount)
	fmt.Println("Total brightness: ", totalBrightness)
}
