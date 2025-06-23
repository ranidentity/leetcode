package main

import (
	"fmt"
	"leetcode/services"
)

func main() {
	fmt.Println(services.LengthOfLongestSubstring("a b c a b c")) // 3
	fmt.Println(services.LengthOfLongestSubstring("abcabcbb"))    // 3
	fmt.Println(services.LengthOfLongestSubstring("  "))          // 0
}
