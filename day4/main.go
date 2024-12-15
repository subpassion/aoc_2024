package main

import (
	utils "aoc_2024/aoc_utils"
	"fmt"
	"strings"
)

const (
	XMAS = "XMAS"
	SAMX = "SAMX"
)

const (
	MAS = "MAS"
	SAM = "SAM"
)

func parse_file(input string) utils.StrGrid {
	lines := strings.Split(input, "\n")
	res := make(utils.StrGrid, uint64(len(lines)))

	for x, line := range lines {
		res[x] = line
	}
	return res
}

func count_number_of_xmas_or_samx(line string) uint64 {
	res := uint64(0)
	for i := 0; i <= len(line)-4; i += 1 {
		target := line[i : i+4]
		if target == XMAS || target == SAMX {
			res += 1
		}
	}
	return res
}

func search_diag(grid utils.StrGrid, curr_x, curr_y int) uint64 {
	res := uint64(0)
	left_diag := ""
	for x, y := curr_x, curr_y; x < min(curr_x+4, len(grid)) && y >= max(curr_y-4, 0); x, y = x+1, y-1 {
		left_diag += string(grid[x][y])
	}

	right_diag := ""
	for x, y := curr_x, curr_y; x < min(curr_x+4, len(grid)) && y < min(curr_y+4, len(grid[x])); x, y = x+1, y+1 {
		right_diag += string(grid[x][y])
	}
	// this is redundant
	res += count_number_of_xmas_or_samx(left_diag)
	res += count_number_of_xmas_or_samx(right_diag)
	return res
}

func part1(input string) uint64 {
	grid := parse_file(input)
	res := uint64(0)

	for x, row := range grid {
		res += count_number_of_xmas_or_samx(row)
		column := ""
		for c := 0; c < len(grid); c++ {
			column += string(grid[c][x])
		}
		for y := range row {
			res += search_diag(grid, x, y)
		}
		res += count_number_of_xmas_or_samx(column)
	}

	return res
}

func scan_right(grid utils.StrGrid, curr_x, curr_y int) bool {
	right_diag := ""
	for x, y := curr_x, curr_y; x < min(curr_x+3, len(grid)) && y < min(curr_y+3, len(grid[x])); x, y = x+1, y+1 {
		right_diag += string(grid[x][y])
	}

	return right_diag == SAM || right_diag == MAS
}

func scan_left(grid utils.StrGrid, curr_x, curr_y int) bool {
	left_diag := ""
	for x, y := curr_x, curr_y; x < min(curr_x+3, len(grid)) && y >= max(curr_y-3, 0); x, y = x+1, y-1 {
		left_diag += string(grid[x][y])
	}

	return left_diag == SAM || left_diag == MAS
}

func part2(input string) uint64 {
	grid := parse_file(input)
	res := uint64(0)

	for x, row := range grid {
		for y := range row {
			right := scan_right(grid, x, y)
			left := scan_left(grid, x, min(y+2, len(row)-1))
			if right && left {
				res += 1
			}
		}
	}

	return res
}

func solve() {
	var input = utils.ReadFile("input.txt")
	part1_res := part1(input)
	fmt.Printf("Part1: %d\n", part1_res)
	part2_res := part2(input)
	fmt.Printf("Part2: %d\n", part2_res)
}

func main() {
	solve()
}
