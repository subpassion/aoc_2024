package main

import (
	utils "aoc_2024/aoc_utils"
	"fmt"
	"math"
	"slices"
	"strings"
)

type Rules map[int64][]int64

func parse_file(input string) (Rules, utils.Grid) {
	rules := make(Rules)
	instructions := make(utils.Grid, 0, 0)

	parse_rules := true
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			parse_rules = false
		} else if parse_rules {
			rules_str := strings.Split(line, "|")
			rule_one, rule_two := utils.ToNumber(rules_str[0]), utils.ToNumber(rules_str[1])
			rules[rule_one] = append(rules[rule_one], rule_two)
		} else {
			instructions_str := strings.Split(line, ",")
			local_instruction := make([]int64, len(instructions_str))
			for i, ins_str := range instructions_str {
				local_instruction[i] = utils.ToNumber(ins_str)
			}
			instructions = append(instructions, local_instruction)
		}
	}
	return rules, instructions
}

func is_correct(rules Rules, page []int64) (bool, int) {
	current_rule := rules[page[0]]
	for j := 1; j < len(page); j++ {
		if !slices.Contains(current_rule, page[j]) {
			return false, j
		}
		current_rule = rules[page[j]]
	}
	return true, math.MaxInt64
}

func part1(input string) int64 {
	ordering_rules, pages := parse_file(input)
	res := int64(0)
	for _, page := range pages {
		if correct, _ := is_correct(ordering_rules, page); correct {
			res += page[len(page)/2]
		}
	}
	return res
}

func do_the_correction(rules Rules, page []int64) int64 {
	for {
		if correct, curr_index := is_correct(rules, page); correct {
			return page[len(page)/2]
		} else {
			page[curr_index-1], page[curr_index] = page[curr_index], page[curr_index-1]
		}
	}
}

func part2(input string) int64 {
	ordering_rules, pages := parse_file(input)
	res := int64(0)
	for _, page := range pages {
		if correct, _ := is_correct(ordering_rules, page); !correct {
			res += do_the_correction(ordering_rules, page)
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
