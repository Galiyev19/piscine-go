package piscine

func Sqrt(nb int) int {
	if nb == 0 || nb == 1 {
		return nb
	}

	x := nb

	for i := 0; i < nb; i++ {
		x = (x + nb/x) / 2
	}

	if x*x == nb {
		return x
	} else {
		return 0
	}
}
