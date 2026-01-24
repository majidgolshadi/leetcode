package main

import (
	"log"
	"reflect"
)

func main() {

	tests := []struct {
		input1         []int
		input1_length  int
		input2         []int
		input2_length  int
		expectedOutput []int
	}{
		{
			input1:         []int{1, 2, 3, 0, 0, 0},
			input1_length:  3,
			input2:         []int{2, 5, 6},
			input2_length:  3,
			expectedOutput: []int{1, 2, 2, 3, 5, 6},
		},
		{
			input1:         []int{1},
			input1_length:  1,
			input2:         []int{},
			input2_length:  0,
			expectedOutput: []int{1},
		},
		{
			input1:         []int{0},
			input1_length:  0,
			input2:         []int{1},
			input2_length:  1,
			expectedOutput: []int{1},
		},
		{
			input1:         []int{4, 0, 0, 0, 0, 0},
			input1_length:  1,
			input2:         []int{1, 2, 3, 5, 6},
			input2_length:  5,
			expectedOutput: []int{1, 2, 3, 4, 5, 6},
		},
		{
			input1:         []int{4, 5, 6, 0, 0, 0},
			input1_length:  3,
			input2:         []int{1, 2, 3},
			input2_length:  3,
			expectedOutput: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, test := range tests {
		merge(test.input1, test.input1_length, test.input2, test.input2_length)

		if reflect.DeepEqual(test.expectedOutput, test.input1) {
			log.Fatalf("output is %v but expected is %v", test.input1, test.expectedOutput)
		}
	}

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := 0, m; i < n; i, j = i+1, j+1 {
		nums1[j] = nums2[i]
	}

	for j := 0; j < m+n-1; j++ {
		for i := 0; i < m+n-1; i++ {
			if nums1[i] > nums1[i+1] {
				b := nums1[i]
				nums1[i] = nums1[i+1]
				nums1[i+1] = b
			}
		}
	}
}
