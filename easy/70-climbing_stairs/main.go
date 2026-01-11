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
			input:          1,
			expectedOutput: 1,
		},
		{
			input:          2,
			expectedOutput: 2,
		},
		{
			input:          3,
			expectedOutput: 3,
		},
		{
			input:          4,
			expectedOutput: 5,
		},
		{
			input:          5,
			expectedOutput: 8,
		},
	}

	for _, test := range tests {
		output := climbStairs(test.input)

		if test.expectedOutput != output {
			log.Fatalf("input %d, output is %d but expected is %d", test.input, output, test.expectedOutput)
		}
	}

}

// f(n)=f(n−1)+f(n−2)
func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	prev2 := 1 // f(1)
	prev1 := 2 // f(2)

	for i := 3; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}