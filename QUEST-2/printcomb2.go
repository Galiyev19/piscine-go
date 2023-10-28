package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 48; i < 58; i++ {
		for j := 48; j < 58; j++ {

			m := j + 1

			for k := i; k < 58; k++ {
				for ; m < 58; m++ {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					z01.PrintRune(' ')
					z01.PrintRune(rune(k))
					z01.PrintRune(rune(m))
					if rune(i) != '9' || rune(j) != '8' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				m = '0'
			}

		}
	}
	z01.PrintRune('\n')
}
