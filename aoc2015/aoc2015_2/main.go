package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calculatePackage(l int, w int, h int) int {

	var side1 int = l * w
	var side2 int = w * h
	var side3 int = h * l

	sides := []int{side1, side2, side3}

	minSide := sides[0]

	for _, side := range sides {
		if side < minSide {
			minSide = side
		}
	}

	var formula int = 2*l*w + 2*w*h + 2*h*l + minSide

	return formula
}

func calculateRibbon(l int, w int, h int) int {

	lengths := []int{l, w, h}

	sort.Ints(lengths)

	var ribbonWrap int = 0

	for i := 0; i < (len(lengths) - 1); i++ {
		ribbonWrap += lengths[i] * 2
	}

	var formula int = ribbonWrap + l*w*h

	return formula

}

func convertToInt(s string) int {

	number, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return number
}

func main() {

	var totalPackageSize int = 0
	var totalRibbonSize int = 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var input string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = scanner.Text()

		replacer := strings.NewReplacer("x", " ")
		input := replacer.Replace(string(input))

		values := strings.Fields(input)
		var numbers = make([]string, len(values))

		for i, value := range values {
			numbers[i] = value
		}

		var packageSize int = calculatePackage(convertToInt(numbers[0]), convertToInt(numbers[1]), convertToInt(numbers[2]))
		totalPackageSize += packageSize

		var ribbonSize int = calculateRibbon(convertToInt(numbers[0]), convertToInt(numbers[1]), convertToInt(numbers[2]))
		totalRibbonSize += ribbonSize
	}

	fmt.Println("Total wrapping paper: ", totalPackageSize)
	fmt.Println("Total ribbon: ", totalRibbonSize)
}
