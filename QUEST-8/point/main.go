package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func PrintNbrBase(nbr int, base string) string {
	if nbr < 0 {
		nbr = -nbr
	}

	length := len(base)
	reversStr := ""
	res := ""

	for nbr > 0 {
		remainder := nbr % length
		digit := rune(base[remainder])
		reversStr += string(digit)
		nbr /= length
	}

	for i := len(reversStr) - 1; i >= 0; i-- {
		res += string(reversStr[i])
	}

	return res
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	result := "x = " + PrintNbrBase(points.x, "0123456789") + ", y = " + PrintNbrBase(points.y, "0123456789") + "\n"

	for _, str := range result {
		z01.PrintRune(rune(str))
	}
}
