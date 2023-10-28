package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	runes1 := []rune(args[0])
	runes2 := []rune(args[1])
	count := 0
	res := ""
	// if ReWrite(args[0], args[1]) {
	for i := 0; i < len(runes1); i++ {
		for j := count; j < len(runes2); j++ {
			if runes1[i] == runes2[j] {
				res += string(runes1[i])
				j = len(runes2) - 1
			}
			count++
		}
	}
	if res == args[0] {
		fmt.Println(res)
	} else {
		return
	}
	// }
}

// func ReWrite(s1, s2 string) bool {
// 	index := 0

// 	for _, ch1 := range s1 {
// 		found := false

// 		for i := index; i < len(s2); i++ {
// 			if rune(s2[i]) == ch1 {
// 				index = i + 1
// 				found = true
// 				break
// 			}
// 		}

// 		if !found {
// 			return false
// 		}
// 	}

// 	return true
// }
