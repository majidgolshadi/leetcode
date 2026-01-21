package main

import (
	"fmt"
)

/*
* Mutable:
* Mutable data type is a data type that can be changed without reallocating any chunk of the memory assigned at the time of decalration
*
* Immutable:
* Immutable data types are ones that can’t be changed once they’re created
*
*  Mutable Data Types:
*   - Slice
*   - Map
*   - Channels
*
* Immutable Data Types:
*   - Boolean, Int, Float
*   - Pointers
*   - String
*   - Interfaces
*/

func main() {

	// mutable
	var x []int = []int{1, 2, 3, 4, 5}
	y := x
	y[0] = 100
	fmt.Printf("x = %v\n", x)
	fmt.Printf("y = %v\n", y)
	// result
	// x = [100 2 3 4 5]
	// y = [100 2 3 4 5]

	// immutable
	var z = 5
	k := z
	k = 7
	fmt.Printf("x = %v\n", z)
	fmt.Printf("y = %v\n", k)
	// result
	// x = 5
	// y = 7
}
