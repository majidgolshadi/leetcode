package main

import "fmt"


/*
* Key Differences from Traditional Inheritance
* Identity: 
* When a method of the embedded struct is called, the receiver is the embedded struct, not the outer one. It has no knowledge of the outer struct.
*/

// The "Base" struct
type Vehicle struct {
	Brand string
}

func (v *Vehicle) GetCarBrand() {
	fmt.Printf("Car Brand is %s\n", v.Brand)
}

// The "Child" struct
type Car struct {
	// We embed Vehicle here (no field name, just the type)
	Vehicle
	Brand string
}

func main() {
	c := Car{
		Brand:   "Ferrari",
		Vehicle: Vehicle{Brand: "Majid"},
	}

	c.GetCarBrand()
}