package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	stringChangeble := []byte(s)
	// fmt.Println(stringChangeble)
	for i := 0; i < len(stringChangeble); i++ {
		z01.PrintRune(rune(stringChangeble[i]))
	}
}
