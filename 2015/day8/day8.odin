package main

import "core:fmt"
import vmem "core:mem/virtual"
import "core:strings"

INPUT :: #load("../input/day8.txt", string)

is_hex :: proc(char: u8) -> bool {
	return (char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F')
}

string_cleanup :: proc(str: string) -> string {
	if len(str) < 2 do return ""

	trimmed := str[1 : len(str) - 1]
	buf := make([]u8, len(trimmed))
	j := 0

	i := 0
	for i < len(trimmed) {
		if trimmed[i] == '\\' {
			switch trimmed[i + 1] {
			case '\\':
				buf[j] = '\\'
				i += 2
			case '"':
				buf[j] = '"'
				i += 2
			case 'x':
				if i + 3 < len(trimmed) && is_hex(trimmed[i + 2]) && is_hex(trimmed[i + 3]) {
					buf[j] = '?'
					i += 4
				} else {
					buf[j] = '\\'
					i += 1
				}
			case:
				buf[j] = '\\'
				i += 1
			}
		} else {
			buf[j] = trimmed[i]
			i += 1
		}
		j += 1
	}
	return string(buf[:j])
}

string_encode :: proc(str: string) -> string {
	buf := make([]u8, len(str) * 3)
	buf[0] = '"'
	j := 1

	for i := 0; i < len(str); i += 1 {
		if str[i] == '"' {
			buf[j] = '\\'
			buf[j + 1] = '"'
			j += 2
		} else if str[i] == '\\' {
			buf[j] = '\\'
			buf[j + 1] = '\\'
			j += 2
		} else {
			buf[j] = str[i]
			j += 1
		}
	}
	buf[j] = '"'
	j += 1
	return string(buf[:j])
}


main :: proc() {
	arena: vmem.Arena
	defer vmem.arena_destroy(&arena)
	context.allocator = vmem.arena_allocator(&arena)

	nr_chars_diff_p1: int
	nr_chars_diff_p2: int

	input := INPUT
	for line in strings.split_lines_iterator(&input) {
		nr_chars_diff_p1 += len(line) - len(string_cleanup(line))
		nr_chars_diff_p2 += len(string_encode(line)) - len(line)
	}

	fmt.printf("Part 1 number of characters difference: %d\n", nr_chars_diff_p1)
	fmt.printf("Part 2 number of characters difference: %d\n", nr_chars_diff_p2)
}
