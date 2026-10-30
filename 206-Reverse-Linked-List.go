/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    current := head

    for current != nil {
        nextNode := current.Next

        current.Next = prev

        prev = current
        current = nextNode
    }

    return prev
}
