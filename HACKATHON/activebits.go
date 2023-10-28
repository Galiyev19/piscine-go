package piscine

func ActiveBits(n int) int {
	var binary string

	for n > 0 {
		remainder := n % 2
		binary = string('0'+remainder) + binary
		n /= 2
	}
	count := 0
	for _, ch := range binary {
		if ch == '1' {
			count++
		}
	}
	return count
}
