package main

import (
	"log"
	"reflect"
)

func main() {

	tests := []struct {
		input          []int
		expectedOutput []int
	}{
		{
			input:          []int{1, 2, 3},
			expectedOutput: []int{1, 2, 4},
		},
		{
			input:          []int{4, 3, 2, 1},
			expectedOutput: []int{4, 3, 2, 2},
		},
		{
			input:          []int{9},
			expectedOutput: []int{1, 0},
		},
		{
			input:          []int{9, 9, 9},
			expectedOutput: []int{1, 0, 0, 0},
		},
	}

	for _, test := range tests {
		output := plusOne(test.input)

		if reflect.DeepEqual(test.expectedOutput, output) == false {
			log.Fatalf("input %v, output is %d but expected is %d", test.input, output, test.expectedOutput)
		}
	}

}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}
