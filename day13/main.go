package main

import (
	utils "aoc_2024/aoc_utils"
	"fmt"
	"math"
	"strings"
)

func cramers(a1, a2, b1, b2 float64, c1, c2 float64) (float64, float64) {
	d := a1*b2 - b1*a2
	if d == 0 {
		return 0, 0
	}
	dx := c1*b2 - c2*b1
	dy := a1*c2 - a2*c1

	return dx / d, dy / d
}

type Equation struct {
	a1, a2 float64
	b1, b2 float64
	c1, c2 float64
}

func parse_file(input string) []Equation {
	lines := strings.Split(input, "\n")
	res := make([]Equation, 0)

	var eq = Equation{}
	for _, line := range lines {
		if line == "" {
			res = append(res, eq)
			eq = Equation{}
			continue
		}

		numbers_str := strings.Split(line[max(0, strings.Index(line, ": "))+1:], ",")
		if strings.HasPrefix(line, "Button A:") {
			eq.a1 = utils.ToFloatNumber(numbers_str[0][2:])
			eq.a2 = utils.ToFloatNumber(numbers_str[1][2:])
		} else if strings.HasPrefix(line, "Button B:") {
			eq.b1 = utils.ToFloatNumber(numbers_str[0][2:])
			eq.b2 = utils.ToFloatNumber(numbers_str[1][2:])
		} else if strings.HasPrefix(line, "Prize: ") {
			eq.c1 = utils.ToFloatNumber(numbers_str[0][3:])
			eq.c2 = utils.ToFloatNumber(numbers_str[1][3:])
		}
	}
	res = append(res, eq)
	return res
}

func part1(input string) float64 {
	eqs := parse_file(input)
	tolerance := 0.00001
	res := float64(0)
	for _, eq := range eqs {
		a, b := cramers(eq.a1, eq.a2, eq.b1, eq.b2, eq.c1, eq.c2)
		if math.Abs(a-math.Round(a)) < tolerance && math.Abs(b-math.Round(b)) < tolerance {
			res += (a*3 + b)
		}
	}
	return res
}

func part2(input string) float64 {
	eqs := parse_file(input)
	tolerance := 0.00001
	scale := float64(10000000000000)
	res := float64(0)
	for _, eq := range eqs {
		a, b := cramers(eq.a1, eq.a2, eq.b1, eq.b2, eq.c1+scale, eq.c2+scale)
		if math.Abs(a-math.Round(a)) < tolerance && math.Abs(b-math.Round(b)) < tolerance {
			res += (a*3 + b)
		}
	}
	return res
}

func solve() {
	var input = utils.ReadFile("input.txt")
	part1_res := part1(input)
	fmt.Printf("Part1: %f\n", part1_res)
	part2_res := part2(input)
	fmt.Printf("Part2: %f\n", part2_res)
}

func main() {
	solve()
}
