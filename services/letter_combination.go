package services

import (
	"strings"
)

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

func LetterCombinations2(input string) []string {
	var alternatives = map[string][]string{
		"jalan":   {"jln", "Jln", "Jalan"},
		"gertak":  {"grtk", "gtk"},
		"sanggul": {"sanggul", "Sanggul"},
	}
	subject := strings.Split(input, " ")
	var result []string
	var backtrack func(i int, data string)
	backtrack = func(i int, data string) {
		if i == len(subject) {
			result = append(result, data)
			return
		}
		letters, ok := alternatives[strings.ToLower(subject[i])]
		if !ok {
			backtrack(i+1, data+subject[i]+" ")
			return
		}
		for _, ea := range letters {
			backtrack(i+1, data+ea+" ")
		}
	}
	backtrack(0, "")
	return result
}
