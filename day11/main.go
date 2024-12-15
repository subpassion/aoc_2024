package main

import (
	utils "aoc_2024/aoc_utils"

	"fmt"
	"math"
	"strings"
)

func number_of_digits(number int64) int64 {
	return int64(math.Log10(float64(number)) + 1)
}

func number_parts(number int64) (int64, int64) {
	n_digits := number_of_digits(number)
	scalar := (int64(math.Pow10(int(n_digits) / 2)))
	first_part := int64(number / scalar)
	second_part := number % (scalar)
	return first_part, second_part
}

func process_numbers(numbers []int64) []int64 {
	res := make([]int64, 0)
	for _, number := range numbers {
		if number == 0 {
			res = append(res, 1)
		} else if number_of_digits(number)%2 == 0 {
			f, s := number_parts(number)
			res = append(res, f, s)
		} else {
			res = append(res, number*2024)
		}
	}
	return res
}

func parse_file(input string) []int64 {
	res := make([]int64, 0)
	for _, number := range strings.Split(input, " ") {
		res = append(res, utils.ToNumber(number))
	}
	return res
}

func part1(input string) int64 {
	numbers := parse_file(input)
	for i := 0; i < 13; i++ {
		fmt.Println(numbers, len(numbers))
		numbers = process_numbers(numbers)
	}

	return int64(len(numbers))
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
