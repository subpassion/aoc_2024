package main

import (
	utils "aoc_2024/aoc_utils"
	"fmt"
	"strings"
)

func is_inc(numbers []int64) bool {
	for i := 0; i < len(numbers)-1; i += 1 {
		diff := numbers[i+1] - numbers[i]
		if !(numbers[i] < numbers[i+1] && diff >= 1 && diff <= 3) {
			return false
		}
	}
	return true
}

func custom_append(slice1, slice2 []int64) []int64 {
	res := make([]int64, 0)
	res = append(res, slice1...)
	res = append(res, slice2...)
	return res
}

func is_inc_one_error(numbers []int64) bool {
	for i := 0; i < len(numbers); i += 1 {
		new_seq := custom_append(numbers[:i], numbers[i+1:])
		if is_inc(new_seq) {
			return true
		}
	}
	return false
}

func covert_to_array(input string) []int64 {
	numbers := strings.Split(input, " ")
	res := make([]int64, len(numbers))
	for i, number := range numbers {
		res[i] = utils.ToNumber(number)
	}
	return res
}

func part1(input string) int64 {
	lines := strings.Split(input, "\n")
	res := int64(0)
	for _, line := range lines {
		numbers := covert_to_array(line)
		if is_inc(numbers) || is_inc(utils.Reverse(numbers)) {
			res += 1
		}
	}
	return res
}

func part2(input string) int64 {
	lines := strings.Split(input, "\n")
	res := int64(0)
	for _, line := range lines {
		numbers := covert_to_array(line)
		if is_inc_one_error(numbers) || is_inc_one_error(utils.Reverse(numbers)) {
			res += 1
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
