package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputText string

	for scanner.Scan() {
		inputText = scanner.Text()
	}

	var x int = 0
	var y int = 0

	var xCoordinates = make([]int, (len(inputText) + 1))
	var yCoordinates = make([]int, (len(inputText) + 1))

	// coordinates of starting house
	xCoordinates[len(inputText)] = 0
	yCoordinates[len(inputText)] = 0

	for i := 0; i < len(inputText); i++ {

		if string(inputText[i]) == ">" {
			x++
		} else if string(inputText[i]) == "<" {
			x--
		} else if string(inputText[i]) == "^" {
			y++
		} else if string(inputText[i]) == "v" {
			y--
		}

		xCoordinates[i] = x
		yCoordinates[i] = y
	}

	var uniqueHouses int = len(xCoordinates)

	for i := 0; i < len(xCoordinates); i++ {

		isDuplicate := false

		for j := i; j < len(xCoordinates); j++ {

			if xCoordinates[i] == xCoordinates[j] &&
				yCoordinates[i] == yCoordinates[j] &&
				i != j {

				isDuplicate = true
			}
		}

		if isDuplicate {
			uniqueHouses--
		}
	}

	fmt.Println("Number of houses with at least one present: ", uniqueHouses)

	// Part2

	var xSanta int = 0
	var ySanta int = 0
	var xRobotSanta int = 0
	var yRobotSanta int = 0

	var xSantaCoordinates = make([]int, 0)
	var ySantaCoordinates = make([]int, 0)

	var xRobotSantaCoordinates = make([]int, 0)
	var yRobotSantaCoordinates = make([]int, 0)
	xSantaCoordinates = append(xSantaCoordinates, 0)
	ySantaCoordinates = append(ySantaCoordinates, 0)

	xRobotSantaCoordinates = append(xRobotSantaCoordinates, 0)
	yRobotSantaCoordinates = append(yRobotSantaCoordinates, 0)

	for i := 0; i < len(inputText); i++ {

		if i%2 == 0 {

			if string(inputText[i]) == ">" {
				xSanta++
			} else if string(inputText[i]) == "<" {
				xSanta--
			} else if string(inputText[i]) == "^" {
				ySanta++
			} else if string(inputText[i]) == "v" {
				ySanta--
			}

			xSantaCoordinates = append(xSantaCoordinates, xSanta)
			ySantaCoordinates = append(ySantaCoordinates, ySanta)

		} else if i%2 != 0 {

			if string(inputText[i]) == ">" {
				xRobotSanta++
			} else if string(inputText[i]) == "<" {
				xRobotSanta--
			} else if string(inputText[i]) == "^" {
				yRobotSanta++
			} else if string(inputText[i]) == "v" {
				yRobotSanta--
			}

			xRobotSantaCoordinates = append(xRobotSantaCoordinates, xRobotSanta)
			yRobotSantaCoordinates = append(yRobotSantaCoordinates, yRobotSanta)
		}
	}

	xVisitedHousesCoordinates := xSantaCoordinates
	yVisitedHousesCoordinates := ySantaCoordinates

	for i := 0; i < len(xRobotSantaCoordinates); i++ {

		xVisitedHousesCoordinates = append(xVisitedHousesCoordinates, xRobotSantaCoordinates[i])
		yVisitedHousesCoordinates = append(yVisitedHousesCoordinates, yRobotSantaCoordinates[i])
	}

	var combinedUniqueVisitedHouses int = len(xVisitedHousesCoordinates)

	for i := 0; i < len(xVisitedHousesCoordinates); i++ {

		isDuplicate := false

		for j := i; j < len(xVisitedHousesCoordinates); j++ {

			if xVisitedHousesCoordinates[i] == xVisitedHousesCoordinates[j] &&
				yVisitedHousesCoordinates[i] == yVisitedHousesCoordinates[j] &&
				i != j {

				isDuplicate = true
			}
		}

		if isDuplicate {
			combinedUniqueVisitedHouses--
		}
	}

	fmt.Println("Santa and Robot Santa visited a number of unique houses: ", combinedUniqueVisitedHouses)
}
