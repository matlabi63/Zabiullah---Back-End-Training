package main

import ( 
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%4 == 0 && i%7 == 0:
			fmt.Println("buzz")
		case i%4 == 0:
			fmt.Println(i * i)
		case i%7 == 0:
			fmt.Println(int(math.Pow(float64(i), 3)))
		default:
			fmt.Println(i)
		}
	}
}
