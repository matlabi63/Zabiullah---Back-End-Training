package main

import "fmt"

// pow calculates x raised to the power n using exponentiation by squaring iteratively
func pow(x, n int) int {
	result := 1
	base := x

	for n > 0 {
		if n%2 == 1 { // If n is odd
			result *= base
		}
		base *= base // Square the base
		n /= 2      // Divide n by 2
	}

	return result
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
