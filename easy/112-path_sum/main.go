package main

import (
	"log"
)

func main() {

	tests := []struct {
		input          *TreeNode
		targetSum      int
		expectedOutput bool
	}{
		{
			input:          nil,
			targetSum:      0,
			expectedOutput: false,
		},
		{
			input: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			targetSum:      22,
			expectedOutput: true,
		},
		{
			input: &TreeNode{
				Val: -2,
				Right: &TreeNode{
					Val: -3,
				},
			},
			targetSum:      -5,
			expectedOutput: true,
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			targetSum:      5,
			expectedOutput: false,
		},
	}

	for _, test := range tests {
		output := hasPathSum(test.input, test.targetSum)

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
// - The number of nodes in the tree is in the range [0, 5000].
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum = targetSum - root.Val
	hps := false

	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}

	if root.Left != nil {
		hps = hasPathSum(root.Left, targetSum)
	}

	if hps {
		return true
	}

	if root.Right != nil {
		hps = hasPathSum(root.Right, targetSum)
	}

	return hps

}
