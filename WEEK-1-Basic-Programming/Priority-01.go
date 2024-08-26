package main

import "fmt"

func main() {
	var phi = 3.14

	
	var radius, height float64

	fmt.Print("Enter the radius of the tube: ")
	fmt.Scanln(&radius)

	fmt.Print("Enter the height of the tube: ")
	fmt.Scanln(&height)

	volume := phi * radius * radius * height

	fmt.Println("The volume of the tube is: ", volume)
}
