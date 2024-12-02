package main

import (
	goUtils "AoC"
	"strings"
	"unicode"
)

func hexStringReplace(str *string) {

	i := 0
	for {

		runes := []rune(*str)

		if runes[i] == '\\' && 
			runes[i+1] == 'x' && 
			unicode.Is(unicode.ASCII_Hex_Digit, rune(runes[i+2])) && 
			unicode.Is(unicode.ASCII_Hex_Digit, rune(runes[i+3])) {

			*str = string(runes[:i]) + "'" + string(runes[i+4:])
		}

		i++
		if i >= len(runes) - 3 {

			break
		}	
	}
}

func main() {

	lines := goUtils.ReadFileLines("../input/day8.txt")

	var charsInMemory int = 0
	var stringliteralsSize int = 0
	var encodedStringSize int = 0

	for _, line := range lines {

		newStr := line[1 : len(line)-1]
		hexStringReplace(&newStr)
		newStr = strings.ReplaceAll(newStr, "\\\"", "\"")
		newStr = strings.ReplaceAll(newStr, "\\\\", "\\")

		encodedString := line[1 : len(line)-1]
		encodedString = strings.ReplaceAll(encodedString, "\\", "\\\\")
		encodedString = strings.ReplaceAll(encodedString, "\"", "\\\"")
		encodedString = "\"\\\"" + encodedString + "\\\"\""

		charsInMemory += len(newStr)
		stringliteralsSize += len(line)
		encodedStringSize += len(encodedString)
	}

	p1Result := "Day 8 - Part 1: String literals - characters in memory: " + goUtils.IntToString(stringliteralsSize - charsInMemory)
	goUtils.WriteToFileAtLine("../results.txt", p1Result, 15)

	p2Result := "Day 8 - Part 2: Encoded string - string literals: " + goUtils.IntToString(encodedStringSize - stringliteralsSize)
	goUtils.WriteToFileAtLine("../results.txt", p2Result, 16)
}