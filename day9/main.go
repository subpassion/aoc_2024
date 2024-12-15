package main

import (
	utils "aoc_2024/aoc_utils"
	"fmt"
)

const (
	FREE int = -1
)

func parse_file(input string) []int {
	res := make([]int, 0)
	free_block := false
	id := 0
	for _, ch := range input {
		for j := 0; j < int(utils.ToNumber(string(ch))); j++ {
			if free_block {
				res = append(res, FREE)
			} else {
				res = append(res, id)
			}
		}
		if !free_block {
			id++
		}
		free_block = !free_block
	}
	return res
}

func find_last_not_free(elements []int, end int) (int, bool) {
	for i := len(elements) - 1; i >= end; i-- {
		if elements[i] != FREE {
			return i, true
		}
	}
	return -1, false
}

func part1(input string) uint64 {
	elements := parse_file(input)
	for i, current_element := range elements {
		if current_element == FREE {
			last, exists := find_last_not_free(elements, i)
			if !exists {
				break
			}

			elements[i], elements[last] = elements[last], elements[i]
		}
	}

	res := uint64(0)
	for pos, elem := range elements {
		if elem == FREE {
			break
		}
		res += uint64(pos * elem)
	}

	return res
}

func find_free_fitting_spot(elements []int, len, end int) (int, bool) {
	current_free_block_len := 0
	for i := 0; i < end; i++ {
		if current_free_block_len == len {
			return i - len, true
		}

		if elements[i] == FREE {
			current_free_block_len++
		} else {
			current_free_block_len = 0
		}
	}
	return -1, false
}

func get_current_len(elements []int, end int) (int, int) {
	len := 0
	for i := end; i >= 0; i-- {
		if elements[end] == elements[i] {
			len++
		} else {
			break
		}
	}
	return end - len + 1, len
}

func part2(input string) uint64 {
	elements := parse_file(input)
	for i := len(elements) - 1; i >= 0; i-- {
		if elements[i] != FREE {
			_, block_len := get_current_len(elements, i)
			free_spot, exits := find_free_fitting_spot(elements, block_len, i)
			if exits {
				for s := 0; s < block_len; s++ {
					elements[free_spot+s], elements[i-s] = elements[i-s], elements[free_spot+s]
				}
			}
			i -= (block_len - 1)
		}
	}
	res := uint64(0)
	for pos, elem := range elements {
		if elem == FREE {
			continue
		}
		res += uint64(pos * elem)
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
