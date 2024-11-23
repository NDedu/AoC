package go_utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func StringToInt(str string) int {

	number, err := strconv.Atoi(str)
	if err != nil {

		log.Printf("Failed to convert string to int: %v", err)
		return 0
	}

	return number
}

func IntToString(number int) string {

	return strconv.Itoa(number)
}

func StringToFloat(str string) float64 {

	number, err := strconv.ParseFloat(str, 64)
	if err != nil {

		log.Printf("Failed to convert string to float: %v", err)
		return 0.0
	}

	return number
}

func StringToBool(str string) bool {

	boolean, err := strconv.ParseBool(str)
	if err != nil {

		log.Printf("Failed to convert string to boolean: %v", err)
		return false
	}

	return boolean
}

func ReadFileLines(filePath string) []string {

	absPath, err := filepath.Abs(filePath)
	if err != nil {

		log.Printf("Could not find path: %v", err)
		return nil
	}

	file, err := os.Open(absPath)
	if err != nil {

		log.Printf("Failed to open file: %v", err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {

		log.Printf("Failed to read file: %v", err)
		return nil
	}

	return lines
}

func ReadFileToString(filePath string) string {

	absPath, err := filepath.Abs(filePath)
	if err != nil {

		log.Printf("Could not find path: %v", err)
		return ""
	}

	file, err := os.ReadFile(absPath)
	if err != nil {

		log.Printf("Failed to read file: %v", err)
		return ""
	}
	str := strings.ReplaceAll(string(file), "\n", "")
	return str
}

func WriteToFile(filePath string, content string) {

	absPath, err := filepath.Abs(filePath)
	if err != nil {

		log.Printf("Could not find path: %v", err)
		return
	}

	writeError := os.WriteFile(absPath, []byte(content), 0644)
	if writeError != nil {

		log.Printf("Failed to write to file: %v", writeError)
		return
	}

	fmt.Println("Successfully wrote to the result to file")
}

func WriteToFileAtLine(filePath string, content string, lineNum int) error {

	if lineNum < 1 {

		log.Printf("Line number must be greater than 1")
		return nil
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {

		log.Printf("Could not find path: %v", err)
		return err
	}

	file, err := os.Open(absPath)
	if err != nil {

		log.Printf("Failed to open file to write: %v", err)
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {

		return err
	}

	for len(lines) < lineNum-1 {

		lines = append(lines, "")
	}

	if lineNum <= len(lines) {

		lines[lineNum-1] = content

	} else {

		lines = append(lines, content)
	}

	output := strings.Join(lines, "\n")
	if len(lines) > 0 {

		output += "\n"
	}

	return os.WriteFile(absPath, []byte(output), 0644)
}
