package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	numArgs := len(os.Args) - 1

	// numbers := ""
	// var str []string

	// fmt.Println(numArgs)

	if numArgs == 0 {
		fmt.Printf(`--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.`)
		fmt.Println()
		return
	}

	if args[0] == "-h" || args[0] == "--help" || args[0] == "-i" || args[0] == "--insert" {
		fmt.Printf(`--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.`)
		fmt.Println()
	} else {

		flagInsert := false
		flagOrder := false
		index := 0

		var res []string

		for _, arg := range args[0] {
			if arg == '=' || arg == '-' {
				flagInsert = true
			}
		}

		if numArgs > 1 {
			for _, str := range args {
				if str == "--order" || str == "-o" {
					flagOrder = true
				}
			}
		}

		if flagInsert && flagOrder {

			for i, ch := range args[0] {
				if ch == '=' {
					index = i
					for i := index + 1; i < len(args[0]); i++ {
						res = append(res, string(args[0][i]))
					}
				}
			}

			if numArgs > 2 {
				for _, ch := range args[2] {
					res = append(res, string(ch))
				}
			} else {
				for _, ch := range args[1] {
					res = append(res, string(ch))
				}
			}
		} else {
			if numArgs == 1 {
				for _, ch := range args[0] {
					res = append(res, string(ch))
				}
			} else if numArgs == 2 {
				for _, ch := range args[1] {
					res = append(res, string(ch))
				}
			}

			for i, ch := range args[0] {
				if ch == '=' {
					index = i
				}
			}
			for i := index + 1; i < len(args[0]); i++ {
				res = append(res, string(args[0][i]))
			}
		}

		if flagOrder {
			for i := 0; i < len(res); i++ {
				for j := 0; j < len(res)-i-1; j++ {
					if res[j] > res[j+1] {
						res[j], res[j+1] = res[j+1], res[j]
					}
				}
			}
		}

		// fmt.Println(flagInsert)
		// fmt.Println(flagOrder)
		temp := ""
		for i := 0; i < len(res); i++ {
			temp += res[i]
		}
		fmt.Println(temp)
	}
}
