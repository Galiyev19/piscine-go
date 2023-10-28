package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		for _, s := range os.Args[1:] {
			file, err := os.Open(s)
			if err != nil {
				for _, e := range "ERROR: " + err.Error() {
					z01.PrintRune(e)
				}
				z01.PrintRune('\n')
				os.Exit(1)
				break
			} else {
				data, err := ioutil.ReadAll(file)
				if err != nil {
					for _, e := range "ERROR: " + err.Error() {
						z01.PrintRune(e)
					}
					z01.PrintRune('\n')
					os.Exit(1)
					break
				} else {
					for _, text := range string(data) {
						z01.PrintRune(text)
					}
				}
			}
		}
	}
}
