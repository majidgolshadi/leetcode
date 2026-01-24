package main

import (
	"log"
	"reflect"
)

func main() {

	tests := []struct {
		input          []int
		target         int
		expectedOutput []int
	}{
		{
			input:          []int{2, 7, 11, 15},
			target:         9,
			expectedOutput: []int{0, 1},
		},
		{
			input:          []int{3, 2, 4},
			target:         6,
			expectedOutput: []int{1, 2},
		},
		{
			input:          []int{3, 3},
			target:         6,
			expectedOutput: []int{0, 1},
		},
	}

	for _, test := range tests {
		output := twoSum(test.input, test.target)

		if reflect.DeepEqual(test.expectedOutput, output) {
			log.Fatalf("input %v, target %d, output is %d", test.input, test.target, output)
		}
	}

}

func twoSum(nums []int, target int) []int {
	// Map allocation.
	// We don't know the exact size needed, but preventing re-hashing
	// for large inputs is a nice touch if N is roughly known.
	// Here, we just use standard initialization.
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		// Go idiom: "comma ok" idiom checks existence efficiently
		if idx, ok := seen[complement]; ok {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil // Or error, depending on spec
}
