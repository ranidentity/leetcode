package main

import (
	"leetcode/services"
)

func main() {
	// arr1 := []int{1, 2}
	// arr2 := []int{3, 4}
	// fmt.Printf("%f\n", services.FindMedianSortedArrays(arr1, arr2))
	// fmt.Println(services.GetPalindromic("labba5789987"))
	services.ImplementingGenerateCombinations()
	"fmt"
)

func main() {
	A := []int{1, 0, 4, 1}
	var outcome []int
	var next int
	var i int
	for next != 0 || len(A) > i {
		var val int
		if next > 0 && len(A) <= i {
			val = next
		} else {
			val = A[i] + next
		}
		remainder := val % 2
		next = val / 2
		outcome = append(outcome, remainder)
		i++
	}
	fmt.Println(outcome)
	var result int
	for _, i := range outcome {
		result = result*10 + i
	}
	fmt.Println(result)
}

func main2() {
	s := "axxaxa"
	hashmap := make(map[rune]int)
	extraCount := 0
	for _, i := range s {
		if _, exists := hashmap[i]; exists {
			hashmap[i]++
		} else {
			hashmap[i] = 1
		}
	}
	for _, i := range hashmap {
		if i%2 == 1 {
			extraCount++
		}
	}

	fmt.Println(extraCount)
}
