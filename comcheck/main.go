package main

import (
	"fmt"
	"os"
)

// func ComCheck(s []string) string {
// 	res := ""

// 	for _, ch := range s {
// 		if string(ch) == "01" || string(ch) == "galaxy" {
// 			res = "Alert!!!"
// 		}
// 	}
// 	return res
// }

func main() {
	args := os.Args[1:]

	for _, ch := range args {
		if string(ch) == "01" || string(ch) == "galaxy" || string(ch) == "or" {
			fmt.Println("Alert!!!")
			break

		}
	}
	return
}
