package go_utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func StringToInt(str string) (int, error) {

	number, err := strconv.Atoi(str)
	if err != nil {

		log.Printf("Failed to convert string to int: %v", err)
		return 0, err
	}

	return number, nil
}

func IntToString(number int) string {

	return strconv.Itoa(number)
}

func StringToFloat(str string) (float64, error) {

	number, err := strconv.ParseFloat(str, 64)
	if err != nil {

		log.Printf("Failed to convert string to float: %v", err)
		return 0.0, err
	}

	return number, nil
}

func StringToBool(str string) (bool, error) {

	boolean, err := strconv.ParseBool(str)
	if err != nil {

		log.Printf("Failed to convert string to boolean: %v", err)
		return false, err
	}

	return boolean, nil
}

func ReadFileLines(filepath string) ([]string, error) {

	file, err := os.Open(filepath)
	if err != nil {

		log.Printf("Failed to open file: %v", err)
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {

		log.Printf("Failed to read file: %v", err)
		return nil, err
	}

	return lines, nil
}

func ReadFileToString(filepath string) (string, error) {

	file, err := os.ReadFile(filepath)
	if err != nil {

		log.Printf("Failed to read file: %v", err)
		return "", err
	}

	return string(file), nil
}
