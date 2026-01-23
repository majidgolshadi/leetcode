package main

import (
	"fmt"
)

// Append might not update the original:
// If append causes a re-allocation (exceeding Cap), the function's local slice header points to a new array,
// but the caller's slice header still points to the old array.
func mutate(s []int) {
	// Slice Growth & Allocation Strategy
	// When you append to a slice that is full (len == cap), Go allocates a new, larger array and copies the data over.
	// Heuristic:
	//   Traditionally, it doubled capacity for small slices and grew by ~1.25x for large slices to avoid memory fragmentation.
	//   (Note: The exact formula evolves with Go versions, but understanding the O(N) copy cost during growth is key).
	// Optimization:
	//   If you know the size beforehand, always pre-allocate using make([]T, len, cap).
	//   This prevents repeated allocations and copies.

	s = append(s, 4)
	s[0] = 9
	// As a result the append is not appear to outside of method
}

func main() {
	// Slice:
	// A dynamically-sized, flexible view into an array. It is a lightweight wrapper over an underlying array.
	slice := []int{1, 2, 3}
	println(slice)

	// Array:
	// A fixed-size sequence of elements of a single type.
	// The size is part of the type signature (e.g., [4]int and [5]int are distinct, incompatible types).
	array := [4]int{1, 2, 3, 4}
	println(array)

	mutate(slice)
	fmt.Println(slice)
	// result: [1 2 3]
	// The Reason:
	// In main, the slice is initialized with len: 3 and cap: 3.
	// When passed to mutate, a copy of the Slice Header is created.
	// Inside mutate, append(slice, 4) exceeds the capacity.
	// Go allocates a new underlying array, copies the elements, and updates the local slice header to point to this new memory.
	// The modification slice[0] = 9 happens on the new array.
	// The original slice header in main still points to the old array, which remains unchanged.
	//
	// Solution:
	// If slice had been initialized with make([]int, 3, 10), 
	// the output would be [9 2 3] because the append wouldn't have triggered a reallocation, 
	// and the slice[0] change would have affected the shared underlying array.
}
