package main

import (
	"fmt"
	"strings"
)

// Define a struct to represent the operation
type SentenceProcessor struct {
	sentence string
}

// Define an interface for processing sentences
type Processor interface {
	Process() string
}

// Implement the Process method for SentenceProcessor
func (sp *SentenceProcessor) Process() string {
	return pickVocals(sp.sentence)
}

// Function to pick vocal letters from a sentence
func pickVocals(sentence string) string {
	vowels := "aeiouAEIOU"
	var result []string

	words := strings.Fields(sentence)
	for _, word := range words {
		var filteredWord strings.Builder
		for _, char := range word {
			if strings.ContainsRune(vowels, char) {
				filteredWord.WriteRune(char)
			}
		}
		result = append(result, filteredWord.String())
	}
	return strings.Join(result, " ")
}

func main() {
	sp1 := &SentenceProcessor{"alterra academy"}
	sp2 := &SentenceProcessor{"i love programming"}
	sp3 := &SentenceProcessor{"go is awesome programming language"}

	fmt.Println(sp1.Process()) //output aea aae
	fmt.Println(sp2.Process()) //output i oe oai
	fmt.Println(sp3.Process()) //output o i aeoe oai auae
}
