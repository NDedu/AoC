package main

import "core:fmt"
import vmem "core:mem/virtual"
import "core:strconv"
import "core:strings"

INPUT :: #load("../input/day6.txt", string)

p1_lights_logic :: proc(grid: ^[1000][1000]bool, str: string) {
	parts := strings.fields(str)

	start := strings.split(parts[len(parts) - 3], ",")
	end := strings.split(parts[len(parts) - 1], ",")

	x1, _ := strconv.parse_int(start[0])
	y1, _ := strconv.parse_int(start[1])
	x2, _ := strconv.parse_int(end[0])
	y2, _ := strconv.parse_int(end[1])

	for x in x1..=x2 {
		for y in y1..=y2 {
			if parts[0] == "turn" && parts[1] == "on" {
				grid[x][y] = true
			} else if parts[0] == "turn" && parts[1] == "off" {
				grid[x][y] = false
			} else if parts[0] == "toggle" && grid[x][y] == true {
				grid[x][y] = false
			} else if parts[0] == "toggle" && grid[x][y] == false {
				grid[x][y] = true
			}
		}
	}
}

p2_lights_logic :: proc(grid: ^[1000][1000]int , str: string) {
	parts := strings.fields(str)

	start := strings.split(parts[len(parts) - 3], ",")
	end := strings.split(parts[len(parts) - 1], ",")

	x1, _ := strconv.parse_int(start[0])
	y1, _ := strconv.parse_int(start[1])
	x2, _ := strconv.parse_int(end[0])
	y2, _ := strconv.parse_int(end[1])

	for x in x1..=x2 {
		for y in y1..=y2 {
			if parts[0] == "turn" && parts[1] == "on" {
				grid[x][y] += 1
			} else if parts[0] == "turn" && parts[1] == "off" && grid[x][y] > 0 {
				grid[x][y] -= 1
			} else if parts[0] == "toggle" {
				grid[x][y] += 2
			}
		}
	}
}

main :: proc() {
	arena: vmem.Arena
	defer vmem.arena_destroy(&arena)
	context.allocator = vmem.arena_allocator(&arena)

	grid_p1 := new([1000][1000]bool)
	grid_p2 := new([1000][1000]int)

	input := INPUT
	for line in strings.split_lines_iterator(&input) {
		p1_lights_logic(grid_p1, line)
		p2_lights_logic(grid_p2, line)
	}

	lit_p1: int
	total_brightness_p2: int
	for x := 0; x < 1000; x += 1 {
		for y := 0; y < 1000; y += 1 {
			if grid_p1[x][y] == true do lit_p1 += 1
			total_brightness_p2 += grid_p2[x][y]
		}
	}

	fmt.printf("Part 1 lit lights: %d\n", lit_p1)
	fmt.printf("Part 2 total brightness: %d\n", total_brightness_p2)
}
