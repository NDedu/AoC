package main

import (
	"fmt"
	"strings"
)

func niceStringPart1(str string) bool {

	badWords := []string{"ab","cd", "pq", "xy"}
	vowels := []string{"a", "e", "i", "o", "u"}

	for _, badWord := range badWords {

		if strings.Contains(str, badWord) {

			return false
		}
	}

	var nrOfVowels int
	for _, char := range str {

		if Contains(vowels, string(char)) {

			nrOfVowels++
		}

		if(nrOfVowels >= 3) {

			break
		}
	}

	if nrOfVowels < 3 {

		return false
	}

	var containsDuplicateLetter bool = false
	for i, _ := range str {

		if i == len(str) - 1 {

			break
		}

		if str[i] == str[i + 1] {

			containsDuplicateLetter = true
			break
		}
	}

	if !containsDuplicateLetter {

		return false
	} 

	return true
}

func niceStringPart2(str string) bool {

	var containsDuplicatePair bool = false
	var containsRepeatingLetter bool = false

	for i := 0; i < len(str) - 2; i++ {

		leftString := str[:i]
		middleString := str[i:i+2]
		rightString := str[i+2:]

		if (len(leftString) > 1 && strings.Contains(leftString, middleString)) || 
			(len(rightString) > 1 && strings.Contains(rightString, middleString)) {

			containsDuplicatePair = true
		}

		if str[i] == str[i+2] {

			containsRepeatingLetter = true
		}

		if containsDuplicatePair && containsRepeatingLetter {

			return true
		}
	}

	return false
}

func main() {

	lines := ReadLinesOrPanic("input/day5.txt")

	var nrOfNiceStringsPart1 int
	var nrOfNiceStringsPart2 int

	for _, line := range lines {

		if niceStringPart1(line) {

			nrOfNiceStringsPart1++
		}

		if niceStringPart2(line) {

			nrOfNiceStringsPart2++
		}
	}

	fmt.Printf("Nice strings in part 1: %v.\n", nrOfNiceStringsPart1)
	fmt.Printf("Nice strings in part 2: %v.\n", nrOfNiceStringsPart2)
}
