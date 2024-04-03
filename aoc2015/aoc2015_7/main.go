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

func convertInt(str string) int {

	number, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}

	return number
}

func sortDesc(slice []int) []int {

	sort.Slice(slice, func(i, j int) bool { return slice[i] > slice[j] })

	return slice
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var input []string
	results := make(map[string]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		input = append(input, line)
	}

	// this is a mess, i should have made a function for the repeating logic but i am not rewriting this
	// hey, as long as it works
	for {

		var indexToRemove []int

		for i, line := range input {

			values := strings.Fields(line)

			if values[0] == "NOT" {

				wire1 := convertInt(values[1])

				if wire1 != -1 {

					results[values[3]] = ^wire1
					indexToRemove = append(indexToRemove, i)

				} else {

					value, ok := results[values[1]]
					if ok {

						results[values[3]] = ^value
						indexToRemove = append(indexToRemove, i)
					}
				}

			} else if values[1] == "AND" {

				wire1 := convertInt(values[0])
				wire2 := convertInt(values[2])

				if wire1 != -1 && wire2 != -1 {

					results[values[4]] = wire1 & wire2
					indexToRemove = append(indexToRemove, i)

				} else if wire1 != -1 && wire2 == -1 {

					value, ok := results[values[2]]
					if ok {

						results[values[4]] = wire1 & value
						indexToRemove = append(indexToRemove, i)
					}

				} else if wire1 == -1 && wire2 != -1 {

					value, ok := results[values[0]]
					if ok {

						results[values[4]] = value & wire2
						indexToRemove = append(indexToRemove, i)
					}

				} else if wire1 == -1 && wire2 == -1 {

					value1, ok1 := results[values[0]]

					value2, ok2 := results[values[2]]
					if ok1 && ok2 {

						results[values[4]] = value1 & value2
						indexToRemove = append(indexToRemove, i)
					}
				}

			} else if values[1] == "OR" {

				wire1 := convertInt(values[0])
				wire2 := convertInt(values[2])

				if wire1 != -1 && wire2 != -1 {

					results[values[4]] = wire1 | wire2
					indexToRemove = append(indexToRemove, i)

				} else if wire1 != -1 && wire2 == -1 {

					value, ok := results[values[2]]
					if ok {

						results[values[4]] = wire1 | value
						indexToRemove = append(indexToRemove, i)
					}

				} else if wire1 == -1 && wire2 != -1 {

					value, ok := results[values[0]]
					if ok {

						results[values[4]] = value | wire2
						indexToRemove = append(indexToRemove, i)
					}

				} else if wire1 == -1 && wire2 == -1 {

					value1, ok1 := results[values[0]]

					value2, ok2 := results[values[2]]
					if ok1 && ok2 {

						results[values[4]] = value1 | value2
						indexToRemove = append(indexToRemove, i)
					}
				}

			} else if values[1] == "LSHIFT" {

				wire1 := convertInt(values[0])
				wire2 := convertInt(values[2])

				if wire1 != -1 && wire2 != -1 {

					results[values[4]] = wire1 << wire2
					indexToRemove = append(indexToRemove, i)

				} else if wire1 != -1 && wire2 == -1 {

					value, ok := results[values[2]]
					if ok {

						results[values[4]] = wire1 << value
						indexToRemove = append(indexToRemove, i)
					}

				} else if wire1 == -1 && wire2 != -1 {

					value, ok := results[values[0]]
					if ok {

						results[values[4]] = value << wire2
						indexToRemove = append(indexToRemove, i)
					}

				} else if wire1 == -1 && wire2 == -1 {

					value1, ok1 := results[values[0]]

					value2, ok2 := results[values[2]]
					if ok1 && ok2 {

						results[values[4]] = value1 << value2
						indexToRemove = append(indexToRemove, i)
					}
				}

			} else if values[1] == "RSHIFT" {

				wire1 := convertInt(values[0])
				wire2 := convertInt(values[2])

				if wire1 != -1 && wire2 != -1 {

					results[values[4]] = wire1 >> wire2
					indexToRemove = append(indexToRemove, i)

				} else if wire1 != -1 && wire2 == -1 {

					value, ok := results[values[2]]
					if ok {

						results[values[4]] = wire1 >> value
						indexToRemove = append(indexToRemove, i)
					}

				} else if wire1 == -1 && wire2 != -1 {

					value, ok := results[values[0]]
					if ok {

						results[values[4]] = value >> wire2
						indexToRemove = append(indexToRemove, i)
					}

				} else if wire1 == -1 && wire2 == -1 {

					value1, ok1 := results[values[0]]

					value2, ok2 := results[values[2]]
					if ok1 && ok2 {

						results[values[4]] = value1 >> value2
						indexToRemove = append(indexToRemove, i)
					}
				}

			} else if values[1] == "->" {

				wire1 := convertInt(values[0])

				if wire1 != -1 {

					results[values[2]] = wire1
					indexToRemove = append(indexToRemove, i)

				} else {

					value, ok := results[values[0]]
					if ok {

						results[values[2]] = value
						indexToRemove = append(indexToRemove, i)
					}
				}
			}
		}

		for _, index := range sortDesc(indexToRemove) {

			input = append(input[:index], input[index+1:]...)
		}

		if len(input) == 0 {

			break
		}
	}

	fmt.Println("Value of a is: ", results["a"])
}
