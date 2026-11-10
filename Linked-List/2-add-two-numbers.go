package main

import "fmt"

// Definition of Linked List Node
type ListNode struct {
	Val  int
	Next *ListNode
}

// Add Two Numbers
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

// Create Linked List
func createList(values []int) *ListNode {
	dummy := &ListNode{Val: 0}
	current := dummy

	for _, value := range values {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}

	return dummy.Next
}

// Print Linked List
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

// Delete Linked List
func deleteList(head *ListNode) {
	for head != nil {
		temp := head
		head = head.Next
		temp.Next = nil
	}
}

// Main Function
func main() {
	// First Number: 342
	l1 := createList([]int{2, 4, 3})

	// Second Number: 465
	l2 := createList([]int{5, 6, 4})

	fmt.Print("First Number: ")
	printList(l1)

	fmt.Print("Second Number: ")
	printList(l2)

	// Add Two Numbers
	result := addTwoNumbers(l1, l2)

	fmt.Print("Result: ")
	printList(result)

	// Clean up
	deleteList(l1)
	deleteList(l2)
	deleteList(result)
}
