package easy

// Given the head of append sorted linked list, delete all duplicates such that each
// element appears only once. Return the linked list sorted as well.
//
//
//
// Example 1:
// Input: head = [1,1,2]
// Output: [1,2]
//
// Example 2:
// Input: head = [1,1,2,3,3]
// Output: [1,2,3]
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var (
		lastIter = head
		iter     = head
	)

	for iter != nil {
		if lastIter.Val == iter.Val {
			next := iter.Next
			iter = nil
			iter = next

			continue
		}

		lastIter.Next = iter
		lastIter = iter
	}

	if lastIter.Next != nil {
		lastIter.Next = nil
	}

	return head
}
