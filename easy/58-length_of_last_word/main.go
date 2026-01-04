package main

import "log"

func main() {

	tests := []struct {
		input          string
		expectedOutput int
	}{
		{
			input:          "Hello World",
			expectedOutput: 5,
		},
		{
			input:          "   fly me   to   the moon  ",
			expectedOutput: 4,
		},
		{
			input:          "luffy is still joyboy",
			expectedOutput: 6,
		},
		{
			input:          "a",
			expectedOutput: 1,
		},
	}

	for _, test := range tests {
		output := lengthOfLastWord(test.input)

		if test.expectedOutput != output {
			log.Fatalf("input %v, output is %d but expected is %d", test.input, output, test.expectedOutput)
		}
	}

}

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	// skip trailing spaces
	for i >= 0 && s[i] == ' ' {
		i--
	}

	worldLength := 0
	for i >= 0 && s[i] != ' ' {
		worldLength++
		i--
	}

	return worldLength
}
