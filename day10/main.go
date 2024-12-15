package main

import (
	utils "aoc_2024/aoc_utils"
	"fmt"
	"strings"
)

func pares_file(input string) ([]utils.Point, utils.Grid) {
	trail_heads := make([]utils.Point, 0)
	lines := strings.Split(input, "\n")
	grid := utils.CreateGrid(int64(len(lines)), int64(len(lines[0])))

	for x, row := range lines {
		for y, ch := range row {
			grid[x][y] = utils.ToNumber(string(ch))
			if grid[x][y] == 0 {
				trail_heads = append(trail_heads, utils.Point{X: int64(x), Y: int64(y)})
			}
		}
	}
	return trail_heads, grid
}

func find_route(grid utils.Grid, x, y int64, visited [][]bool) int64 {
	if grid[x][y] == 9 {
		return 1
	}

	directions := []utils.Point{
		{-1, 0}, // move up
		{0, 1},  // move right
		{1, 0},  // move down
		{0, -1}} // move left
	res := int64(0)
	for _, dir := range directions {
		next_x, next_y := x+dir.X, y+dir.Y
		if grid.InRange(next_x, next_y) && grid[next_x][next_y]-grid[x][y] == 1 && !visited[next_x][next_y] {
			res += find_route(grid, next_x, next_y, visited)
			visited[next_x][next_y] = true
		}
	}
	return res
}

func find_route_not_visited(grid utils.Grid, x, y int64) int64 {
	if grid[x][y] == 9 {
		return 1
	}

	directions := []utils.Point{
		{-1, 0}, // move up
		{0, 1},  // move right
		{1, 0},  // move down
		{0, -1}} // move left
	res := int64(0)
	for _, dir := range directions {
		next_x, next_y := x+dir.X, y+dir.Y
		if grid.InRange(next_x, next_y) && grid[next_x][next_y]-grid[x][y] == 1 {
			res += find_route_not_visited(grid, next_x, next_y)
		}
	}
	return res
}

func part1(input string) int64 {
	heads, grid := pares_file(input)
	res := int64(0)
	for _, head := range heads {
		visited := make([][]bool, len(grid))
		for i := range visited {
			visited[i] = make([]bool, len(grid[i]))
			for j := range visited[i] {
				visited[i][j] = false
			}
		}
		res += find_route(grid, head.X, head.Y, visited)
	}

	return int64(res)
}

func part2(input string) int64 {
	heads, grid := pares_file(input)
	res := int64(0)
	for _, head := range heads {
		res += find_route_not_visited(grid, head.X, head.Y)
	}

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
