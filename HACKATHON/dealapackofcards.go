package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	for _, num := range deck {
		if num == 1 {
			fmt.Printf("Player 1: %v, %v, %v", deck[0], deck[1], deck[2])
			z01.PrintRune('\n')
		} else if num == 3 {
			fmt.Printf("Player 2: %v, %v, %v", deck[3], deck[4], deck[5])
			z01.PrintRune('\n')
		} else if num == 6 {
			fmt.Printf("Player 3: %v, %v, %v", deck[6], deck[7], deck[8])
			z01.PrintRune('\n')
		} else if num == 9 {
			fmt.Printf("Player 4: %v, %v, %v", deck[9], deck[10], deck[11])
			z01.PrintRune('\n')
		}
	}
}
