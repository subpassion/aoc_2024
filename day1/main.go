package main

import (
	"fmt"
	"slices"
	"strings"

	utils "aoc_2024/aoc_utils"
)

func parse_file1(input string) ([]int64, []int64) {
	left_list := make([]int64, 0)
	right_list := make([]int64, 0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		left_list = append(left_list, utils.ToNumber(numbers[0]))
		right_list = append(right_list, utils.ToNumber(numbers[1]))
	}

	slices.Sort(left_list)
	slices.Sort(right_list)
	return left_list, right_list
}

func part1(input string) int64 {
	left_distances, right_distances := parse_file1(input)
	var res int64
	for i := 0; i < len(left_distances); i += 1 {
		res += max(left_distances[i], right_distances[i]) - min(left_distances[i], right_distances[i])
	}
	return res
}

func parse_file2(input string) ([]int64, map[int64]int64) {
	left_list := make([]int64, 0)
	right_list_count := make(map[int64]int64)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		left_list = append(left_list, utils.ToNumber(numbers[0]))
		right_number := utils.ToNumber(numbers[1])
		right_list_count[right_number] += 1
	}

	return left_list, right_list_count
}

func part2(input string) int64 {
	left_distances, right_count := parse_file2(input)
	res := int64(0)
	for i := 0; i < len(left_distances); i += 1 {
		value, exists := right_count[left_distances[i]]
		if exists {
			res += value * left_distances[i]
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
