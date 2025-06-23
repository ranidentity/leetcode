package services

// fmt.Println(services.LengthOfLongestSubstring("a b c a b c")) // 3
// fmt.Println(services.LengthOfLongestSubstring("abcabcbb"))    // 3
// fmt.Println(services.LengthOfLongestSubstring("  "))          // 0

func LengthOfLongestSubstring(str string) int {
	m := make(map[rune]int)
	max, left := 0, 0
	for idx, c := range str {
		if _, okay := m[c]; okay == true && m[c] >= left {
			// m[c] = location of the character, if exist then move left +1,
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
