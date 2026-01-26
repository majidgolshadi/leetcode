package main

import (
	"log"
)

func main() {

	tests := []struct {
		input          *TreeNode
		expectedOutput bool
	}{
		{
			input:          nil,
			expectedOutput: true,
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
					Left: &TreeNode{
						Val: 4,
					},
				},
			},
			expectedOutput: true,
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expectedOutput: false,
		},
	}

	for _, test := range tests {
		output := isSymmetric(test.input)

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
//   - The number of nodes in the tree is in the range [1, 1000].
//   - 100 <= Node.val <= 100

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val &&
		isMirror(p.Left, q.Right) &&
		isMirror(p.Right, q.Left)
}
