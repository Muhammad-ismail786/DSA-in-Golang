# LeetCode 25 — Reverse Nodes in k-Group (Go)

## Intuition

We need to reverse the linked list in groups of `k` nodes.

If fewer than `k` nodes remain, we leave those nodes unchanged.

We change the links (`Next`) between nodes. We do not change node values.

## Approach

1. First, check whether `k` nodes are available.
2. If fewer than `k` nodes are available, return the original head.
3. Reverse the first `k` nodes using three pointers: `prev`, `curr`, and `nextNode`.
4. Connect the original head to the recursively reversed remaining list.
5. Return `prev`, which is the new head of the reversed group.

## Complexity

- **Time complexity:** O(n) — Every node is visited a constant number of times.
- **Space complexity:** O(n/k) — Recursive call stack for each group. In the worst case, O(n) when k = 1.

## Code — Go

```go
func reverseKGroup(head *ListNode, k int) *ListNode {

    // Step 1: Check if k nodes are available
    temp := head

    for i := 0; i < k; i++ {
        if temp == nil {
            return head
        }

        temp = temp.Next
    }

    // Step 2: Reverse k nodes
    var prev *ListNode
    curr := head

    for i := 0; i < k; i++ {
        nextNode := curr.Next

        curr.Next = prev

        prev = curr
        curr = nextNode
    }

    // Step 3: Connect remaining nodes
    head.Next = reverseKGroup(curr, k)

    // Step 4: Return new head
    return prev
}
```

## Example

**Input:**
```text
head = [1,2,3,4,5]
k = 2
```

**Output:**
```text
[2,1,4,3,5]
```

**Explanation:**

- First group: `[1,2]` becomes `[2,1]`.
- Second group: `[3,4]` becomes `[4,3]`.
- Remaining node: `[5]` stays unchanged.
