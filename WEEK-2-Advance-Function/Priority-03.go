package main

import (
	"fmt"
)

func main() {
	fmt.Println(groupPalindrome([]string{"katak", "civic", "kasur", "rusak"}))                  //output [[katak, civic], kasur, rusak]
	fmt.Println(groupPalindrome([]string{"racecar", "seru", "kasur", "civic", "bilik", "kak"})) //output [[racecar, civic, kak], seru, kasur, bilik]
	fmt.Println(groupPalindrome([]string{"masuk", "civic", "hahah", "garam"}))                  //output [[civic, hahah], masuk, garam]
}

func groupPalindrome(words []string) [][]string {
	var result [][]string
	var palindromes []string

	for _, word := range words {
		if isPalindrome(word) {
			palindromes = append(palindromes, word)
		} else {
			if len(palindromes) > 0 {
				result = append(result, palindromes)
				palindromes = nil
			}
			result = append(result, []string{word})
		}
	}

	if len(palindromes) > 0 {
		result = append(result, palindromes)
	}

	return result
}

func isPalindrome(word string) bool {
	runes := []rune(word)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}
	return true
}
