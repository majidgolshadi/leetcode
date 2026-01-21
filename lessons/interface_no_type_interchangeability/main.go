package main

import "fmt"

/*
* Key Differences from Traditional Inheritance
* No Type Interchangeability:
* A Car is not of type Vehicle. You cannot pass a Car struct into a function that expects a Vehicle struct.
* You must use an Interface to achieve this.
 */

type Vehicle struct {
	Brand string
}

type Car struct {
	Vehicle // Embedding Vehicle
	Doors   int
}

// This function strictly expects a 'Vehicle' struct
func checkVehicle(v Vehicle) {
	fmt.Println("Checking vehicle:", v.Brand)
}

func main() {
	myCar := Car{
		Vehicle: Vehicle{Brand: "Toyota"},
		Doors:   4,
	}

	// --- THIS WILL FAIL ---
	// Error: cannot use myCar (type Car) as type Vehicle in argument to checkVehicle
	checkVehicle(myCar)

	// To make this work without interfaces, you have to explicitly
	// access the embedded field:
	checkVehicle(myCar.Vehicle) // This works, but it strips away the 'Car' identity
}
