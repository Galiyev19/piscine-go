package main

import "fmt"

func main() {
	inputByte := byte(0b10101010) // Example byte

	result := ReverseBits(inputByte)

	fmt.Printf("Original Byte: %08b\n", inputByte)
	fmt.Printf("Reversed Byte : %08b\n", result)
}

func ReverseBits(oct byte) byte {
	var res byte

	for i := 0; i < 8; i++ {
		bit := oct & 1
		res = res << 1
		res = res | bit
		oct = oct >> 1
	}

	return res
}
