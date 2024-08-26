package main

import (
	"fmt"
)


func generateComposites(limit int, ch chan<- int) {
	for num := 2; num <= limit; num++ {
		if isComposite(num) {
			ch <- num
		}
	}
	close(ch) 
}

func isComposite(n int) bool {
	if n <= 1 {
		return false
	}
	factorCount := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factorCount++
		}
	}
	return factorCount > 2
}

func calculateSquares(composites <-chan int, squares chan<- int) {
	for composite := range composites {
		squares <- composite * composite
	}
	close(squares) 
}

func checkEvenOdd(squares <-chan int) {
	for square := range squares {
		if square%2 == 0 {
			fmt.Println("Pong")
		} else {
			fmt.Println("Ping")
		}
	}
}

func main() {
	limit := 100
	compositeChannel := make(chan int)
	squareChannel := make(chan int)

	go generateComposites(limit, compositeChannel)

	go calculateSquares(compositeChannel, squareChannel)

	go checkEvenOdd(squareChannel)

	var input string
	fmt.Scanln(&input)
}
