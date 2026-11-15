package main

import "fmt"

// Definition for singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

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

		// Save next node
		nextNode := curr.Next

		// Reverse current node's link
		curr.Next = prev

		// Move pointers forward
		prev = curr
		curr = nextNode
	}

	// Step 3: Connect remaining nodes
	head.Next = reverseKGroup(curr, k)

	// Step 4: Return new head
	return prev
}

// Print linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)

		if head.Next != nil {
			fmt.Print(" -> ")
		}

		head = head.Next
	}

	fmt.Println()
}

// Main function
func main() {
	var n int

	fmt.Print("Enter number of nodes: ")
	fmt.Scan(&n)

	var head *ListNode
	var tail *ListNode

	fmt.Println("Enter node values:")

	for i := 0; i < n; i++ {

		var value int
		fmt.Scan(&value)

		newNode := &ListNode{
			Val: value,
		}

		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}

	var k int

	fmt.Print("Enter k: ")
	fmt.Scan(&k)

	fmt.Print("\nOriginal List: ")
	printList(head)

	head = reverseKGroup(head, k)

	fmt.Print("After Reversing: ")
	printList(head)
}
