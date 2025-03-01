package main

import (
	goUtils "AoC"
	"strings"
)

func main() {

	lines := goUtils.ReadFileLines("../input/day7.txt")

	for _, line := range lines {

		lineStrings := strings.Fields(line)

		if len(lineStrings) > 3 && lineStrings[0] == "NOT" {

		}
	}
}
