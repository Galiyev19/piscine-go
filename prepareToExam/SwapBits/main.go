package main

import "fmt"

func main() {
	inputByte := byte(0b00100110) // Example byte

	result := SwapBits(inputByte)

	fmt.Printf("Original Byte: %08b\n", inputByte)
	fmt.Printf("SwapBits Byte : %08b\n", result)
}

func SwapBits(octet byte) byte {
	return (octet >> 4) | (octet << 4)
}
