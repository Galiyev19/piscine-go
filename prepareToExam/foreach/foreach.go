package main

import piscine "piscine/QUEST-2"

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		f(a[i])
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(piscine.PrintNbr, a)
}
