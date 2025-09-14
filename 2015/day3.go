package main

import "fmt"

func main() {

	inputString := ReadFileOrPanic("input/day3.txt")

	santaCoordinates := make(map[string]struct{})

	var santaXCoordinate int
	var santaYCoordinate int

	santaCoordinates[SliceToString([]int{santaXCoordinate, santaYCoordinate})] = struct{}{}

	// Will have a set to determine unique coordinates visited by both
	// then with a pointer will decide which coordinates to change
	robotAndSantaCoordinates := make(map[string]struct{})

	var santa2XCoordinate int
	var santa2YCoordinate int
	var robotXCoordinate int
	var robotYCoordinate int

	santaCoordinates[SliceToString([]int{santa2XCoordinate, santa2YCoordinate})] = struct{}{}

	for i, char := range inputString {

		var pointerXCoordinate *int
		var pointerYCoordinate *int

		if i % 2 == 0 {

			pointerXCoordinate = &santa2XCoordinate
			pointerYCoordinate = &santa2YCoordinate

		} else {

			pointerXCoordinate = &robotXCoordinate
			pointerYCoordinate = &robotYCoordinate
		}

		if string(char) == "^" {

			santaYCoordinate++
			(*pointerYCoordinate)++

		} else if string(char) == ">" {

			santaXCoordinate++
			(*pointerXCoordinate)++

		} else if string(char) == "v" {

			santaYCoordinate--
			(*pointerYCoordinate)--

		} else if string(char) == "<" {

			santaXCoordinate--
			(*pointerXCoordinate)--
		}

		santaCoordinates[SliceToString([]int{santaXCoordinate, santaYCoordinate})] = struct{}{}

		robotAndSantaCoordinates[SliceToString([]int{santa2XCoordinate, santa2YCoordinate})] = struct{}{}
		robotAndSantaCoordinates[SliceToString([]int{robotXCoordinate, robotYCoordinate})] = struct{}{}
	}

	fmt.Printf("Houses with at least one present: %v.\n", len(santaCoordinates))
	fmt.Printf("Houses with at least one present from Santa and the robot: %v.\n", len(robotAndSantaCoordinates))
}
