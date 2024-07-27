package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(groupPrime([]int{2, 3, 4, 5, 6, 7, 8})) //output [[2,3,5,7],4,6,8]
	fmt.Println(groupPrime([]int{15, 24, 3, 9, 5}))     //output [[3,5],15,24,9]
	fmt.Println(groupPrime([]int{4, 8, 9, 12}))         //output [4,8,9,12]
}

func groupPrime(numbers []int) []any {
	var primes []int
	var nonPrimes []any

	for _, number := range numbers {
		if isPrime(number) {
			primes = append(primes, number)
		} else {
			nonPrimes = append(nonPrimes, number)
		}
	}

	if len(primes) > 0 {
		nonPrimes = append([]any{primes}, nonPrimes...)
	}

	return nonPrimes
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
