package main

import (
	"log"
	"reflect"
)

func main() {

	tests := []struct {
		input          *TreeNode
		expectedOutput []int
	}{
		{
			input:          nil,
			expectedOutput: nil,
		},
		{
			input: &TreeNode{
				Val: 1,
			},
			expectedOutput: []int{1},
		},
		{
			input: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			expectedOutput: []int{1, 3, 2},
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 7,
						},
						Left: &TreeNode{
							Val: 6,
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 9,
						},
					},
				},
			},
			expectedOutput: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
	}

	for _, test := range tests {
		output := inorderTraversal(test.input)

		if reflect.DeepEqual(test.expectedOutput, output) {
			log.Fatalf("output is %v but expected is %v for input %v", output, test.expectedOutput, test.input)
		}
	}

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Constraints:
//  * The number of nodes in the tree is in the range [0, 100].
//  * -100 <= Node.val <= 100

// left -> root -> right
func inorderTraversal(root *TreeNode) []int {
	var result []int
	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)
		result = append(result, root.Val)
		traverse(root.Right)
	}

	traverse(root)

	return result
}
