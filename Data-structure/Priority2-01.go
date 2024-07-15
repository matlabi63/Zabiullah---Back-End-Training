package main

import "fmt"

func main() {
	fmt.Println(spinSlice([]int{1, 2, 3, 4, 5}))    //output [1, 5, 2, 4, 3]
	fmt.Println(spinSlice([]int{6, 7, 8}))          //output [6, 8, 7]
	fmt.Println(spinSlice([]int{8, 5, 4, 3, 2, 1})) //ouput [8, 1, 5, 2, 4, 3]
}

func spinSlice(numbers []int) []int {
	n := len(numbers)
	result := make([]int, n)
	
	left := 0
	right := n - 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[left] = numbers[i]
			left++
		} else {
			result[right] = numbers[i]
			right--
		}
	}

	return result
}
