package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(convertToBinaryList(2)) // Output: [0, 1, 10]
	fmt.Println(convertToBinaryList(3)) // Output: [0, 1, 10, 11]
}


func convertToBinaryList(n int) []string {

	binaryList := make([]string, n+1)
	

	for i := 0; i <= n; i++ {
		binaryList[i] = toBinary(i)
	}
	
	return binaryList
}

func toBinary(num int) string {
	return strconv.FormatInt(int64(num), 2)
}
