package main

import (
	utils "aoc_2024/aoc_utils"
	"fmt"
	"strings"
)

type Robot struct {
	pos utils.Point
	vel utils.Point
}

func parse_file(input string) []Robot {
	res := make([]Robot, 0)
	for _, line := range strings.Split(input, "\n") {
		line_data := strings.Split(line, " ")
		pos_str := strings.TrimPrefix(line_data[0], "p=")
		vel_str := strings.TrimPrefix(line_data[1], "v=")

		pos := strings.Split(pos_str, ",")
		vel := strings.Split(vel_str, ",")

		res = append(res, Robot{
			pos: utils.Point{
				X: utils.ToNumber(pos[1]),
				Y: utils.ToNumber(pos[0]),
			},
			vel: utils.Point{
				X: utils.ToNumber(vel[1]),
				Y: utils.ToNumber(vel[0]),
			},
		})
	}
	return res
}

func clamp(val, border int64) int64 {
	if val < 0 {
		return border + val
	} else if val >= border {
		return val - border
	}
	return val
}

func part1(input string) int64 {
	robots := parse_file(input)
	width := int64(101)
	height := int64(103)
	for i := 0; i < 1000; i++ {
		for r := 0; r < len(robots); r++ {
			robot := robots[r]
			robots[r].pos.X = clamp(robot.pos.X+robot.vel.X, height)
			robots[r].pos.Y = clamp(robot.pos.Y+robot.vel.Y, width)
		}
	}

	var q1, q2, q3, q4 int64
	for r := 0; r < len(robots); r++ {
		robot := robots[r]
		if robot.pos.X >= 0 && robot.pos.X < height/2 {
			if robot.pos.Y >= 0 && robot.pos.Y < width/2 {
				q1 += 1
			} else if robot.pos.Y > width/2 && robot.pos.Y < width {
				q2 += 1
			}
		}
		if robot.pos.X > height/2 && robot.pos.X < height {
			if robot.pos.Y >= 0 && robot.pos.Y < width/2 {
				q3 += 1
			} else if robot.pos.Y > width/2 && robot.pos.Y < width {
				q4 += 1
			}
		}
	}
	return q1 * q2 * q3 * q4
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
