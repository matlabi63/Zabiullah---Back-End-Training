package main

import (
	"fmt"
	"sort"
	"strings"
)
// Test case     
// Input:
// First word: cafe
// Second word: face

// Output:
// Anagram

// Input:
// First word: donut
// Second word: peanut

// Output:
// Not Anagram

func main() {
	var firstWord, secondWord string

	// Input from user
	fmt.Println("Enter the first word:")
	fmt.Scanln(&firstWord)
	fmt.Println("Enter the second word:")
	fmt.Scanln(&secondWord)

	firstWord = normalizeString(firstWord)
	secondWord = normalizeString(secondWord)

	// Check if words are anagrams
	if isAnagram(firstWord, secondWord) {
		fmt.Println("Anagram")
	} else {
		fmt.Println("Not Anagram")
	}
}

// Function to normalize string by converting to lowercase and removing spaces
func normalizeString(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", ""))
}

// Function to check if two strings are anagrams
func isAnagram(s1, s2 string) bool {
	// If lengths of strings are different, they can't be anagrams
	if len(s1) != len(s2) {
		return false
	}

	// Convert strings to rune slices
	r1 := []rune(s1)
	r2 := []rune(s2)

	// Sort rune slices
	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })

	// Compare sorted rune slices
	for i := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}
