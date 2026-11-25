package main

import "core:fmt"
import vmem "core:mem/virtual"
import "core:strings"

INPUT :: #load("../input/day1.txt", string)

main :: proc() {
	arena: vmem.Arena
	defer vmem.arena_destroy(&arena)
	context.allocator = vmem.arena_allocator(&arena)

	floor_p1: int
	first_basement_step: int

	input := INPUT
	for line in strings.split_lines_iterator(&input) {
		for char, i in line {
			if char == '(' {
				floor_p1 += 1
				if first_basement_step == 0 && floor_p1 < 0 do first_basement_step = i + 1 // because first position is 1 and i is 0
			} else if char == ')' {
				floor_p1 -= 1
				if first_basement_step == 0 && floor_p1 < 0 do first_basement_step = i + 1
			}
		}
	}

	fmt.println(floor_p1)
	fmt.println(first_basement_step)
}
