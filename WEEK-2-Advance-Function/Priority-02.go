package main

import "fmt"

func main() {
	fmt.Println(spinString("alta"))    //output aalt
	fmt.Println(spinString("alterra")) //output aalrtre
	fmt.Println(spinString("sepulsa")) //output saesplu
}

func spinString(word string) string {
	runes := []rune(word)
	n := len(runes)

	// Initialize two pointers
	left := 0
	right := n - 1

	// Create a slice to store the result
	result := make([]rune, n)

	// Iterate through the word
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = runes[left]
			left++
		} else {
			result[i] = runes[right]
			right--
		}
	}

	return string(result)
}
