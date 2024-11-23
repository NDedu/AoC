package main

import (
	goUtils "AoC"
	"slices"
)

func duplicateChars(str string) bool {

	for i := range str {

		if i < len(str)-1 && str[i] == str[i+1] {

			return true
		}
	}

	return false
}

func wrongStrings(str string) bool {

	wrongString := []string{"ab", "cd", "pq", "xy"}

	for i := range str {

		if i < len(str)-1 && slices.Contains(wrongString, string(str[i])+string(str[i+1])) {

			return true
		}
	}

	return false
}

func numberOfVowels(str string) int {

	vowels := []string{"a", "e", "i", "o", "u"}
	var number int = 0

	for _, char := range str {

		if slices.Contains(vowels, string(char)) {

			number++
		}
	}

	return number
}

func pairOfLetters(str string) bool {

	var pairOfChars []string
	uniquePairs := make(map[string]struct{})

	i := 0
	for i < len(str)-1 {

		if i <= len(str)-3 && str[i] == str[i+1] && str[i] == str[i+2] {

			pairOfChars = append(pairOfChars, string(str[i])+string(str[i+1]))
			i += 2
			continue
		}

		if i < len(str)-1 {

			pairOfChars = append(pairOfChars, string(str[i])+string(str[i+1]))
		}
		i++
	}

	for _, pairs := range pairOfChars {

		uniquePairs[pairs] = struct{}{}
	}

	if len(pairOfChars) != len(uniquePairs) {

		return true
	}

	return false
}

func duplicateCharsInBetween(str string) bool {

	for i := range str {

		if i <= len(str)-3 && str[i] == str[i+2] {

			return true
		}
	}

	return false
}

func main() {

	lines, err := goUtils.ReadFileLines("../input/day5.txt")
	if err != nil {
		return
	}

	var numberNiceStringsP1 int = 0
	var numberNiceStringsP2 int = 0

	for _, line := range lines {

		if numberOfVowels(line) >= 3 && duplicateChars(line) && !wrongStrings(line) {

			numberNiceStringsP1++
		}

		if pairOfLetters(line) && duplicateCharsInBetween(line) {

			numberNiceStringsP2++
		}
	}

	p1Result := "Day 5 - Part 1: Number of nice strings " + goUtils.IntToString(numberNiceStringsP1)
	goUtils.WriteToFileAtLine("../results.txt", p1Result, 9)

	p2Result := "Day 5 - Part 2: Number of nice strings " + goUtils.IntToString(numberNiceStringsP2)
	goUtils.WriteToFileAtLine("../results.txt", p2Result, 10)
}
