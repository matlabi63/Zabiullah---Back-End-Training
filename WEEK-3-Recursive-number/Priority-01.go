package main

import "fmt"

func main() {
    fmt.Println(sumFibonacci(5))  // Output: 12
    fmt.Println(sumFibonacci(10)) // Output: 143
}

func sumFibonacci(n int) int {
    sum := 0
    for i := 0; i <= n; i++ {
        sum += fibonacciRecursive(i)
    }
    return sum
}

func fibonacciRecursive(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}
