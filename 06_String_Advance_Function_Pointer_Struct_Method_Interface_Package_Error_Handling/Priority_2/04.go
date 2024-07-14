package main

import (
	"fmt"
)

// Define the Customer struct
type Customer struct {
	Name           string
	Address        string
	RentedVehicles []Vehicle
}

// Define the Vehicle struct
type Vehicle struct {
	Brand     string
	Type      string
	RentPrice float64
}

// Define a list to hold customers and vehicles
var customers []Customer
var availableVehicles []Vehicle

// Add a vehicle to the available vehicles list
func addVehicle(brand string, vehicleType string, rentPrice float64) {
	vehicle := Vehicle{Brand: brand, Type: vehicleType, RentPrice: rentPrice}
	availableVehicles = append(availableVehicles, vehicle)
}

// Display the list of customers
func displayCustomers() {
	for _, customer := range customers {
		fmt.Printf("Name: %s, Address: %s\n", customer.Name, customer.Address)
	}
}

// Display the list of rented vehicles
func displayRentedVehicles() {
	for _, customer := range customers {
		fmt.Printf("Customer: %s\n", customer.Name)
		for _, vehicle := range customer.RentedVehicles {
			fmt.Printf("  Rented Vehicle - Brand: %s, Type: %s, Rent Price: %.2f\n", vehicle.Brand, vehicle.Type, vehicle.RentPrice)
		}
	}
}

// Rent a vehicle to a customer
func rentVehicle(customerName, vehicleBrand, vehicleType string) {
	var customer *Customer
	for i := range customers {
		if customers[i].Name == customerName {
			customer = &customers[i]
			break
		}
	}
	if customer == nil {
		fmt.Println("Customer not found!")
		return
	}

	var vehicleIndex int = -1
	for i, vehicle := range availableVehicles {
		if vehicle.Brand == vehicleBrand && vehicle.Type == vehicleType {
			vehicleIndex = i
			break
		}
	}
	if vehicleIndex == -1 {
		fmt.Println("Vehicle not available!")
		return
	}

	// Append the vehicle to the customer's rented vehicles
	customer.RentedVehicles = append(customer.RentedVehicles, availableVehicles[vehicleIndex])
	// Remove the vehicle from the available vehicles list
	availableVehicles = append(availableVehicles[:vehicleIndex], availableVehicles[vehicleIndex+1:]...)
	fmt.Printf("Vehicle rented: %s %s to %s\n", vehicleBrand, vehicleType, customerName)
}

// Return a vehicle from a customer
func returnVehicle(customerName, vehicleBrand, vehicleType string) {
	var customer *Customer
	for i := range customers {
		if customers[i].Name == customerName {
			customer = &customers[i]
			break
		}
	}
	if customer == nil {
		fmt.Println("Customer not found!")
		return
	}

	var vehicleIndex int = -1
	for i, vehicle := range customer.RentedVehicles {
		if vehicle.Brand == vehicleBrand && vehicle.Type == vehicleType {
			vehicleIndex = i
			break
		}
	}
	if vehicleIndex == -1 {
		fmt.Println("Vehicle not rented by this customer!")
		return
	}

	// Return the vehicle to available vehicles
	availableVehicles = append(availableVehicles, customer.RentedVehicles[vehicleIndex])
	// Remove the vehicle from the customer's rented vehicles list
	customer.RentedVehicles = append(customer.RentedVehicles[:vehicleIndex], customer.RentedVehicles[vehicleIndex+1:]...)
	fmt.Printf("Vehicle returned: %s %s from %s\n", vehicleBrand, vehicleType, customerName)
}

func main() {
	// Add some initial vehicles to the system
	addVehicle("Toyota", "Sedan", 50.0)
	addVehicle("Honda", "SUV", 70.0)

	// Add some customers to the system
	customers = append(customers, Customer{Name: "Alice", Address: "123 Main St"})
	customers = append(customers, Customer{Name: "Bob", Address: "456 Elm St"})

	// Display initial lists
	fmt.Println("Customers:")
	displayCustomers()
	fmt.Println("\nAvailable Vehicles:")
	for _, v := range availableVehicles {
		fmt.Printf("Brand: %s, Type: %s, Rent Price: %.2f\n", v.Brand, v.Type, v.RentPrice)
	}

	// Rent a vehicle
	rentVehicle("Alice", "Toyota", "Sedan")

	// Display lists after renting
	fmt.Println("\nCustomers:")
	displayCustomers()
	fmt.Println("\nRented Vehicles:")
	displayRentedVehicles()
	fmt.Println("\nAvailable Vehicles:")
	for _, v := range availableVehicles {
		fmt.Printf("Brand: %s, Type: %s, Rent Price: %.2f\n", v.Brand, v.Type, v.RentPrice)
	}

	// Return a vehicle
	returnVehicle("Alice", "Toyota", "Sedan")

	// Display lists after returning
	fmt.Println("\nCustomers:")
	displayCustomers()
	fmt.Println("\nRented Vehicles:")
	displayRentedVehicles()
	fmt.Println("\nAvailable Vehicles:")
	for _, v := range availableVehicles {
		fmt.Printf("Brand: %s, Type: %s, Rent Price: %.2f\n", v.Brand, v.Type, v.RentPrice)
	}
}
