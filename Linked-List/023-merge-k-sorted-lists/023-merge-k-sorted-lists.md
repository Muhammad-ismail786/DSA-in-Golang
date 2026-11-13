
# 23. Merge k Sorted Lists

**Difficulty:** Hard

**Topic:** Linked List, Heap, Priority Queue

**Language:** Go (Golang)

## Problem Description

You are given an array of `k` linked lists, where each linked list is sorted in ascending order.

Merge all the linked lists into one sorted linked list and return its head.

### Example

**Input:**

```text
lists = [
    1 -> 4 -> 5,
    1 -> 3 -> 4,
    2 -> 6
]
```

**Output:**

```text
1 -> 1 -> 2 -> 3 -> 4 -> 4 -> 5 -> 6
```

## Intuition

We are given `k` sorted linked lists. We need to merge all of them into one sorted linked list.

Instead of comparing every node manually, we can use a Min Heap to efficiently find the smallest available node.

In Go, we can implement a Min Heap using the `container/heap` package.

Initially, we insert the first node of every non-empty linked list into the Min Heap.

The Min Heap always gives us the smallest node. We remove that node, attach it to our result list, and insert its next node into the heap.

We repeat this process until the heap becomes empty.

## Approach

1. Import the `container/heap` package.
2. Create a `MinHeap` type using a slice of `*ListNode`.
3. Implement the required heap methods:
   - `Len()` returns the number of elements.
   - `Less()` compares node values.
   - `Swap()` swaps two elements.
   - `Push()` inserts a node.
   - `Pop()` removes a node.
4. Initialize the Min Heap.
5. Insert the first node of every non-empty linked list.
6. Create a dummy node to build the merged linked list.
7. While the heap is not empty:
   - Remove the smallest node using `heap.Pop()`.
   - Attach it to the merged linked list.
   - Move the current pointer forward.
   - Insert its next node into the heap if it exists.
8. Return `dummy.Next`.

## Complexity

- **Time Complexity:** O(N log k)
  - `N` is the total number of nodes across all linked lists.
  - `k` is the number of linked lists.
  - Every node is inserted into and removed from the heap once.

- **Space Complexity:** O(k)
  - The Min Heap stores at most one available node from each linked list at a time.

## Code

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "container/heap"

type MinHeap []*ListNode

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)

    node := old[n-1]
    old[n-1] = nil
    *h = old[:n-1]

    return node
}

func mergeKLists(lists []*ListNode) *ListNode {

    // Create Min Heap
    minHeap := &MinHeap{}
    heap.Init(minHeap)

    // Insert first node of every non-empty list
    for _, list := range lists {
        if list != nil {
            heap.Push(minHeap, list)
        }
    }

    // Create dummy node
    dummy := &ListNode{}
    current := dummy

    // Merge all linked lists
    for minHeap.Len() > 0 {

        // Get smallest node
        smallest := heap.Pop(minHeap).(*ListNode)

        // Attach node to result
        current.Next = smallest
        current = current.Next

        // Insert next node
        if smallest.Next != nil {
            heap.Push(minHeap, smallest.Next)
        }
    }

    return dummy.Next
}
```

## Key Concepts

| Concept | Meaning |
|---|---|
| `container/heap` | Go package for heap operations |
| `MinHeap` | Custom Min Heap implementation |
| `heap.Init()` | Initializes the heap |
| `heap.Push()` | Inserts an element |
| `heap.Pop()` | Removes the smallest element |
| `dummy` | Starting node used to build the result |
| `current` | Pointer to the last node in the result |

