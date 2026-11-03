package main

import "fmt"

// Step 1: TreeNode structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Step 2: Invert Binary Tree function
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// Step 3: Tree print karne ka function
func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Print(root.Val, " ")

	printTree(root.Left)
	printTree(root.Right)
}

// Step 4: Main function
func main() {
	root := &TreeNode{Val: 4}

	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}

	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}

	fmt.Print("Original Tree: ")
	printTree(root)

	root = invertTree(root)

	fmt.Print("\nInverted Tree: ")
	printTree(root)
}
