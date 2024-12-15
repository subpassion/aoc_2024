package main

import (
	utils "aoc_2024/aoc_utils"
	"fmt"
	"regexp"
	"strings"
)

func mul(input string) int64 {
	res := strings.TrimSuffix(strings.TrimPrefix(input, "mul("), ")")
	parts := strings.Split(res, ",")
	return utils.ToNumber(parts[0]) * utils.ToNumber(parts[1])
}

func part1(input string) int64 {
	res := int64(0)
	lines := strings.Split(input, "\n")

	r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	for _, line := range lines {
		multiplies := r.FindAllString(line, -1)
		local_res := int64(0)
		for _, mul_op := range multiplies {
			local_res += mul(mul_op)
		}
		res += local_res
	}
	return res
}

func part2(input string) int64 {
	res := int64(0)
	lines := strings.Split(input, "\n")

	r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	enabled := true
	for _, line := range lines {
		instructions := r.FindAllString(line, -1)
		local_res := int64(0)
		for _, instruction := range instructions {
			if strings.HasPrefix(instruction, "don't()") {
				enabled = false
			} else if strings.HasPrefix(instruction, "do()") {
				enabled = true
			} else if strings.HasPrefix(instruction, "mul(") && enabled {
				local_res += mul(instruction)
			}
		}
		res += local_res
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
