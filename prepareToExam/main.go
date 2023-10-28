package main

import (
	"fmt"
)

func main() {
	// args := os.Args[1:]
	// fmt.Println("PARAMS")
	// fParam := DisplayFirstParam(args)
	// lParam := DisplayLastParam(args)
	// fmt.Println(fParam)
	// fmt.Println(lParam)
	// fmt.Println("STRING LENGTH")
	// l := StrLen("Hello World!")
	// fmt.Println(l)

	// fmt.Println("String Reverse")
	// s := "Hello World!"
	// s = StrRev(s)
	// fmt.Println(s)

	// fmt.Println("WDMATCH")
	// match := wdMatch(args[0], args[1])
	// fmt.Println(match)

	// fmt.Println("ISPOWER2")
	// fmt.Println(IsPower2(2))

	// fmt.Println("MAX")
	// a := []int{23, 123, 1, 11, 55, 93}
	// max := Max(a)
	// fmt.Println(max)

	// fmt.Println("REDUCEINT")
	// mul := func(acc int, cur int) int {
	// 	return acc * cur
	// }
	// sum := func(acc int, cur int) int {
	// 	return acc + cur
	// }
	// div := func(acc int, cur int) int {
	// 	return acc / cur
	// }
	// as := []int{500, 2}
	// ReduceInt(as, mul)
	// ReduceInt(as, sum)
	// ReduceInt(as, div)

	// fmt.Println(ParamCount(args))

	// fmt.Println(SplitStr(args[0]))

	// result := Rot14("Hello! How are You?")

	// for _, r := range result {
	// 	z01.PrintRune(r)
	// }
	// z01.PrintRune('\n')

	// fmt.Println(switchCase(args))s
	// fmt.Println(AlphaMirror(args[0]))

	// fmt.Println(Atoi("12345"))
	// fmt.Println(Atoi("0000000012345"))
	// fmt.Println(Atoi("012 345"))
	// fmt.Println(Atoi("Hello World!"))
	// fmt.Println(Atoi("+1234"))
	// fmt.Println(Atoi("-1234"))
	// fmt.Println(Atoi("++1234"))
	// fmt.Println(Atoi("--1234"))

	// fmt.Println(WDMATCH(args[0], args[1]))
	// fmt.Println(LASTWORD(args[0]))
	// fmt.Println(RomanNumber(123))
	fmt.Println(AtoiBase("125", "0123456789"))
	// fmt.Println(AtoiBase("1111101", "01"))
	// fmt.Println(AtoiBase("bbbbbab", "-ab"))

	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

func Capitalize(s string) string {
	nextToUpper := false
	result := make([]rune, len(s))
	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
			if nextToUpper {
				result[i] = ch - 'a' - 'A'
				nextToUpper = false
			} else {
				result[i] = ch + 'a' - 'A'
			}
		} else {
			result[i] = ch
			nextToUpper = true
		}
	}
	return string(result)
}

func AtoiBase(num, base string) int {
	res := 0
	index := 0

	if !isValid(base) {
		return 0
	}

	for _, ch := range num {
		digit := string(ch)

		for i, char := range base {
			if string(char) == digit {
				index = i
				break
			}
		}
		res = res*len(base) + index
	}

	return res
}

func isValid(base string) bool {
	if len(base) < 2 {
		return false
	}

	uniqueChars := make(map[rune]bool)

	for _, ch := range base {
		if uniqueChars[ch] || ch == '-' || ch == '+' {
			return false
		}
		uniqueChars[ch] = true
	}
	return true
}

func RomanNumber(num int) string {
	if num <= 0 || num >= 4000 {
		return "ERROR: cannot convert to roman digit"
	}

	res := ""
	romanSymbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i := 0; i < len(romanSymbols); i++ {
		for num >= romanValues[i] {
			res += romanSymbols[i]
			num -= romanValues[i]
		}
	}

	return res
}

func WDMATCH(s1, s2 string) string {
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	count := 0
	var res string
	if ReWrite(s1, s2) {
		for i := 0; i < len(runes1); i++ {
			for j := count; j < len(runes2); j++ {
				if runes1[i] == runes2[j] {
					res += string(runes1[i])
					j = len(runes2) - 1
				}
				count++
			}
		}
	} else {
		return ""
	}

	return res
}

func ReWrite(s1, s2 string) bool {
	index := 0

	for _, ch1 := range s1 {
		found := false

		for i := index; i < len(s2); i++ {
			if rune(s2[i]) == ch1 {
				index = i + 1
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func TRIM(s string) string {
	start, end := 0, len(s)

	// Trim spaces from the beginning
	for start < end && s[start] == ' ' {
		start++
	}

	// Trim spaces from the end
	for start < end && s[end-1] == ' ' {
		end--
	}

	return s[start:end]
}

func LASTWORD(s string) string {
	var strArr []string

	trimStr := TRIM(s)
	startIndex := 0
	for i := 0; i < len(trimStr); i++ {
		if trimStr[i] == ' ' {
			strArr = append(strArr, trimStr[startIndex:i])
			startIndex = i + 1
		}
	}
	strArr = append(strArr, trimStr[startIndex:])
	return strArr[len(strArr)-1]
}

func DisplayFirstParam(args []string) string {
	return args[0]
}

func DisplayLastParam(args []string) string {
	return args[len(args)-1]
}

func StrLen(str string) int {
	count := 0
	for _, str := range str {
		str++
		count++
	}

	return count
}

func StrRev(s string) string {
	strRev := ""

	for i := len(s) - 1; i >= 0; i-- {
		strRev += string(s[i])
	}

	return strRev
}

func wdMatch(s1, s2 string) string {
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	var res string
	count := 0
	for i := 0; i < len(runes1); i++ {
		for j := count; j < len(runes2); j++ {
			if runes1[i] == runes2[j] {
				res += string(runes1[i])
				j = len(runes2) - 1
			}
			count++
		}
	}
	return res
}

func IsPower2(num int) bool {
	return num > 0 && (num&(num-1)) == 0
}

func Max(arr []int) int {
	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func ReduceInt(a []int, f func(int, int) int) {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res = f(res, a[i])
	}
	fmt.Println(res)
}

func ParamCount(arr []string) int {
	return len(arr)
}

func Trim(s string) string {
	start, end := 0, len(s)

	// Trim spaces from the beginning
	for start < end && s[start] == ' ' {
		start++
	}

	// Trim spaces from the end
	for start < end && s[end-1] == ' ' {
		end--
	}

	return s[start:end]
}

func SplitStr(s string) string {
	var strArr []string

	trimStr := Trim(s)
	startIndex := 0
	for i := 0; i < len(trimStr); i++ {
		if trimStr[i] == ' ' {
			strArr = append(strArr, trimStr[startIndex:i])
			startIndex = i + 1
		}
	}
	strArr = append(strArr, trimStr[startIndex:])
	return strArr[len(strArr)-1]
}

func Rot14(s string) string {
	str := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			if ch >= 'm' {
				ch -= 12
			} else {
				ch += 14
			}
		} else if ch >= 'A' && ch <= 'z' {
			if ch >= 'M' {
				ch -= 12
			} else {
				ch += 14
			}
		}
		str += string(ch)
	}

	return str
}

func switchCase(s []string) string {
	res := ""

	for _, arg := range s {
		for _, ch := range arg {
			if ch >= 'A' && ch <= 'Z' {
				ch = rune(ch + 32)
			} else if ch >= 'a' && ch <= 'z' {
				ch = rune(ch - 32)
			}
			res += string(ch)
		}
	}
	return res
}

func AlphaMirror(s string) string {
	res := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			res += string('z' - ch + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			res += string('Z' - ch + 'A')
		} else {
			res += string(ch)
		}
	}
	return res
}

func Atoi(s string) int {
	res := 0
	sign := 1
	count := 0

	for _, ch := range s {
		if ch == '+' || ch == '-' {
			count++
		}

		if count > 1 {
			return 0
		}

		if ch == '+' || ch == '-' {
			if ch == '-' {
				sign = -1
			}
		} else if ch >= '0' && ch <= '9' {
			digit := int(ch - '0')
			res = res*10 + digit
		} else {
			return 0
		}
	}

	return res * sign
}

// func PrintBits(num int) {
// 	for i := 7; i >= 0; i-- {
// 		bit := (num >> uint(i)) & 1
// 		char := '0'
// 		if bit == 1 {
// 			char = '1'
// 		}
// 		z01.PrintRune(char)
// 	}
// }

// func SwapBits(oct byte) byte {
// 	A := oct / byte(16)
// 	B := oct % byte(16)
// 	C := B*byte(16) + A
// 	return C
// }

// func ReverseBits2(oct byte) byte {
// 	var result byte
// 	for i := 0; i < 8; i++ {
// 		bit := oct & 1
// 		result = result << 1
// 		result = result | bit
// 		oct = oct >> 1
// 	}
// 	return result
// }

// func RevBits(oct byte) byte {
// 	var result byte
// 	return result<<1 | oct>>1
// }

// func ReverseBits(oct byte) byte {
// 	var result byte

// 	for i := 0; i < 8; i++ {
// 		result = (result << 1) | (oct & 1)

// 		oct >>= 1
// 	}
// 	return result
// }

// // DISPLAY FIRST PARAM
// func DisplayFirstParam(s []string) {
// 	if len(s) == 0 {
// 		return
// 	}
// 	fParam := s[0]
// 	for _, arg := range fParam {
// 		z01.PrintRune(arg)
// 	}
// 	z01.PrintRune('\n')
// }

// // DISPLAY LAST PARAM

// func DisplayLastParam(s []string) {
// 	if len(s) == 0 {
// 		return
// 	}
// 	lParam := s[len(s)-1]

// 	for _, arg := range lParam {
// 		z01.PrintRune(arg)
// 	}
// 	z01.PrintRune('\n')
// }

// // LENGTH STRING

// func StrLen(s string) int {
// 	count := 0

// 	for _, ch := range s {
// 		if ch != ' ' {
// 			count++
// 		}
// 	}
// 	return count
// }

// // DISPLAYALPHAM

// func DisplayAlphaM() {
// 	for i := 'a'; i <= 'z'; i++ {
// 		if i%2 == 0 {
// 			z01.PrintRune(i - 32)
// 		} else {
// 			z01.PrintRune(i)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func DisplayAlRevM() {
// 	for i := 'z'; i >= 'a'; i-- {
// 		if i%2 != 0 {
// 			z01.PrintRune(i - 32)
// 		} else {
// 			z01.PrintRune(i)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// // WDMATCH
// func wdMatch(fStr, lStr string) {
// 	if canRewrite(fStr, lStr) {
// 		for _, ch := range fStr {
// 			if ContainChar(lStr, ch) {
// 				z01.PrintRune(ch)
// 			}
// 		}
// 		z01.PrintRune('\n')
// 	} else {
// 		return
// 	}
// }

// func canRewrite(str1, str2 string) bool {
// 	index2 := 0

// 	for _, char1 := range str1 {
// 		found := false

// 		for i := index2; i < len(str2); i++ {
// 			if rune(str2[i]) == char1 {
// 				index2 = i + 1
// 				found = true
// 				break
// 			}
// 		}

// 		if !found {
// 			return false
// 		}
// 	}

// 	return true
// }

// func ContainChar(str string, ch rune) bool {
// 	for _, char := range str {
// 		if char == ch {
// 			return true
// 		}
// 	}
// 	return false
// }

// // LastWord
// func LastWord(str string) string {
// 	res := ""

// 	firstIndex := 0

// 	trimmedStr := trimSpaces(str)

// 	for i := len(trimmedStr) - 1; i >= 0; i-- {
// 		if trimmedStr[i] == ' ' {
// 			firstIndex = i + 1
// 			break
// 		}
// 	}

// 	if firstIndex >= len(trimmedStr) {
// 		return res
// 	}

// 	// Build the result string
// 	res = trimmedStr[firstIndex:]

// 	return res
// }

// // Function to trim leading and trailing spaces without using strings.TrimSpace
// func trimSpaces(s string) string {
// 	start, end := 0, len(s)

// 	// Trim leading spaces
// 	for start < end && s[start] == ' ' {
// 		start++
// 	}

// 	// Trim trailing spaces
// 	for end > start && s[end-1] == ' ' {
// 		end--
// 	}

// 	return s[start:end]
// }

// // ISPOWER2

// func IsPower2(n int) bool {
// 	return n > 0 && (n&(n-1)) == 0
// }

// func ROT13(s string) {
// 	for i := 0; i < len(s); i++ {
// 		if s[i] >= 'a' && s[i] < 'z' {
// 			if s[i]+13 >= 'z' {
// 				z01.PrintRune(rune(s[i] - 13))
// 			} else {
// 				z01.PrintRune(rune(s[i] + 13))
// 			}
// 		} else if s[i] >= 'A' && s[i] < 'Z' {
// 			if s[i]+13 >= 'Z' {
// 				z01.PrintRune(rune(s[i] - 13))
// 			} else {
// 				z01.PrintRune(rune(s[i] + 13))
// 			}
// 		} else {
// 			z01.PrintRune(rune(s[i]))
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func Max(arr []int) int {
// 	max := arr[0]

// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] > max {
// 			max = arr[i]
// 		}
// 	}

// 	return max
// }

// func ReduceInt(a []int, f func(int, int) int) {
// 	res := a[0]
// 	for i := 1; i < len(a); i++ {
// 		res = f(res, a[i])
// 	}
// }

// // PrintBits

// // func PrintBits(num int) {
// // 	for i := 8; i >= 0; i-- {
// // 		bit := (num >> uint(i)) & 1
// // 		char := '0'
// // 		if bit == 1 {
// // 			char = '1'
// // 		}
// // 		z01.PrintRune(char)
// // 	}
// // }

// func Addprimesum(num int) {
// 	res := 0

// 	if !isPrime(num) {
// 		return
// 	}
// 	for i := 1; i <= num; i++ {
// 		if isPrime(i) {
// 			res += i
// 		}
// 	}
// 	fmt.Println(res)
// }

// func isPrime(nb int) bool {
// 	if nb <= 1 {
// 		return false
// 	}

// 	for i := 2; i*i <= nb; i++ {
// 		if nb%i == 0 {
// 			return false
// 		}
// 	}

// 	return true
// }
