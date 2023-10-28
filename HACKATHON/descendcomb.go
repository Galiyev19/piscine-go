package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 57; i >= 48; i-- {
		for j := 57; j >= 48; j-- {

			m := j - 1

			for k := i; k >= 48; k-- {
				for ; m >= 48; m-- {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					z01.PrintRune(' ')
					z01.PrintRune(rune(k))
					z01.PrintRune(rune(m))
					if rune(i) != '0' || rune(j) != '1' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				m = '9'
			}

		}
	}
}
