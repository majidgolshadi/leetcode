package main

import "log"

func main() {

	tests := []struct {
		input          []int
		target         int
		expectedOutput int
	}{
		{
			input:          []int{1, 2, 5, 6},
			target:         5,
			expectedOutput: 2,
		},
		{
			input:          []int{1, 3, 5, 6},
			target:         2,
			expectedOutput: 1,
		},
		{
			input:          []int{1, 3, 5, 6},
			target:         7,
			expectedOutput: 4,
		},
		{
			input:          []int{},
			target:         1,
			expectedOutput: 0,
		},
		{
			input:          []int{1},
			target:         2,
			expectedOutput: 1,
		},
	}

	for _, test := range tests {
		output := searchInsert(test.input, test.target)

		if test.expectedOutput != output {
			log.Fatalf("input %v, target %d, output is %d", test.input, test.target, output)
		}
	}

}

// algorithm with O(log n) runtime complexity
func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}
