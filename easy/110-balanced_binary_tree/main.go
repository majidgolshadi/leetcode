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
			expectedOutput: true,
		},
		{
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
					Left: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
						},
						Left: &TreeNode{
							Val: 4,
						},
					},
				},
			},
			expectedOutput: false,
		},
	}

	for _, test := range tests {
		output := isBalanced(test.input)

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
func isBalanced(root *TreeNode) bool {
    // If checkHeight doesn't return -1, the tree is balanced
    return checkHeight(root) != -1
}

func checkHeight(node *TreeNode) int {
    if node == nil {
        return 0
    }
	
    leftHeight := checkHeight(node.Left)
    if leftHeight == -1 {
        return -1
    }

	
    rightHeight := checkHeight(node.Right)
    if rightHeight == -1 {
        return -1
    }

    // 1. Check if current node is balanced
    // 2. Return actual height if balanced, else -1
    diff := leftHeight - rightHeight
    if diff < 0 {
        diff = -diff
    }

    if diff > 1 {
        return -1
    }

    // Return the height of this node
    if leftHeight > rightHeight {
        return leftHeight + 1
    }

    return rightHeight + 1
}
