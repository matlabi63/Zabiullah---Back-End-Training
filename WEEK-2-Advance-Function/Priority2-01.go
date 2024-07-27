package main

import (
	"fmt"
)

// Vehicle struct to hold vehicle information
type Vehicle struct {
	Brand     string
	Type      string
	RentPrice float64
}

// Customer struct to hold customer information and rented vehicles
type Customer struct {
	Name          string
	Address       string
	RentedVehicles []Vehicle
}

// RentVehicle method adds a vehicle to the customer's rented vehicles list
func (c *Customer) RentVehicle(vehicle Vehicle) {
	c.RentedVehicles = append(c.RentedVehicles, vehicle)
}

// ReturnVehicle method removes a vehicle from the customer's rented vehicles list
func (c *Customer) ReturnVehicle(vehicle Vehicle) {
	for i, v := range c.RentedVehicles {
		if v.Brand == vehicle.Brand && v.Type == vehicle.Type {
			c.RentedVehicles = append(c.RentedVehicles[:i], c.RentedVehicles[i+1:]...)
			break
		}
	}
}

// ListCustomers function displays a list of all customers
func ListCustomers(customers []Customer) {
	fmt.Println("List of Customers:")
	for _, customer := range customers {
		fmt.Printf("Name: %s, Address: %s\n", customer.Name, customer.Address)
	}
}

// ListRentedVehicles function displays a list of all rented vehicles
func ListRentedVehicles(customers []Customer) {
	fmt.Println("List of Rented Vehicles:")
	for _, customer := range customers {
		if len(customer.RentedVehicles) > 0 {
			fmt.Printf("Customer: %s\n", customer.Name)
			for _, vehicle := range customer.RentedVehicles {
				fmt.Printf("- %s %s\n", vehicle.Brand, vehicle.Type)
			}
		}
	}
}

func main() {
	// Initialize customers slice
	var customers []Customer

	// Add some sample customers and vehicles
	customer1 := Customer{Name: "Alice", Address: "123 Main St"}
	customer2 := Customer{Name: "Bob", Address: "456 Elm St"}

	vehicle1 := Vehicle{Brand: "Toyota", Type: "Camry", RentPrice: 50.0}
	vehicle2 := Vehicle{Brand: "Honda", Type: "Accord", RentPrice: 60.0}
	vehicle3 := Vehicle{Brand: "Ford", Type: "Fusion", RentPrice: 55.0}

	// Rent vehicles for customers
	customer1.RentVehicle(vehicle1)
	customer1.RentVehicle(vehicle2)
	customer2.RentVehicle(vehicle3)

	// Add customers to the slice
	customers = append(customers, customer1)
	customers = append(customers, customer2)

	// List all customers
	ListCustomers(customers)

	// List rented vehicles
	ListRentedVehicles(customers)

	// Example: Return a vehicle for customer1
	customer1.ReturnVehicle(vehicle1)

	// Updated list of rented vehicles
	ListRentedVehicles(customers)
}
