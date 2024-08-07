package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func GetMD5Hash(text string) string {

	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {

	var input string = "iwrupvqb"
	var value int = 0

	for i := 0; i < 1000000000; i++ {

		var secret string = input + strconv.Itoa(i)
		hash := GetMD5Hash(secret)

		if string(hash[0]) == "0" &&
			string(hash[1]) == "0" &&
			string(hash[2]) == "0" &&
			string(hash[3]) == "0" &&
			string(hash[4]) == "0" &&
			string(hash[5]) == "0" {

			value = i
			break
		}

	}

	fmt.Println(value)
}
