package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isHexadecimal(char byte) bool {

	if (char >= '0' && char <= '9') || (char >= 'A' && char <= 'F') || (char >= 'a' && char <= 'f') {

		return true

	} else {

		return false
	}
}

func logic(line string, numStringCharacters int) int {

	for i := 1; i < len(line)-1; i++ {

		if line[i] == '\\' && line[i+1] == '\\' && i < (len(line)-2) {

			numStringCharacters++
			i += 1

		} else if line[i] == '\\' && line[i+1] == '"' && i < (len(line)-2) {

			numStringCharacters++
			i += 1

		} else if line[i] == '\\' &&
			line[i+1] == 'x' &&
			i < (len(line)-4) &&
			isHexadecimal(line[i+2]) &&
			isHexadecimal(line[i+3]) {

			numStringCharacters++
			i += 3

		} else if len(line) != 2 {

			numStringCharacters++
		}
	}

	return numStringCharacters
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {

		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numStringLiteralsP1 int
	var numStringCharactersP1 int
	var numStringLiteralsP2 int

	for scanner.Scan() {

		var encodedLine string
		line := scanner.Text()
		numStringLiteralsP1 += len(line)
		numStringCharactersP1 = logic(line, numStringCharactersP1)

		startingEncodingLine := string(line[0]) + string('\\') + string('"') + line[1:(len(line)-1)] + string('\\') + string('"') + string(line[len(line)-1])
		encodedLine = startingEncodingLine

		for i := 3; i < len(startingEncodingLine)-3; i++ {

			if startingEncodingLine[i] == '\\' || startingEncodingLine[i] == '"' {

				encodedLine = string(encodedLine[0:i]) + string('\\') + encodedLine[i:]
			}
		}

		numStringLiteralsP2 += len(encodedLine)
	}

	fmt.Println("Part 1 answer: ", numStringLiteralsP1-numStringCharactersP1)
	fmt.Println("Part 2 answer: ", numStringLiteralsP2-numStringLiteralsP1)
}
