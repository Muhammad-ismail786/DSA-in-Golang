package main

import (
	"container/heap"
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Min Heap implementation
type MinHeap []*ListNode

// Number of elements in heap
func (h MinHeap) Len() int {
	return len(h)
}

// Compare elements
func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

// Swap elements
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Add element to heap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

// Remove last element from heap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)

	node := old[n-1]
	old[n-1] = nil
	*h = old[:n-1]

	return node
}

// Merge k sorted linked lists
func mergeKLists(lists []*ListNode) *ListNode {
	// Create Min Heap
	minHeap := &MinHeap{}

	// Initialize heap
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

		// Insert next node from the same list
		if smallest.Next != nil {
			heap.Push(minHeap, smallest.Next)
		}
	}

	return dummy.Next
}

// Create linked list
func createLinkedList(values []int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for _, value := range values {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}

	return dummy.Next
}

// Print linked list
func printLinkedList(head *ListNode) {
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
	// First list: 1 -> 4 -> 5
	list1 := createLinkedList([]int{1, 4, 5})

	// Second list: 1 -> 3 -> 4
	list2 := createLinkedList([]int{1, 3, 4})

	// Third list: 2 -> 6
	list3 := createLinkedList([]int{2, 6})

	// Store all linked lists
	lists := []*ListNode{
		list1,
		list2,
		list3,
	}

	// Merge linked lists
	result := mergeKLists(lists)

	// Print result
	fmt.Print("Merged Linked List: ")
	printLinkedList(result)
}
