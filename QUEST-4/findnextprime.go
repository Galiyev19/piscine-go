package piscine

func FindNextPrime(nb int) int {
	prime := IsPrime(nb)
	if prime {
		return nb
	}

	next := nb + 1

	for {
		if IsPrime(next) {
			return next
		}
		next++
	}
}
