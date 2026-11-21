package main

import "core:fmt"
import vmem "core:mem/virtual"
import "core:strings"

INPUT :: #load("../input/day2.txt", string)

main :: proc() {
	arena: vmem.Arena
	defer vmem.arena_destroy(&arena)
	context.allocator = vmem.arena_allocator(&arena)

	input := INPUT
	for line in strings.split_lines_iterator(&input) {
		fmt.println(line)
	}
}
