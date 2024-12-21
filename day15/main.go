package main

import (
	utils "aoc_2024/aoc_utils"
	"fmt"
	"slices"
	"strings"
)

func parse_file(input string) (utils.Grid, utils.StrGrid, utils.Point) {
	grid_and_moves := strings.Split(input, "\n\n")
	grid := make(utils.Grid, 0)
	moves := make(utils.StrGrid, 0)
	start := utils.Point{}
	for x, line := range strings.Split(grid_and_moves[0], "\n") {
		row := make([]int64, 0)
		for y, val := range line {
			row = append(row, int64(val))
			if val == '@' {
				start = utils.Point{X: int64(x), Y: int64(y)}
			}
		}
		grid = append(grid, row)
	}

	for _, line := range strings.Split(grid_and_moves[1], "\n") {
		moves = append(moves, line)
	}

	return grid, moves, start
}

func find_vacant_spot(input []int64, current int64, move_to_the_right bool) int64 {
	if move_to_the_right {
		for x := current; x < int64(len(input)); x++ {
			if input[x] == int64('#') {
				break
			}
			if input[x] == int64('.') {
				return x
			}
		}
	} else {
		for x := current; x >= 0; x-- {
			if input[x] == int64('#') {
				break
			}
			if input[x] == int64('.') {
				return x
			}
		}
	}
	return -1
}

func move_through(grid utils.Grid, start utils.Point, dir string) utils.Point {
	directions := map[string]utils.Point{
		"^": {X: -1, Y: 0},
		">": {X: 0, Y: 1},
		"v": {X: 1, Y: 0},
		"<": {X: 0, Y: -1},
	}
	current_dir := directions[dir]
	next := utils.Point{X: start.X + current_dir.X, Y: start.Y + current_dir.Y}
	if grid.InRange(next.X, next.Y) {
		path := make([]int64, 0)
		path_start := int64(-1)
		if dir == "v" || dir == "^" {
			path_start = start.X
			for x := 0; x < len(grid); x++ {
				path = append(path, grid[x][start.Y])
			}
		} else {
			path_start = start.Y
			path = grid[start.X]
		}

		ascending := dir == ">" || dir == "v"
		vacant_spot := find_vacant_spot(path, int64(path_start), ascending)
		if vacant_spot != -1 {
			if ascending {
				for y := vacant_spot; y > int64(path_start); y-- {
					path[y], path[y-1] = path[y-1], path[y]
				}
			} else {
				for y := vacant_spot; y < int64(path_start); y++ {
					path[y], path[y+1] = path[y+1], path[y]
				}
			}

			if dir == ">" || dir == "<" {
				copy(grid[start.X], path)
				return utils.Point{X: start.X, Y: int64(slices.Index(path, int64('@')))}
			} else {
				for x := 0; x < len(path); x++ {
					grid[x][start.Y] = path[x]
				}
				return utils.Point{X: int64(slices.Index(path, int64('@'))), Y: start.Y}
			}
		}
	}
	return start
}
func part1(input string) int64 {
	grid, moves, start := parse_file(input)
	for _, move := range moves {
		for _, dir := range move {
			start = move_through(grid, start, string(dir))
		}
	}

	grid.Print()

	res := int64(0)
	for x, row := range grid {
		for y := range row {
			if grid[x][y] == int64('O') {
				res += int64(100*x + y)
			}
		}
	}
	return res
}

func part2(input string) int64 {
	return 0
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
