package main

import (
	"log"
)

// An integer is a palindrome when it reads the same forward and backward.
// For example, 121 is a palindrome while 123 is not.

func main() {

	tests := []struct {
		input          string
		expectedOutput bool
	}{
		{
			input:          "()",
			expectedOutput: true,
		},
		{
			input:          "()[]{}",
			expectedOutput: true,
		},
		{
			input:          "(]",
			expectedOutput: false,
		},
		{
			input:          "([)]",
			expectedOutput: false,
		},
	}

	for _, test := range tests {
		output := isValid(test.input)

		if test.expectedOutput != output {
			log.Fatalf("input %v, expected %t, output is %t", test.input, test.expectedOutput, output)
		}
	}

}

func isValid(s string) bool {
	// If input is huge, pre-allocating slice capacity
	// (e.g., len(s)/2) can prevent resizing overhead.
	stack := make([]rune, 0, len(s)/2)

	// Map for cleaner lookup logic
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		// If it's a closing bracket
		if open, isClose := pairs[char]; isClose {

			// Check empty stack or mismatch
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}

			// Pop
			stack = stack[:len(stack)-1]
		} else {
			// Push
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
