package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StringToInt(str string) int {

	n, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("failed to parse int from '%s': %v", str, err))
	}

	return n
}

func StringToInt64(str string) int64 {

	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse int from '%s': %v", str, err))
	}

	return n
}

func StringToFloat(str string) float64 {

	n, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse float from '%s': %v", str, err))
	}

	return n
}

func StringToBool(str string) bool {

	b, err := strconv.ParseBool(str)
	if err != nil {
		panic(fmt.Sprintf("failed to parse bool from '%s': %v", str, err))
	}

	return b
}

func IntToString(n int) string {

	return strconv.Itoa(n)
}

func Int64ToString(n int64) string {

	return strconv.FormatInt(n, 10)
}

func FloatToString(f float64) string {

	return strconv.FormatFloat(f, 'f', -1, 64)
}

func FloatToStringPrecision(f float64, precision int) string {

	return strconv.FormatFloat(f, 'f', precision, 64)
}

func SliceToString[T any](slice []T) string {

	if len(slice) == 0 {

		return ""
	}
	
	var elements []string

	for _, item := range slice {

		elements = append(elements, fmt.Sprintf("%v", item))
	}
	
	return strings.Join(elements, ", ")
}

func Contains[T comparable](slice []T, variable T) bool {

	for _, value := range slice {

		if value == variable {

			return true
		}
	}

	return false
}

func ReadFile(filePath string) (string, error) {

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	return strings.ReplaceAll(string(content), "\n", ""), nil
}

func ReadFileBytes(filePath string) ([]byte, error) {

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	return content, nil
}

func ReadLines(filePath string) ([]string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	const maxCapacity = 1024 * 1024 // 1MB max line length
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	return lines, nil
}

// For large files
func ReadLinesStreaming(filePath string, processLine func(int, string) error) error {

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	lineNum := 0
	for scanner.Scan() {

		if err := processLine(lineNum, scanner.Text()); err != nil {	
			return fmt.Errorf("error processing line %d: %w", lineNum, err)
		}

		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	return nil
}

func ReadNonEmptyLines(filePath string) ([]string, error) {

	lines, err := ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	nonEmpty := make([]string, 0, len(lines))
	for _, line := range lines {

		if trimmed := strings.TrimSpace(line); trimmed != "" {
			
			nonEmpty = append(nonEmpty, line)
		}
	}

	return nonEmpty, nil
}

func ReadFileOrPanic(filePath string) string {

	content, err := ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	
	return content
}

func ReadLinesOrPanic(filePath string) []string {
	
	lines, err := ReadLines(filePath)
	if err != nil {
		panic(err)
	}
	
	return lines
}

func FileExists(filePath string) bool {

	_, err := os.Stat(filePath)

	return err == nil
}

func CountLines(filePath string) (int, error) {

	lines, err := ReadLines(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}

	return len(lines), nil
}
