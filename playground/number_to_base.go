package playground

import (
	//"strconv"
	"strings")


// func NumberToBase(number int, base int) string {
// 	var remainder int = number
// 	var modulus int
// 	var result strings.Builder

// 	for remainder > 0 {
// 		modulus = remainder % base
// 		remainder = remainder / base

// 		c := castInt(modulus)
// 		result.WriteString(c)
// 	}
// 	return ReverseString(result.String())
// }

// func castInt(n int) string {
// 	return strconv.Itoa(n)
// }

// NumberToBase converts a number to a string representation in the specified base.
func NumberToBase(number int, base int) string {
	if base < 2 || base > 36 {
		return ""
	}

	if number == 0 {
		return "0"
	}

	var result strings.Builder
	remainder := number

	for remainder > 0 {
		modulus := remainder % base
		remainder /= base

		if modulus < 10 {
			result.WriteByte(byte(modulus + '0'))
		} else {
			result.WriteByte(byte(modulus-10+'A'))
		}
	}

	// Reverse the result string
	return ReverseString(result.String())
}