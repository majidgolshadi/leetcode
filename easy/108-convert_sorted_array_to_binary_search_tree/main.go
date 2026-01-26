package main

import (
	"log"
)

func main() {

	tests := []struct {
		input          []int
		expectedOutput *TreeNode
	}{
		// {
		// 	input:          nil,
		// 	expectedOutput: nil,
		// },
		// {
		// 	input:          []int{},
		// 	expectedOutput: &TreeNode{},
		// },
		{
			input: []int{-10, -3, 0, 5, 9},
			expectedOutput: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -10,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}

	for _, test := range tests {
		output := sortedArrayToBST(test.input)

		if test.expectedOutput != output {
			log.Fatalf("output is %t but expected is %t", output, test.expectedOutput)
		}
	}

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Constraints:
//	- 1 <= nums.length <= 10^4
//  - 10^4 <= nums[i] <= 10^4
//  - nums is sorted in a strictly increasing order.

// A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.
func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}

	mid := len(nums) / 2

	node := &TreeNode{}

	if mid < len(nums) {
		node.Val = nums[mid]

		if mid-1 >= 0 {
			// exclusive
			node.Left = sortedArrayToBST(nums[:mid])
		}

		if mid+1 < len(nums) {
			// inclusive
			node.Right = sortedArrayToBST(nums[mid+1:])
		}
	}

	return node
}
