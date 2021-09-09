package easy

// Given the head of a singly linked list, return true if it is a palindrome.
//
//
//
// Example 1:
// Input: head = [1,2,2,1]
// Output: true
//
// Example 2:
// Input: head = [1,2]
// Output: false
func isLinkedListPalindrome(head *ListNode) bool {
	var slow, fast = head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow = reverseLinkedList(slow)
	fast = head

	for fast != nil || slow != nil {
		if fast.Val != slow.Val {
			return false
		}

		fast = fast.Next
		slow = slow.Next
	}

	return true
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}

	return prev
}
