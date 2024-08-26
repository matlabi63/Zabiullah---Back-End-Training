package main

import (
	"fmt"
)

func main() {
	primeFactors(20)   // 2, 2, 5
	primeFactors(75)   // 3, 5, 5
	primeFactors(1024) // 2, 2, 2, 2, 2, 2, 2, 2, 2, 2
}

func primeFactors(n int) {

	factors := []int{}

	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	if n > 2 {
		factors = append(factors, n)
	}

	for i, factor := range factors {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(factor)
	}
	fmt.Println()
}
