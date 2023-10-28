package piscine

func ComCheck(s []string) string {
	res := ""

	for _, ch := range s {
		if string(ch) == "01" || string(ch) == "galaxy" {
			res = "Alert!!!"
		}
	}
	return res
}
