
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	var previous *ListNode = nil
	current := head

	for current != nil {
		if current.Val == val {
			if previous == nil {
				head = current.Next
			} else {
				previous.Next = current.Next
			}
		} else {
			previous = current
		}

		current = current.Next
	}

	return head
}
