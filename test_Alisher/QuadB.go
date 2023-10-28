package piscine

import (
	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	height := y
	width := x

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 || i == height-1 || j == 0 || j == width-1 {
				if j == 0 || j == width-1 {
					if i > 0 && i < height-1 {
						z01.PrintRune('*')
					} else if i == 0 && i <= height-1 {
						if j == 0 {
							z01.PrintRune('/')
						} else {
							z01.PrintRune('\\')
						}
					} else if i >= height-1 || i != 0 {
						if j == 0 {
							z01.PrintRune('\\')
						} else {
							z01.PrintRune('/')
						}
					}
				} else {
					z01.PrintRune('*')
				}
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
