package main

import (
	goUtils "AoC"
)

func main() {

	file := goUtils.ReadFileToString("../input/day3.txt")

	uniqueHouses := make(map[string]struct{})
	var xPosition int = 0
	var yPosition int = 0
	uniqueHouses[goUtils.IntToString(xPosition)+" "+goUtils.IntToString(yPosition)] = struct{}{}

	var xRobotPosition int = 0
	var yRobotPosition int = 0
	robotSantaHouses := make(map[string]struct{})
	robotSantaHouses["0 0"] = struct{}{}

	for _, char := range file {

		if string(char) == "<" {

			xPosition--
			uniqueHouses[goUtils.IntToString(xPosition)+" "+goUtils.IntToString(yPosition)] = struct{}{}

		} else if string(char) == ">" {

			xPosition++
			uniqueHouses[goUtils.IntToString(xPosition)+" "+goUtils.IntToString(yPosition)] = struct{}{}

		} else if string(char) == "^" {

			yPosition++
			uniqueHouses[goUtils.IntToString(xPosition)+" "+goUtils.IntToString(yPosition)] = struct{}{}

		} else if string(char) == "v" {

			yPosition--
			uniqueHouses[goUtils.IntToString(xPosition)+" "+goUtils.IntToString(yPosition)] = struct{}{}
		}
	}

	xPosition = 0
	yPosition = 0

	for i, char := range file {

		if i%2 == 0 {

			if string(char) == "<" {

				xPosition--
				robotSantaHouses[goUtils.IntToString(xPosition)+" "+goUtils.IntToString(yPosition)] = struct{}{}

			} else if string(char) == ">" {

				xPosition++
				robotSantaHouses[goUtils.IntToString(xPosition)+" "+goUtils.IntToString(yPosition)] = struct{}{}

			} else if string(char) == "^" {

				yPosition++
				robotSantaHouses[goUtils.IntToString(xPosition)+" "+goUtils.IntToString(yPosition)] = struct{}{}

			} else if string(char) == "v" {

				yPosition--
				robotSantaHouses[goUtils.IntToString(xPosition)+" "+goUtils.IntToString(yPosition)] = struct{}{}
			}

		} else {

			if string(char) == "<" {

				xRobotPosition--
				robotSantaHouses[goUtils.IntToString(xRobotPosition)+" "+goUtils.IntToString(yRobotPosition)] = struct{}{}

			} else if string(char) == ">" {

				xRobotPosition++
				robotSantaHouses[goUtils.IntToString(xRobotPosition)+" "+goUtils.IntToString(yRobotPosition)] = struct{}{}

			} else if string(char) == "^" {

				yRobotPosition++
				robotSantaHouses[goUtils.IntToString(xRobotPosition)+" "+goUtils.IntToString(yRobotPosition)] = struct{}{}

			} else if string(char) == "v" {

				yRobotPosition--
				robotSantaHouses[goUtils.IntToString(xRobotPosition)+" "+goUtils.IntToString(yRobotPosition)] = struct{}{}
			}
		}
	}

	p1Result := "Day 3 - Part 1: Santa delivers presents to " + goUtils.IntToString(len(uniqueHouses)) + " houses"
	goUtils.WriteToFileAtLine("../results.txt", p1Result, 5)

	p2Resut := "Day 3 - Part 2: Santa and the robot delivers presents to " + goUtils.IntToString(len(robotSantaHouses)) + " houses"
	goUtils.WriteToFileAtLine("../results.txt", p2Resut, 6)
}
