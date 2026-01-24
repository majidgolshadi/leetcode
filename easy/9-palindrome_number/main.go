package main

import (
	"log"
	"strconv"
)

// An integer is a palindrome when it reads the same forward and backward.
// For example, 121 is a palindrome while 123 is not.

func main() {

	tests := []struct {
		input          int
		expectedOutput bool
	}{
		{
			input:          121,
			expectedOutput: true,
		},
		{
			input:          -121,
			expectedOutput: false,
		},
		{
			input:          10,
			expectedOutput: false,
		},
	}

	for _, test := range tests {
		output := isPalindrome(test.input)

		if test.expectedOutput != output {
			log.Fatalf("input %v, expected %t, output is %t", test.expectedOutput, output)
		}
	}

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	number := strconv.FormatInt(int64(x), 10)
	length := len(number)

	for i := 0; i <= length/2; i++ {
		if number[i] != number[length-i-1] {
			return false
		}
	}

	return true
}
