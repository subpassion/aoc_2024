package main

import (
	utils "aoc_2024/aoc_utils"
	"fmt"
	"strings"
)

func apply_plus_mul(current_value int64, index int, numbers []int64, target int64) int64 {
	if current_value == target && index == len(numbers) {
		return target
	}

	if index < len(numbers) {
		next := current_value * numbers[index]
		res := apply_plus_mul(next, index+1, numbers, target)
		if res != 0 {
			return res
		}

		next = current_value + numbers[index]
		res = apply_plus_mul(next, index+1, numbers, target)
		if res != 0 {
			return res
		}
	}

	return 0
}

func part1(input string) int64 {
	res := int64(0)
	for _, line := range strings.Split(input, "\n") {
		line_parts := strings.Split(line, ":")
		target := utils.ToNumber(line_parts[0])
		numbers := make([]int64, 0)
		for _, number_str := range strings.Split(line_parts[1], " ") {
			if number_str != "" {
				numbers = append(numbers, utils.ToNumber(number_str))
			}
		}

		if apply_plus_mul(0, 0, numbers, target) == target {
			res += target
		}
	}
	return res
}

func apply_plus_mul_concat(current_value int64, index int, numbers []int64, target int64) int64 {
	if current_value == target && index == len(numbers) {
		return target
	}

	if index < len(numbers) {
		next := current_value * numbers[index]
		res := apply_plus_mul_concat(next, index+1, numbers, target)
		if res != 0 {
			return res
		}

		next = current_value + numbers[index]
		res = apply_plus_mul_concat(next, index+1, numbers, target)
		if res != 0 {
			return res
		}

		next = utils.ToNumber(fmt.Sprintf("%d%d", current_value, numbers[index]))
		res = apply_plus_mul_concat(next, index+1, numbers, target)
		if res != 0 {
			return res
		}
	}

	return 0
}

func part2(input string) int64 {
	res := int64(0)
	for _, line := range strings.Split(input, "\n") {
		line_parts := strings.Split(line, ":")
		target := utils.ToNumber(line_parts[0])
		numbers := make([]int64, 0)
		for _, number_str := range strings.Split(line_parts[1], " ") {
			if number_str != "" {
				numbers = append(numbers, utils.ToNumber(number_str))
			}
		}

		if apply_plus_mul_concat(0, 0, numbers, target) == target {
			res += target
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
