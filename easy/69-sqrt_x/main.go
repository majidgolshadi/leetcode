package main

import (
	"log"
)

func main() {

	tests := []struct {
		input          int
		expectedOutput int
	}{
		{
			input:          4,
			expectedOutput: 2,
		},
		{
			input:          8,
			expectedOutput: 2,
		},
		{
			input:          1,
			expectedOutput: 1,
		},
	}

	for _, test := range tests {
		output := mySqrt(test.input)

		if test.expectedOutput != output {
			log.Fatalf("input %d, output is %d but expected is %d", test.input, output, test.expectedOutput)
		}
	}

}

func mySqrt(x int) int {
	number := 1
	for result := 1; result <= x; result = number * number {
		number++
	}

	return number - 1
}
