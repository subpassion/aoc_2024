package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Point struct {
	x int64
	y int64
}

type Grid [][]int64
type StrGrid []string

func (grid Grid) InRange(x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func (grid StrGrid) InRange(x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func (grid Grid) Print() {
	for _, row := range grid {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Println()
	}
}

func (grid StrGrid) Print() {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func ToNumber(number string) int64 {
	processed_number := strings.TrimSpace(number)
	res, err := strconv.ParseInt(processed_number, 10, 0)
	Report(err)
	return res
}

func IsNumber(input string) bool {
	number := strings.TrimSpace(input)
	for _, ch := range number {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

func CreateGrid(n_rows, n_cols int64) Grid {
	array := make([][]int64, n_rows)
	for i := range array {
		array[i] = make([]int64, n_cols)
	}
	return array
}

func Report(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(file_path string) string {
	content, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func Count[T comparable](slice []T, element T) int64 {
	count := int64(0)
	for _, v := range slice {
		if v == element {
			count++
		}
	}
	return count
}

func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}

func Reverse[T any](slice []T) []T {
	reversed := make([]T, len(slice))
	copy(reversed, slice)

	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return reversed
}
