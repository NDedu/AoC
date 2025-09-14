package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {

	inputString := ReadFileOrPanic("input/day4.txt")

	part1repeat := 5
	part2repeat := 6

	var part1string string = ""
	var part2string string = ""

	n := 0
	for {

		nString := IntToString(n)
		
		hash := md5.Sum([]byte(inputString + nString))
		hashString := hex.EncodeToString(hash[:])

		if hashString[:part1repeat] == strings.Repeat("0", part1repeat) && part1string == "" {

			part1string = nString
		}

		if  hashString[:part2repeat] == strings.Repeat("0", part2repeat) && part2string == "" {

			part2string = nString
		}
		

		if part1string != "" && part2string != "" {
			break
		}

		n++
	}

	fmt.Println("Number where the hash starts with 5 zeros: " + part1string)
	fmt.Println("Number where the hast starts with 6 zeros: " + part2string)
}
