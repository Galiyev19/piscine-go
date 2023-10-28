package piscine

func UltimateDivMod(a *int, b *int) {
	res := *a / *b
	res2 := *a % *b
	*a = res
	*b = res2
}
