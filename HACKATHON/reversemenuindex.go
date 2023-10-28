package piscine

func ReverseMenuIndex(menu []string) []string {
	var res []string
	// for i, j := 0, len(menu)-1; i <= j; i, j = i+1, j-1 {
	// 	menu[i], menu[j] = menu[j], menu[i]
	// }

	for i := len(menu) - 1; i >= 0; i-- {
		res = append(res, menu[i])
	}
	return res
}

//
