package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var inputText string
	var floor int = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputText = scanner.Text()
	}

	for i := 0; i < len(inputText); i++ {
		if string(inputText[i]) == "(" {
			floor++
		} else if string(inputText[i]) == ")" {
			floor--
		}

		if floor == -1 {
			fmt.Println("First levels of basement: ", i+1)
		}
	}

	fmt.Println("The reached floor is: ", floor)

}
