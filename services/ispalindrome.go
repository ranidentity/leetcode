package services

// 1234567
func IsPalindrome(x int) bool {
	reverseX := 0 // 76
	for num := x; num > 0; {
		rem := num % 10              //7,6
		reverseX = reverseX*10 + rem //76
		num = num / 10
	}
	if x != reverseX {
		return false
	}
	return true
}
