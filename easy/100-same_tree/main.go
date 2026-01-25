package main

import (
	"log"
)

func main() {

	tests := []struct {
		input1         *TreeNode
		input2         *TreeNode
		expectedOutput bool
	}{
		{
			input1:         nil,
			input2:         nil,
			expectedOutput: true,
		},
		{
			input1: &TreeNode{
				Val: 1,
			},
			input2:         nil,
			expectedOutput: false,
		},
		{
			input1: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			input2: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expectedOutput: false,
		},
		{
			input1: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
				},
			},
			input2: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 2,
				},
			},
			expectedOutput: false,
		},
	}

	for _, test := range tests {
		output := isSameTree(test.input1, test.input2)

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
//   - The number of nodes in both trees is in the range [0, 100].
//   -10^4 <= Node.val <= 10^4

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
