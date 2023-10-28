package piscine

func Swap(a *int, b *int) {
	temp := *a
	temp1 := *b
	*a = temp1
	*b = temp
}
