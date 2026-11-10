# 2. Add Two Numbers

**LeetCode:** https://leetcode.com/problems/add-two-numbers/

**Difficulty:** Medium

**Topic:** Linked List

## Intuition

We are given two numbers represented by two linked lists.

Each node contains one digit, and the digits are stored in reverse order.

We need to add both numbers and return the result as a linked list.

For example:

```text
l1 = 2 -> 4 -> 3
l2 = 5 -> 6 -> 4
```

These linked lists represent:

```text
342 + 465 = 807
```

Our answer will be:

```text
7 -> 0 -> 8
```

We use a `carry` variable to store the extra value when the sum is greater than or equal to 10.

## Approach

1. Create a dummy node to make the result linked list.
2. Create a `current` pointer to track the last node.
3. Initialize `carry = 0`.
4. Use a loop while either linked list has nodes or carry is not zero.
5. Start `sum` with the carry value.
6. If `l1` is not nil, add its value to sum and move to the next node.
7. If `l2` is not nil, add its value to sum and move to the next node.
8. If sum is greater than or equal to 10, set carry to 1; otherwise, set it to 0.
9. Calculate the digit using `sum % 10`.
10. Create a new node with that digit.
11. Move the current pointer to the newly created node.
12. Return `dummy.Next`, which is the actual result list.

## Example Dry Run

Input:

```text
l1 = 2 -> 4 -> 3
l2 = 5 -> 6 -> 4
```

| Round | l1 | l2 | Carry | Sum | Digit |
|---|---:|---:|---:|---:|---:|
| 1 | 2 | 5 | 0 | 7 | 7 |
| 2 | 4 | 6 | 0 | 10 | 0 |
| 3 | 3 | 4 | 1 | 8 | 8 |

Output:

```text
7 -> 0 -> 8
```

## Complexity

**Time Complexity:** O(max(n, m))

We visit each node of both linked lists once.

**Space Complexity:** O(max(n, m))

We create a new linked list to store the result.

## Code

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    current := dummy
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        sum := carry

        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        if sum >= 10 {
            carry = 1
        } else {
            carry = 0
        }

        digit := sum % 10

        current.Next = &ListNode{Val: digit}
        current = current.Next
    }

    return dummy.Next
}
```
