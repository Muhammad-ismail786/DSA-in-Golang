# 226. Invert Binary Tree

## Intuition

The problem asks us to invert a binary tree by swapping the left and right children of every node.

We can use recursion to visit every node and swap its left and right children.

## Approach

1. Check if the root is `nil`. If it is, return `nil`.
2. Swap the current node's left and right children using multiple assignment.
3. Recursively invert the left subtree.
4. Recursively invert the right subtree.
5. Finally, return the root of the inverted tree.

## Complexity

- **Time Complexity:** \(O(n)\) — We visit every node exactly once.
- **Space Complexity:** \(O(h)\) — Recursion uses stack space proportional to the height of the tree. In the worst case, it can be \(O(n)\).

## Code

See [226-invert-binary-tree.go](226-invert-binary-tree.go).

## Output

```text
Original Tree: 4 2 1 3 7 6 9
Inverted Tree: 4 7 9 6 2 3 1
```
