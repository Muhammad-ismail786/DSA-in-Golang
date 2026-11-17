
# LeetCode 32: Longest Valid Parentheses

## # Intuition

We are given a string containing only opening `(` and closing `)` parentheses.

We need to find the length of the longest valid parentheses substring.

We can solve this problem using a Stack.

Instead of storing parentheses, we store their indexes in the stack.

We initialize the stack with `-1` as the starting boundary.

## # Approach

1. Create a stack and initialize it with `-1`.
2. Initialize an answer variable with `0`.
3. Traverse the string using a loop.
4. If the current character is `(`, push its index into the stack.
5. Otherwise, remove the top index from the stack.
6. If the stack becomes empty, push the current index as a new boundary.
7. Otherwise, calculate the valid substring length using `i - stack[len(stack)-1]`.
8. Update the maximum length.
9. Return the final answer.

## # Complexity

- Time complexity: `O(n)`

  We traverse the string only once.

- Space complexity: `O(n)`

  In the worst case, the stack can store up to `n` indexes.

## # Code

```go
func longestValidParentheses(s string) int {

    stack := []int{-1}
    ans := 0

    for i := 0; i < len(s); i++ {

        if s[i] == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1]
        }

        if len(stack) == 0 {
            stack = append(stack, i)
        } else {
            length := i - stack[len(stack)-1]

            if length > ans {
                ans = length
            }
        }
    }

    return ans
}
```
