package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	
	x := (a + c - b) / 2
	y := (a + b - c) / 2
	z := (b + c - a) / 2

	if (a + b == x + y) && (b + c == y + z) && (a + c == x + z) && x != y && y != z && x != z {
		fmt.Printf("%d %d %d\n", x, y, z)
	} else {
		fmt.Println("no solution")
	}
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
