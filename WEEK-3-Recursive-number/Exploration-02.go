package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getItem([]int{10, 10, 5, 30, 40, 10, 5}, 80)) // [40 30 10]
	fmt.Println(getItem([]int{50, 20, 10, 25, 25}, 100))      // [50 25 25]
}

func getItem(items []int, target int) []int {

	sort.Slice(items, func(i, j int) bool {
		return items[i] > items[j]
	})

	var result []int
	currentSum := 0

	for _, item := range items {
		if currentSum+item <= target {
			result = append(result, item)
			currentSum += item
		}
		if currentSum == target {
			break
		}
	}

	return result
}
