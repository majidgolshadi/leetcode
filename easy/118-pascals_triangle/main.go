package main

import (
	"log"
	"reflect"
)

func main() {

	tests := []struct {
		numRows        int
		expectedOutput [][]int
	}{
		{
			numRows:        5,
			expectedOutput: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			numRows:        1,
			expectedOutput: [][]int{{1}},
		},
	}

	for _, test := range tests {
		output := generate(test.numRows)

		if reflect.DeepEqual(test.expectedOutput, output) {
			log.Fatalf("output is %t but expected is %t", output, test.expectedOutput)
		}
	}

}

// Constraints:
// - 1 <= numRows <= 30

func generate_solution1(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}

	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	result := make([][]int, numRows)
	result[0] = []int{1}
	result[1] = []int{1, 1}

	for i := 2; i < numRows; i++ {
		rowLength := len(result[i-1]) + 1
		row := make([]int, rowLength)
		
		for j := 0; j < rowLength; j++ {
			if j == 0 || j == rowLength - 1{
				row[j] = 1
				continue
			}

			row[j] = result[i-1][j-1] + result[i-1][j]
		}

		result[i] = row
	}

	return result
}

func generate(numRows int) [][]int {
    triangle := make([][]int, numRows)

    for i := 0; i < numRows; i++ {
        // 1. Create a row with (i + 1) elements, all set to 1
        row := make([]int, i+1)
        row[0], row[i] = 1, 1

        // 2. Fill in the middle elements (if any)
        for j := 1; j < i; j++ {
            row[j] = triangle[i-1][j-1] + triangle[i-1][j]
        }

        triangle[i] = row
    }

    return triangle
}