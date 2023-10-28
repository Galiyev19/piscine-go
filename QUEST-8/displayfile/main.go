package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]

	// fmt.Println(args[0])

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}

	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	content, err := ioutil.ReadFile(args[0])
	check(err)
	fmt.Printf("%s", content)

	// fmt.Println(content)
}
