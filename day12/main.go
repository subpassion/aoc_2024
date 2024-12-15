package main

import (
	utils "aoc_2024/aoc_utils"
	"fmt"
	"strings"
)

func parse_file(input string) utils.StrGrid {
	lines := strings.Split(input, "\n")
	res := make(utils.StrGrid, 0)

	for _, line := range lines {
		res = append(res, line)
	}
	return res
}

func compute(current byte, grid utils.StrGrid, x, y int64, visited [][]bool) (int64, int64) {
	directions := []utils.Point{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1}}

	res := int64(0)
	per := int64(4)
	for _, dir := range directions {
		next_x, next_y := x+dir.X, y+dir.Y
		if grid.InRange(next_x, next_y) && grid[next_x][next_y] == current {
			if !visited[next_x][next_y] {
				visited[next_x][next_y] = true
				a, p := compute(current, grid, next_x, next_y, visited)
				res, per = res+1+a, per-1+p
			} else {
				per -= 1
			}
		}
	}
	return res, per
}

func part1(input string) int64 {
	garden := parse_file(input)

	visited := make([][]bool, len(garden))
	for i := range visited {
		visited[i] = make([]bool, len(garden[i]))
		for j := range visited[i] {
			visited[i][j] = false
		}
	}

	res := int64(0)
	for i, row := range garden {
		for j, current := range row {
			if !visited[i][j] {
				visited[i][j] = true
				a, p := compute(byte(current), garden, int64(i), int64(j), visited)
				res += ((a + 1) * p)
			}
		}
	}
	return int64(res)
}

func part2(input string) int64 {
	res := 0
	return int64(res)
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
