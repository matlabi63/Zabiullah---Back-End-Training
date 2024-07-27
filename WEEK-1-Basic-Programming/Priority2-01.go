package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	result := ""
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			if i%2 == 0 {
				result += "I"
			} else {
				result += "O"
			}
		}
	}

	fmt.Println("Output:", result)
}
