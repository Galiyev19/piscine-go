package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	width := x
	height := y

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 || i == height-1 {
				if j == 0 || j == width-1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}

				if j == width-1 {
					z01.PrintRune('\n')
				}

			}

			if i != 0 && i != height-1 {
				if j == 0 || j == width-1 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}

				if j == width-1 {
					z01.PrintRune('\n')
				}
			}
		}
	}
}
