package services

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	digitMap := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi",
		'5': "jkl", '6': "mno", '7': "pqrs",
		'8': "tuv", '9': "wxyz",
	}
	var result []string
	var backtrack func(int, string)
	backtrack = func(index int, path string) {
		if index == len(digits) {
			result = append(result, path)
			return
		}
		letters := digitMap[digits[index]]
		for _, c := range letters {
			backtrack(index+1, path+string(c))
		}
	}
	backtrack(0, "")
	return result
}
