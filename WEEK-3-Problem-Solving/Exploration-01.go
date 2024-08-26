package main

import "fmt"


func Frog(jumps []int) int {
	n := len(jumps)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}

	
	dp := make([]int, n)
	dp[0] = 0 


	if n > 1 {
		dp[1] = abs(jumps[1] - jumps[0])
	}

	
	for i := 2; i < n; i++ {
		
		dp[i] = min(dp[i-1]+abs(jumps[i]-jumps[i-1]),
			dp[i-2]+abs(jumps[i]-jumps[i-2]))
	}

	return dp[n-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
