package main

import (
	"log"
)

func main() {

	tests := []struct {
		input_1        string
		input_2        string
		expectedOutput string
	}{
		{
			input_1:        "11",
			input_2:        "1",
			expectedOutput: "100",
		},
		{
			input_1:        "111",
			input_2:        "1",
			expectedOutput: "1000",
		},
		{
			input_1:        "1010",
			input_2:        "1011",
			expectedOutput: "10101",
		},
		{
			input_1:        "1000",
			input_2:        "0111",
			expectedOutput: "1111",
		},
		{
			input_1:        "1",
			input_2:        "1",
			expectedOutput: "10",
		},
		{
			input_1:        "100",
			input_2:        "110010",
			expectedOutput: "110110",
		},
	}

	for _, test := range tests {
		output := addBinary(test.input_1, test.input_2)

		if test.expectedOutput != output {
			log.Fatalf("input_1 %s and input_2 %s, output is %s but expected is %s", test.input_1, test.input_2, output, test.expectedOutput)
		}
	}

}

func addBinary_versionOne(a string, b string) string {
	// equal size
	for len(b) != len(a) {
		if len(b) > len(a) {
			a = "0" + a
		} else {
			b = "0" + b
		}
	}

	result := ""
	carry := 0

	for i := len(a) - 1; i >= 0; i-- {
		t := int((a[i]-'0')+(b[i]-'0')) + carry
		result = string('0'+rune(t%2)) + result
		carry = int(t / 2)
	}

	if carry == 1 {
		return "1" + result
	}

	return result
}

func addBinary(a string, b string) string {
	result := ""
	carry := 0
	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 || carry > 0 {
		bitA, bitB := 0, 0
		if i >= 0 {
			bitA = int(a[i] - '0')
			i--
		}
		if j >= 0 {
			bitB = int(b[j] - '0')
			j--
		}

		sum := bitA + bitB + carry
		result = string('0'+rune(sum%2)) + result
		carry = sum / 2
	}

	return result
}
