package main

import "fmt"

func main() {
  fmt.Println(getSequence(4))   //output 10
  fmt.Println(getSequence(15))  //output 120
  fmt.Println(getSequence(100)) //output 5050
}

func getSequence(n int) int {
  return n * (n + 1) / 2
}
