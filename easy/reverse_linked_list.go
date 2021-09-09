package easy

// Given the head of a singly linked list, reverse the list, and return the reversed list.
//
//
//
// Example 1:
//
//
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
// Example 2:
//
//
// Input: head = [1,2]
// Output: [2,1]
// Example 3:
//
// Input: head = []
// Output: []
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		nextNode := head.Next
		head.Next = prev
		prev = head
		head = nextNode
	}

	return prev
}
