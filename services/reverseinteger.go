package services

import (
	"math"
	"strconv"
	"strings"
)

func BestReverseInt(x int) int {
	var rev int
	for x != 0 {
		digit := x % 10 // Extract last digit
		x /= 10         // Remove last digit

		// Check overflow/underflow before updating rev
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && digit > 7) {
			return 0 // Overflow (MaxInt32 = 2,147,483,647 → last digit 7)
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && digit < -8) {
			return 0 // Underflow (MinInt32 = -2,147,483,648 → last digit -8)
		}

		rev = rev*10 + digit // Append digit
	}
	return rev
}

func ReverseInt(x int) int {
	positive := true
	if x < 0 {
		positive = false
		x = x * -1
	}
	s := strconv.Itoa(x)
	var sb strings.Builder
	sb.Grow(len(s)) // Pre-allocate memory
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i]) // Works for ASCII only
	}
	num, _ := strconv.Atoi(sb.String())
	if !positive {
		num = num * -1
	}
	return num
}
