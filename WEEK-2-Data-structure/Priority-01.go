package main

import "fmt"

func main() {
	res1 := merge([][]int{
		{1, 2, 5, 4, 3},
		{1, 2, 7, 8},
	})
	fmt.Println(res1) //output [1,2,5,4,3,7,8]

	res2 := merge([][]int{
		{1, 2, 5, 4, 5},
		{7, 9, 10, 5},
	})
	fmt.Println(res2) //output [1, 2, 5, 4, 7, 9, 10]

	res3 := merge([][]int{
		{1, 4, 5},
		{7},
	})
	fmt.Println(res3) //output [1, 4, 5, 7]
}

func merge(data [][]int) []int {
	
	uniqueElements := make(map[int]bool)

	for _, arr := range data {
		for _, num := range arr {
			uniqueElements[num] = true
		}
	}

	result := make([]int, 0, len(uniqueElements))
	for num := range uniqueElements {
		result = append(result, num)
	}

	return result
}
