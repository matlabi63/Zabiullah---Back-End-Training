package main

import (
	"fmt"
	"time"
)

func reverseWord(word string, ch chan<- string) {
	length := len(word)
	var intermediate string
	for i := 0; i < length; i++ {
		
		intermediate = string(word[:i+1])
		reversed := reverse(intermediate)
		
		ch <- reversed
		time.Sleep(3 * time.Second) 
	}
	close(ch) 
}

func reverse(s string) string {
	r := []rune(s) 

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	word := "phone"
	ch := make(chan string)

	go reverseWord(word, ch)

	for reversed := range ch {
		fmt.Println(reversed)
	}
}
