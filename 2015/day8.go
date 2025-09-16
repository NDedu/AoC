package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	inputStringList := ReadLinesOrPanic("input/day8.txt")
	
	var numberOfCharactersLiterals int
	var numberOfCharactersMemory int
	var encodedStringSize int

	for _, inputString := range inputStringList {

		formattedString := inputString[1:len(inputString) - 1]
		hexStringReplace(&formattedString)
		formattedString = strings.ReplaceAll(formattedString, "\\\"", "\"")
		formattedString = strings.ReplaceAll(formattedString, "\\\\", "\\")

		encodedString := inputString[1 : len(inputString)-1]
		encodedString = strings.ReplaceAll(encodedString, "\\", "\\\\")
		encodedString = strings.ReplaceAll(encodedString, "\"", "\\\"")
		encodedString = "\"\\\"" + encodedString + "\\\"\""

		numberOfCharactersLiterals += len(inputString)
		numberOfCharactersMemory += len(formattedString)
		encodedStringSize += len(encodedString)
	}

	fmt.Printf("Character literals - characters in memory: %v.\n", numberOfCharactersLiterals - numberOfCharactersMemory)
	fmt.Printf("Encoded string size - character literals: %v.\n", encodedStringSize - numberOfCharactersLiterals)
}

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
