package services

import "fmt"

func GetPalindromic(s string) string {
	len := len(s)
	if len == 0 {
		return ""
	}

	var l, r, pl, pr int // l is the initial index, r is the moving index that decide the length
	for r < len {
		fmt.Printf("begin: r = %d l = %d s[l]:%c s[r]:%c \n", r, l, s[l], s[r])
		// gobble up dup chars
		for r+1 < len && s[l] == s[r+1] { // if same with next character. this will form the center part of the palindrome
			fmt.Printf("loop 1: r:%d s[l]:%c s[r+1]:%c \n", r, s[l], s[r+1])
			r++ // check next character
		}
		// find size of this palindrome
		for l-1 >= 0 && r+1 < len && s[l-1] == s[r+1] { // this will form the "wings" on both side
			fmt.Printf("loop 2: l:%d r:%d s[l-1]:%c s[r+1]:%c\n", l, r, s[l-1], s[r+1])
			l--
			r++
		}
		if r-l > pr-pl {
			pl, pr = l, r
		}
		fmt.Printf("pl:%d pr:%d \n", pl, pr)
		// reset to next mid point
		l = (l+r)/2 + 1
		r = l
	}
	return s[pl : pr+1]
}
