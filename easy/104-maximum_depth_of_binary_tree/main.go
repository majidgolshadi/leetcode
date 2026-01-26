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
					Right: &TreeNode{
						Val: 7,
					},
					Left: &TreeNode{
						Val: 15,
					},
				},
			},
			expectedOutput: 3,
		},
	}

	for _, test := range tests {
		output := maxDepth(test.input)

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
//   - The number of nodes in the tree is in the range [0, 10^4].
//   - 100 <= Node.val <= 100

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}



func maxDepth_solution1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return depth(root, 1)
}

func depth(node *TreeNode, depthLevel int) int {
	if node == nil {
		return depthLevel - 1
	}

	if node.Left == nil {
		return depth(node.Right, depthLevel + 1)	
	}

	if node.Right == nil {
		return depth(node.Left, depthLevel + 1)	
	}

	return maxInt(depth(node.Left, depthLevel + 1), depth(node.Right, depthLevel + 1))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}