package main

import "core:slice"
import "core:strconv"
import "core:fmt"
import vmem "core:mem/virtual"
import "core:strings"

INPUT :: #load("../input/day2.txt", string)

box_surface_area :: proc(l: int, w: int, h: int) -> int {
	return 2 * l * w + 2 * w * h + 2 * h * l
}

main :: proc() {
	arena: vmem.Arena
	defer vmem.arena_destroy(&arena)

	context.allocator = vmem.arena_allocator(&arena)

	wrapping_paper_p1: int
	feet_of_ribbon_p2: int

	input := INPUT
	for line in strings.split_lines_iterator(&input) {
	    parts := strings.split(line, "x")
		l, _ := strconv.parse_int(parts[0])
		w, _ := strconv.parse_int(parts[1])
		h, _ := strconv.parse_int(parts[2])
		dimensions := [3]int{ l, w, h }
		slice.sort(dimensions[:])
		smallest_side := dimensions[0] * dimensions[1]
		wrapping_paper_p1 = wrapping_paper_p1 + box_surface_area(l, w, h) + smallest_side
		feet_of_ribbon_p2 = feet_of_ribbon_p2 + dimensions[0] * 2 + dimensions[1] * 2 + dimensions[0] * dimensions[1] * dimensions[2]
	}

	fmt.printf("Wrapping paper: %d\n", wrapping_paper_p1)
	fmt.printf("Feet of ribbon: %d\n", feet_of_ribbon_p2)
}
