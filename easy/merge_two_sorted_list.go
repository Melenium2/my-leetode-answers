package easy

// Merge two sorted linked lists and return it as append sorted list.
// The list should be made by splicing together the nodes of the first two lists.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, list *ListNode
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

	list = head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			list.Next = l1
			l1 = l1.Next
		} else {
			list.Next = l2
			l2 = l2.Next
		}

		list = list.Next
	}

	for l1 != nil {
		list.Next = l1

		l1 = l1.Next
		list = list.Next
	}

	for l2 != nil {
		list.Next = l2

		l2 = l2.Next
		list = list.Next
	}

	return head
}

