# Intuition

We need to find the smallest missing positive integer from an unsorted array.

A positive integer means a number greater than `0`.

For example:

```go
nums := []int{3, 4, -1, 1}
```

The smallest missing positive integer is `2`.

Instead of sorting the array, we can use **Cyclic Sort** to place each number at its correct index.

For example:

- Number `1` belongs at index `0`.
- Number `2` belongs at index `1`.
- Number `3` belongs at index `2`.
- Number `4` belongs at index `3`.

# Approach

1. Find the length of the array using `len(nums)`.
2. Traverse the array using a `for` loop.
3. If the current number is between `1` and `n`, move it to its correct index.
4. Use swapping to place each number at its correct position.
5. After rearranging, traverse the array again.
6. If `nums[i] != i+1`, return `i+1` because that positive number is missing.
7. If all positions contain their correct numbers, return `n+1`.

# Complexity

- Time complexity: **O(n)**

  Every number is moved to its correct position at most once.

- Space complexity: **O(1)**

  We modify the original array without using an extra array.

# Code

```go
package main

import "fmt"

func firstMissingPositive(nums []int) int {
    n := len(nums)

    for i := 0; i < n; i++ {
        for nums[i] >= 1 &&
            nums[i] <= n &&
            nums[nums[i]-1] != nums[i] {

            correctIndex := nums[i] - 1

            nums[i], nums[correctIndex] = nums[correctIndex], nums[i]
        }
    }

    for i := 0; i < n; i++ {
        if nums[i] != i+1 {
            return i + 1
        }
    }

    return n + 1
}

func main() {
    nums := []int{3, 4, -1, 1}

    result := firstMissingPositive(nums)

    fmt.Println("First Missing Positive:", result)
}
```

**Output:**

```text
First Missing Positive: 2
```
