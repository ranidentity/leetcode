package services

import (
	"fmt"
	"strings"
)

var alternatives = map[string][]string{
	"jalan":   {"jln", "Jln", "Jalan"},
	"gertak":  {"grtk", "gtk"},
	"sanggul": {"sanggul", "Sanggul"},
}

// Recursive function to generate combinations
func generateCombinations(words []string, index int, current []string, result *[]string) {
	if index == len(words) {
		*result = append(*result, strings.Join(current, " "))
		return
	}

	word := strings.ToLower(words[index])
	if alt, ok := alternatives[word]; ok { // if exist
		for _, option := range alt {
			fmt.Println(option)
			generateCombinations(words, index+1, append(current, option), result)
		}
	}
}

func ImplementingGenerateCombinations() {
	input := "Jalan Gertak Sanggul"
	words := strings.Split(input, " ")

	var result []string
	generateCombinations(words, 0, []string{}, &result)

	// for _, combo := range result {
	// 	fmt.Println(combo)
	// }

	fmt.Printf("Total combinations: %d\n", len(result))
}
