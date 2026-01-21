package main

import "fmt"

/*
* Key Differences from Traditional Inheritance
* No Method Overriding (Shadowing Only):
* You cannot truly "override" a method.
* If the outer struct defines a method with the same name as the embedded struct, the outer one takes precedence when called on the outer struct.
* However, methods belonging to the inner struct will still call the inner struct's version of the method, not the "shadowed" one.
 */

// The "Base" struct
type Vehicle struct {
	Speed int
	Brand string
}

func (v *Vehicle) Move() {
	fmt.Printf("Moving at %d km/h\n", v.Speed)
}

// The "Child" struct
type Car struct {
	// We embed Vehicle here (no field name, just the type)
	Vehicle
	Brand string
}

func (c *Car) Move() {
	fmt.Printf("Moving car at %d km/h\n", c.Speed)
}

func main() {
	c := Car{
		Brand:   "Ferrari",
		Vehicle: Vehicle{Speed: 200, Brand: "Majid"},
	}

	// We can call Move() directly on c, as if Car inherited it.
	// This is method promotion.
	c.Move()

	// We can also access the embedded field directly
	fmt.Println(c.Speed)
}
