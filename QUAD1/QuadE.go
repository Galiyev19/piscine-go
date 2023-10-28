package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	width := x
	height := y

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 {
				if j == 0 {
					z01.PrintRune('A')
				} else if j == width-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}

				if j == width-1 {
					z01.PrintRune('\n')
				}
			}

			if i != 0 && i != height-1 {
				if j == 0 || j == width-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}

				if j == width-1 {
					z01.PrintRune('\n')
				}
			}

			if i == height-1 && i != 0 {
				if j == 0 {
					z01.PrintRune('C')
				} else if j == width-1 {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}

				if j == width-1 {
					z01.PrintRune('\n')
				}
			}
		}
	}
}
