# 234. Palindrome Linked List

## Intuition

We need to check whether a linked list reads the same forward and backward.

For example:

`1 -> 2 -> 2 -> 1`

This is a palindrome because the values are the same from both directions.

## Approach

1. Use two pointers, `slow` and `fast`, to find the middle of the linked list.
2. Reverse the second half of the linked list.
3. Compare the first half with the reversed second half.
4. If all corresponding values match, return `true`.
5. Otherwise, return `false`.

## Complexity

- **Time Complexity:** O(n) — We traverse the linked list to find the middle, reverse the second half, and compare the values.
- **Space Complexity:** O(1) — We use only a few pointers and no extra data structure.

## Code

```go
func (s Solution) IsPalindrome(head *ListNode) bool {
    slow := head
    fast := head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    var prev *ListNode
    curr := slow

    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    first := head
    second := prev

    for second != nil {
        if first.Val != second.Val {
            return false
        }

        first = first.Next
        second = second.Next
    }

    return true
}
```
