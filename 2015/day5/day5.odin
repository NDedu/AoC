package main

import "core:fmt"
import vmem "core:mem/virtual"
import "core:strings"

INPUT :: #load("../input/day5.txt", string)

p1_nice_string :: proc(str: string) -> bool {
	vowel_count: int
	twice_char: bool

	if strings.contains(str, "ab") || strings.contains(str, "cd") || strings.contains(str, "pq") || strings.contains(str, "xy") do return false

	for char, i in str {
		switch char {
			case 'a', 'e', 'i', 'o', 'u': vowel_count += 1
		}
		if !twice_char && i < len(str) - 1 && str[i] == str[i+1] do twice_char = true
		if twice_char && vowel_count >= 3 do return true
	}

	return false
}

p2_nice_string :: proc(str: string) -> bool {
	has_pair_chars: bool
	twice_char: bool

	for _, i in str {
		if !has_pair_chars && i < len(str) - 3 && strings.contains(str[i + 2:], str[i:i + 2]) do has_pair_chars = true
		if !twice_char && i < len(str) - 2 && str[i] == str[i+2] do twice_char = true
		if has_pair_chars && twice_char do return true
	}

	return false
}

main :: proc() {
	arena: vmem.Arena
	defer vmem.arena_destroy(&arena)
	context.allocator = vmem.arena_allocator(&arena)

	nice_strings_p1: int
	nice_strings_p2: int

	input := INPUT
	for line in strings.split_lines_iterator(&input) {
		if p1_nice_string(line) do nice_strings_p1 += 1
		if p2_nice_string(line) do nice_strings_p2 += 1
	}

	fmt.printf("Part 1 nice strings: %d\n", nice_strings_p1)
	fmt.printf("Part 2 nice strings: %d\n", nice_strings_p2)
}
