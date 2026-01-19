package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	// Create the head node
	head := &ListNode{Val: nums[0]}
	current := head

	// Loop through the rest of the numbers
	for i := 1; i < len(nums); i++ {
		newNode := &ListNode{Val: nums[i]}
		current.Next = newNode // Link current node to new node
		current = newNode      // Move pointer forward
	}
	return head
}

func main() {

	tests := []struct {
		input          *ListNode
		expectedOutput *ListNode
	}{
		{
			input:          createList([]int{}),
			expectedOutput: createList([]int{}),
		},
		{
			input:          createList([]int{1, 1, 2}),
			expectedOutput: createList([]int{1, 2}),
		},
		{
			input:          createList([]int{1,1,2,3,3}),
			expectedOutput: createList([]int{1,2,3}),
		},
	}

	for _, test := range tests {
		output := deleteDuplicates(test.input)

		if test.expectedOutput != output {
			log.Fatalf("input %v, output is %v but expected is %v", test.input, output, test.expectedOutput)
		}
	}

}

/*
* Constraints:
*
*   The number of nodes in the list is in the range [0, 300].
*   -100 <= Node.val <= 100
*   The list is guaranteed to be sorted in ascending order.
*
 */

func deleteDuplicates_versionOne(head *ListNode) *ListNode {
	var prevNodePointer *ListNode
	preVal := 200

	if head == nil {
		return head
	}

	for cursor := head; cursor != nil; cursor = cursor.Next {
		if preVal == cursor.Val {
			prevNodePointer.Next = cursor.Next
		} else {
			prevNodePointer = cursor
		}

		preVal = cursor.Val
	}

	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    current := head

    for current.Next != nil {
        if current.Val == current.Next.Val {
            current.Next = current.Next.Next
        } else {
            current = current.Next
        }
    }

    return head
}
