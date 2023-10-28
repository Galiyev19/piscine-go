package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func Chunk(slice []int, size int) {
	if size == 0 {
		z01.PrintRune('\n')
		return
	}

	length := len(slice)
	var result [][]int

	for i := 0; i < length; i += size {
		end := i + size
		if end > length {
			end = length
		}
		result = append(result, slice[i:end])
	}

	fmt.Print("[")
	for i, subSlice := range result {
		fmt.Print(subSlice)
		if i < len(result)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}

func getDigits(n int) []rune {
	var digits []rune

	if n == 0 {
		return []rune{'0'}
	}

	for n > 0 {
		digit := '0' + rune(n%10)
		digits = append([]rune{digit}, digits...)
		n /= 10
	}

	return digits
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
