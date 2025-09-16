package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	linesPart1 := ReadLinesOrPanic("input/day7.txt")
	linesPart2 := ReadLinesOrPanic("input/day7.txt")

	signalsByWirePart1 := make(map[string]int)
	signalsByWirePart1 = bitOperations(signalsByWirePart1, linesPart1)

	signalsByWirePart2 := make(map[string]int)
	signalsByWirePart2["b"] = signalsByWirePart1["a"]
	signalsByWirePart2 = bitOperations(signalsByWirePart2, linesPart2)

	fmt.Printf("Value of A in part 1: %v.\n", signalsByWirePart1["a"])

	fmt.Printf("Value of A in part 2: %v.\n", signalsByWirePart2["a"])
}

func getValue(signalMap map[string]int, keyOrValue string) (int, error) {

	
	valueMap, exists := signalMap[keyOrValue]

	value, err := strconv.Atoi(keyOrValue)

	if exists {

		return valueMap, nil

	} else if err == nil {

		return value, nil 

	} else {

		return 0, err
	}
}

func bitOperations(signalsByWire map[string]int, lines []string) map[string]int {

	i := 0
	for {

		if len(lines) == 0 {

			break
		}

		words := strings.Split(lines[i], " ")

		if words[1] == "->" {

			//Check for part 2
			if _, exists := signalsByWire[words[2]]; exists {

				lines = append(lines[:i], lines[i+1:]...)
				i = 0
				continue
			}

			value, err := getValue(signalsByWire, words[0])

			if err == nil {

				signalsByWire[words[2]] = value
				lines = append(lines[:i], lines[i+1:]...)
				i = 0
				continue
			}

		} else if words[0] == "NOT" {
			
			value, err := getValue(signalsByWire, words[1])

			if err == nil {

				signalsByWire[words[3]] = ^value
				lines = append(lines[:i], lines[i+1:]...)
				i = 0
				continue
			}

		} else if words[1] == "OR" {

			value1, err1 := getValue(signalsByWire, words[0])
			value2, err2 := getValue(signalsByWire, words[2])

			if err1 == nil && err2 == nil {

				signalsByWire[words[4]] = value1 | value2
				lines = append(lines[:i], lines[i+1:]...)
				i = 0
				continue
			}

		} else if words[1] == "AND" {
			
			value1, err1 := getValue(signalsByWire, words[0])
			value2, err2 := getValue(signalsByWire, words[2])

			if err1 == nil && err2 == nil {

				signalsByWire[words[4]] = value1 & value2
				lines = append(lines[:i], lines[i+1:]...)
				i = 0
				continue
			}

		} else if words[1] == "LSHIFT" {
				
			value1, err1 := getValue(signalsByWire, words[0])
			value2, err2 := getValue(signalsByWire, words[2])

			if err1 == nil && err2 == nil {

				signalsByWire[words[4]] = value1 << value2
				lines = append(lines[:i], lines[i+1:]...)
				i = 0
				continue
			}		

		} else if words[1] == "RSHIFT" {
			
			value1, err1 := getValue(signalsByWire, words[0])
			value2, err2 := getValue(signalsByWire, words[2])

			if err1 == nil && err2 == nil {

				signalsByWire[words[4]] = value1 >> value2
				lines = append(lines[:i], lines[i+1:]...)
				i = 0
				continue
			}
		}

		i++
		if i == len(lines) {

			i = 0
		}
	}

	return signalsByWire
}
