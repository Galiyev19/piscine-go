package main

import "fmt"

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s string, sep string) []string {
	startIndex := 0
	var res []string
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			res = append(res, s[startIndex:i])
			startIndex = i + len(sep)
			i = startIndex - 1
		}
	}
	res = append(res, s[startIndex:])
	return res
}
