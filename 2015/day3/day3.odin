package main

import "core:fmt"
import vmem "core:mem/virtual"
import "core:strings"

INPUT :: #load("../input/day3.txt", string)

main :: proc() {
	arena: vmem.Arena
	defer vmem.arena_destroy(&arena)
	context.allocator = vmem.arena_allocator(&arena)

	input := INPUT
	for line in strings.split_lines_iterator(&input) {

		seen_p1 := make(map[[2]int]struct{})
		seen_p2 := make(map[[2]int]struct{})

		x_y_coordinates_p1 := [2]int{0, 0}
		seen_p1[x_y_coordinates_p1] = {}

		x_y_coordinates_p2 := [2]int{0, 0}
		seen_p2[x_y_coordinates_p2] = {}
		x_y_coordinates_robot_p2 := [2]int{0, 0}

	    for char, i in line {
			switch char {
				case '>': {
					x_y_coordinates_p1[0] += 1
					if i % 2 != 0 do x_y_coordinates_p2[0] += 1
					if i % 2 == 0 do x_y_coordinates_robot_p2[0] += 1
				}
				case '<': {
					x_y_coordinates_p1[0] -= 1
					if i % 2 != 0 do x_y_coordinates_p2[0] -= 1
					if i % 2 == 0 do x_y_coordinates_robot_p2[0] -= 1
				}
				case '^': {
					x_y_coordinates_p1[1] += 1
					if i % 2 != 0 do x_y_coordinates_p2[1] += 1
					if i % 2 == 0 do x_y_coordinates_robot_p2[1] += 1
				}
				case 'v': {
					x_y_coordinates_p1[1] -= 1
					if i % 2 != 0 do x_y_coordinates_p2[1] -= 1
					if i % 2 == 0 do x_y_coordinates_robot_p2[1] -= 1
				}
				case: fmt.printf("No match")
			}
			seen_p1[x_y_coordinates_p1] = {}
			seen_p2[x_y_coordinates_p2] = {}
			seen_p2[x_y_coordinates_robot_p2] = {}
		}
		fmt.printf("Santa at least one preset: %d\n", len(seen_p1))
		fmt.printf("Santa and Robot at least one preset: %d\n", len(seen_p2))
	}
}
