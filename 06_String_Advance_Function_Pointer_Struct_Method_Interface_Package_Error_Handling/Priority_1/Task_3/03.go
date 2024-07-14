package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(groupPalindrome([]string{"katak", "civic", "kasur", "rusak"}))                  // [[katak, civic], kasur, rusak]
	fmt.Println(groupPalindrome([]string{"racecar", "seru", "kasur", "civic", "bilik", "kak"})) // [[racecar, civic, kak], seru, kasur, bilik]
	fmt.Println(groupPalindrome([]string{"masuk", "civic", "hahah", "garam"}))                  // [[civic, hahah], masuk, garam]
}

func groupPalindrome(words []string) []any {
	isPalindrome := func(word string) bool {
		normalized := strings.ToLower(word)
		for i, j := 0, len(normalized)-1; i < j; i, j = i+1, j-1 {
			if normalized[i] != normalized[j] {
				return false
			}
		}
		return true
	}

	var palindromes, others []any

	for _, word := range words {
		if isPalindrome(word) {
			palindromes = append(palindromes, fmt.Sprintf("[%s]", word))
		} else {
			others = append(others, word)
		}
	}

	return append(palindromes, others...)
}
