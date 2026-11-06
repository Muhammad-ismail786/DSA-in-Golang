# 232. Implement Queue using Stacks (Golang)

## Intuition

A Queue follows FIFO (First In, First Out), which means the element inserted first must be removed first.

A Stack follows LIFO (Last In, First Out), which means the last inserted element is removed first.

We can implement a Queue using two stacks (`s1` and `s2`).

- `s1` is used to insert new elements.
- `s2` is used to access and remove the front element.

When `s2` is empty, we transfer all elements from `s1` to `s2`. This reverses their order and gives us the correct FIFO behavior.

In Go, we use slices (`[]int`) to implement stacks.

## Approach

1. **Push(x):**
   - Insert the new element into `s1` using `append()`.
   - This operation takes O(1) amortized time.

2. **Pop():**
   - If `s2` is empty, transfer all elements from `s1` to `s2`.
   - Remove and return the last element of `s2`.

3. **Peek():**
   - If `s2` is empty, transfer all elements from `s1` to `s2`.
   - Return the last element of `s2` without removing it.

4. **Empty():**
   - Return true only when both `s1` and `s2` are empty.

## Complexity

- **Time complexity:**
  - `Push()`: O(1) amortized
  - `Pop()`: O(1) amortized
  - `Peek()`: O(1) amortized
  - `Empty()`: O(1)
  - Overall: O(n) for n operations.

- **Space complexity:** O(n), where n is the number of elements stored in the Queue.

## Code

```go
type MyQueue struct {
    s1 []int
    s2 []int
}

func Constructor() MyQueue {
    return MyQueue{}
}

func (this *MyQueue) Push(x int) {
    this.s1 = append(this.s1, x)
}

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

func (this *MyQueue) Empty() bool {
    return len(this.s1) == 0 && len(this.s2) == 0
}
