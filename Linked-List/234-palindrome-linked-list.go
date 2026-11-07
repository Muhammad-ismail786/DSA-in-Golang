package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct{}

func (s Solution) IsPalindrome(head *ListNode) bool {
	// Find the middle of linked list
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the second half
	var prev *ListNode
	curr := slow

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// Compare both halves
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

func main() {
	// Hardcoded input: 1 -> 2 -> 2 -> 1
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}

	solution := Solution{}

	result := solution.IsPalindrome(head)

	fmt.Println("Is Palindrome:", result)
}
