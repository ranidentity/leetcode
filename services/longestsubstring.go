package services

import "fmt"

func LengthOfLongestSubstring(str string) int {
	m := make(map[rune]int)
	max, left := 0, 0
	fmt.Println(str)
	for idx, c := range str {
		fmt.Printf("rune: %c idx: %d\n ", c, idx)
		if _, okay := m[c]; okay == true && m[c] >= left {
			fmt.Printf("repeated char with pointer m[c]: %d left: %d \n", m[c], left)
			if idx-left > max {
				max = idx - left
			}
			left = m[c] + 1
		}
		m[c] = idx
	}
	if len(str)-left > max {
		max = len(str) - left
	}
	return max
}
