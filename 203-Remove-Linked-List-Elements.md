
# 203. Remove Linked List Elements

## Intuition
We need to remove all nodes from the linked list whose value equals `val`.

We can traverse the linked list and remove matching nodes by updating pointers.

## Approach
1. Initialize `current` to `head` and `previous` to `nil`.
2. Traverse the linked list using a `for` loop.
3. If `current.Val == val`:
   - If `previous == nil`, update `head` to `current.Next`.
   - Otherwise, update `previous.Next` to `current.Next`.
4. If the value does not match, move `previous` to `current`.
5. Move `current` to the next node.
6. Return the updated `head`.

## Complexity

- **Time Complexity:** O(n), where n is the number of nodes.
- **Space Complexity:** O(1), because we use only a constant number of pointers.

## Code

```go
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
```
