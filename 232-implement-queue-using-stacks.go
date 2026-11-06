package main

import "fmt"

type MyQueue struct {
    s1 []int
    s2 []int
}

func Constructor() MyQueue {
    return MyQueue{}
}

// Insert element into Queue
func (this *MyQueue) Push(x int) {
    this.s1 = append(this.s1, x)
}

// Remove and return front element
func (this *MyQueue) Pop() int {
    if len(this.s2) == 0 {
        for len(this.s1) > 0 {
            top := this.s1[len(this.s1)-1]

            this.s2 = append(this.s2, top)
            this.s1 = this.s1[:len(this.s1)-1]
        }
    }

    top := this.s2[len(this.s2)-1]
    this.s2 = this.s2[:len(this.s2)-1]

    return top
}

// Return front element without removing
func (this *MyQueue) Peek() int {
    if len(this.s2) == 0 {
        for len(this.s1) > 0 {
            top := this.s1[len(this.s1)-1]

            this.s2 = append(this.s2, top)
            this.s1 = this.s1[:len(this.s1)-1]
        }
    }

    return this.s2[len(this.s2)-1]
}

// Check whether Queue is empty
func (this *MyQueue) Empty() bool {
    return len(this.s1) == 0 && len(this.s2) == 0
}

func main() {
    myQueue := Constructor()

    var n int

    fmt.Print("How many elements do you want to insert? ")
    fmt.Scan(&n)

    // Insert elements
    for i := 0; i < n; i++ {
        var value int

        fmt.Printf("Enter element %d: ", i+1)
        fmt.Scan(&value)

        myQueue.Push(value)
    }

    // Display front element
    fmt.Println("\nFront element:", myQueue.Peek())

    // Remove front element
    fmt.Println("Removed element:", myQueue.Pop())

    // Display new front element
    if !myQueue.Empty() {
        fmt.Println("Front element after Pop:", myQueue.Peek())
    } else {
        fmt.Println("Queue is empty.")
    }

    // Check Queue status
    fmt.Println("Is Queue empty?", myQueue.Empty())
}
