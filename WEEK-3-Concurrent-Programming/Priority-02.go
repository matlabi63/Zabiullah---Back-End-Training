package main

import (
	"fmt"
	"math"
)

func generatePrimes(limit int, ch chan<- int) {
	for num := 2; num <= limit; num++ {
		if isPrime(num) {
			ch <- num
		}
	}
	close(ch) 
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func calculateSquares(primes <-chan int, results chan<- int) {
	for prime := range primes {
		results <- prime * prime
	}
	close(results)
}

func main() {
	limit := 100
	primeChannel := make(chan int)
	resultChannel := make(chan int)

	go generatePrimes(limit, primeChannel)

	go calculateSquares(primeChannel, resultChannel)

	fmt.Println("Squares of prime numbers from 1 to 100:")
	for result := range resultChannel {
		fmt.Println(result)
	}
}
