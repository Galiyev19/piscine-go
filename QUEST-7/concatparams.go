package piscine

func ConcatParams(args []string) string {
	concatStr := ""

	for idx, str := range args {
		if !(idx == len(args)-1) {
			concatStr += str + "\n"
		} else {
			concatStr += str
		}
	}

	return concatStr
}
