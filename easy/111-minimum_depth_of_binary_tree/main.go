package main

import (
	"log"
)

func main() {

	tests := []struct {
		input          *TreeNode
		expectedOutput int
	}{
		{
			input:          nil,
			expectedOutput: 0,
		},
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expectedOutput: 2,
		},
		{
			input: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 6,
							},
						},
					},
				},
			},
			expectedOutput: 5,
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			expectedOutput: 2,
		},
	}

	for _, test := range tests {
		output := minDepth(test.input)

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
// - The number of nodes in the tree is in the range [0, 10^5].
// - 1000 <= Node.val <= 1000

func minDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	minLeft := 0
	if node.Left != nil {
		minLeft = minDepth(node.Left) + 1
	}

	minRight := 0
	if node.Right != nil {
		minRight = minDepth(node.Right) + 1
	}

	if minLeft == 0 {
		return minRight
	}

	if minRight == 0 {
		return minLeft
	}

	if minLeft < minRight {
		return minLeft
	}

	return minRight
}
