package main

import "core:strconv"
import "core:fmt"
import vmem "core:mem/virtual"
import "core:strings"

INPUT :: #load("../input/day7.txt", string)

get_value :: proc(m: ^map[string]u16, str: string) -> (u16, bool) {
	if value, ok := strconv.parse_int(str); ok do return u16(value), true
	value, ok := m[str]
	return value, ok
}

bitwise_logc :: proc(bitwise_map: ^map[string]u16, str: string) -> bool {
	parts := strings.fields(str)

	if parts[0] == "NOT" {
		value, ok := get_value(bitwise_map, parts[1])
		if ok {
			bitwise_map[parts[3]] = ~value
			return true
		}
	} else if parts[1] == "AND" {
		value1, ok1 := get_value(bitwise_map, parts[0])
		value2, ok2 := get_value(bitwise_map, parts[2])
		if ok1 && ok2 {
			bitwise_map[parts[4]] = value1 & value2
			return true
		}
	} else if parts[1] == "OR" {
		value1, ok1 := get_value(bitwise_map, parts[0])
		value2, ok2 := get_value(bitwise_map, parts[2])
		if ok1 && ok2 {
			bitwise_map[parts[4]] = value1 | value2
			return true
		}
	} else if parts[1] == "LSHIFT" {
		value1, ok1 := get_value(bitwise_map, parts[0])
		value2, ok2 := get_value(bitwise_map, parts[2])
		if ok1 && ok2 {
			bitwise_map[parts[4]] = value1 << value2
			return true
		}
	} else if parts[1] == "RSHIFT" {
		value1, ok1 := get_value(bitwise_map, parts[0])
		value2, ok2 := get_value(bitwise_map, parts[2])
		if ok1 && ok2 {
			bitwise_map[parts[4]] = value1 >> value2
			return true
		}
	} else {
		value, ok := get_value(bitwise_map, parts[0])
		if ok {
			bitwise_map[parts[2]] = value
			return true
		}
	}
	return false
}

main :: proc() {
	arena: vmem.Arena
	defer vmem.arena_destroy(&arena)
	context.allocator = vmem.arena_allocator(&arena)

	wires1: map[string]u16
	p1_lines: [dynamic]string
	wires2: map[string]u16
	p2_lines: [dynamic]string

	input := INPUT
	for line in strings.split_lines_iterator(&input) {
		append(&p1_lines, line)
		append(&p2_lines, line)
	}

	for len(p1_lines) > 0 {
		#reverse for line, i in p1_lines {
			if bitwise_logc(&wires1, line) do unordered_remove(&p1_lines, i)
		}
	}

	wires2["b"] = wires1["a"]
	for len(p2_lines) > 0 {
		#reverse for line, i in p2_lines {
			if bitwise_logc(&wires2, line) do unordered_remove(&p2_lines, i)
		}
	}

	fmt.println("Part 1 a =", wires1["a"])
	fmt.println("Part 2 a =", wires2["a"])
}
