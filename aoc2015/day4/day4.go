package main

import (
	goUtils "AoC"
	"crypto/md5"
	"encoding/hex"
)

func md5Hash(text string) string {

	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {

	fileString := goUtils.ReadFileToString("../input/day4.txt")

	var number int = 1
	for {

		md5 := md5Hash(fileString + goUtils.IntToString(number))
		if string(md5[0]) == "0" && string(md5[1]) == "0" && string(md5[2]) == "0" && string(md5[3]) == "0" && string(md5[4]) == "0" {
			break
		}

		number++
	}

	p1Result := "Day 4 - Part 1: The number where the hash starts with 5 zeroes is " + goUtils.IntToString(number)
	goUtils.WriteToFileAtLine("../results.txt", p1Result, 7)

	number = 1
	for {

		md5 := md5Hash(fileString + goUtils.IntToString(number))
		if string(md5[0]) == "0" && string(md5[1]) == "0" && string(md5[2]) == "0" && string(md5[3]) == "0" && string(md5[4]) == "0" && string(md5[5]) == "0" {
			break
		}

		number++
	}

	p2Result := "Day 4 - Part 2: The number where the hash starts with 6 zeroes is " + goUtils.IntToString(number)
	goUtils.WriteToFileAtLine("../results.txt", p2Result, 8)
}
