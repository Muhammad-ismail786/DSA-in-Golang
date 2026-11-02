
# 225. Implement Stack using Queues

**Difficulty:** Easy

**Topics:** Stack, Queue, Design

## Intuition

We need to implement a Stack using two Queues.

A Stack follows LIFO (Last In, First Out), which means the last inserted element is removed first.

A Queue follows FIFO (First In, First Out), which means the first inserted element is removed first.

We use two slices (`q1` and `q2`) to simulate Queues in Go.

Whenever we insert a new element, we add it to `q2` first, then move all elements from `q1` to `q2`. Finally, we swap both slices so that the newest element always stays at the front of `q1`.

## Approach

1. **Push(x):**
   - Insert `x` into `q2` using `append()`.
   - Move all elements from `q1` to `q2`.
   - Remove moved elements from `q1` using slicing.
   - Swap `q1` and `q2`.
   - The newest element is now at the front of `q1`.

2. **Pop():**
   - Get the front element of `q1`.
   - Remove it using slicing.
   - Return the removed element.

3. **Top():**
   - Return the front element of `q1` without removing it.

4. **Empty():**
   - Check whether `q1` has zero elements using `len()`.

## Complexity

- **Time Complexity:**
  - Push: O(n)
  - Pop: O(1)
  - Top: O(1)
  - Empty: O(1)

- **Space Complexity:** O(n)

## Code

```go
type MyStack struct {
    q1 []int
    q2 []int
}

func Constructor() MyStack {
    return MyStack{}
}

func (this *MyStack) Push(x int) {
    this.q2 = append(this.q2, x)

    for len(this.q1) > 0 {
        this.q2 = append(this.q2, this.q1[0])
        this.q1 = this.q1[1:]
    }

    this.q1, this.q2 = this.q2, this.q1
}

func (this *MyStack) Pop() int {
    x := this.q1[0]
    this.q1 = this.q1[1:]
    return x
}

func (this *MyStack) Top() int {
    return this.q1[0]
}

func (this *MyStack) Empty() bool {
    return len(this.q1) == 0
}
```
