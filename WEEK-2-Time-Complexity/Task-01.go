package main

import (
	"fmt"
	"math"
)

// primeNumber checks if a given number is prime
func primeNumber(number int) bool {
	// Any number less than or equal to 1 is not prime
	if number <= 1 {
		return false
	}
	// 2 and 3 are prime numbers
	if number <= 3 {
		return true
	}
	// Eliminate even numbers and multiples of 3
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	// Check for factors up to the square root of the number
	sqrt := int(math.Sqrt(float64(number)))
	for i := 5; i <= sqrt; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(1500450271)) // false
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
