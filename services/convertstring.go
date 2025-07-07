package services

import (
	"math"
	"strings"
	"unicode"
)

func MyAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	sign := 1
	result := 0
	i := 0

	// handle sign
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')

		// Check overflow/underflow
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}
