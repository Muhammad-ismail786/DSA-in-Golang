# 231. Power of Two

## Intuition

A power of two is a number that can be divided by 2 repeatedly until it becomes 1.

For example:

16 → 8 → 4 → 2 → 1

If the number becomes 1, it is a power of two.

If the number is less than or equal to 0, it cannot be a power of two.

## Approach

1. First, check if the number is less than or equal to 0. If yes, return false.
2. Use a while loop to divide the number by 2 as long as it is completely divisible by 2.
3. After the loop, check whether the number is equal to 1.
4. If it is 1, return true; otherwise, return false.

## Complexity

- **Time Complexity:** O(log n) — We divide the number by 2 repeatedly.
- **Space Complexity:** O(1) — We use only a constant amount of extra space.

## Code

Language: Golang

```go
package main

import "fmt"

func isPowerOfTwo(n int) bool {
    if n <= 0 {
        return false
    }

    for n%2 == 0 {
        n = n / 2
    }

    return n == 1
}

func main() {
    n := 16

    result := isPowerOfTwo(n)

    fmt.Println("Input:", n)
    fmt.Println("Is Power Of Two:", result)
}
```

## Example

Input: n = 16

Output: true

Explanation:

16 / 2 = 8

8 / 2 = 4

4 / 2 = 2

2 / 2 = 1

Since the final number is 1, it is a power of two.
