package piscine

// a into c.
// c into d.
// d into b.
// b into a.

func Enigma(a ***int, b *int, c *******int, d ****int) {
	tempA := ***a
	tempB := *b
	tempC := *******c
	tempD := ****d

	*******c = tempA
	****d = tempC
	*b = tempD
	***a = tempB
}
