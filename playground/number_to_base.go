package playground

import (
	"strconv"
	"strings")


func NumberToBase(number int, base int) string {
	var remainder int = number
	var modulus int
	var result strings.Builder

	for remainder > 0 {
		modulus = remainder % base
		remainder = remainder / base

		c := castInt(modulus)
		result.WriteString(c)
	}
	return ReverseString(result.String())
}

func castInt(n int) string {
	return strconv.Itoa(n)
}