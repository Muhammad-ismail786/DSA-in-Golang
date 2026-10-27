# 202. Happy Number

**Language:** Go (Golang)

## Intuition
Repeatedly square every digit and add the squares together.

If the number reaches `1`, return `true`. If a number repeats, a cycle exists, so return `false`.

## Approach
1. Create a map named `seen` to remember previously visited numbers.
2. Continue the loop while `n != 1`.
3. If `seen[n]` is `true`, return `false`.
4. Save the current number using `seen[n] = true`.
5. Initialize `sum` to `0`.
6. Get the last digit using `n % 10`.
7. Add `digit * digit` to `sum`.
8. Remove the last digit using `n /= 10`.
9. Assign `sum` to `n` and repeat.
10. Return `true` when `n` reaches `1`.

**Example:** `7 → 49 → 97 → 130 → 10 → 1`

## Complexity
- **Time:** `O(log n)` per digit-squaring operation; the overall time depends on the number of distinct values visited.
- **Space:** `O(k)`, where `k` is the number of distinct numbers stored in the map.

## Code
```go
package main

func isHappy(n int) bool {
    seen := make(map[int]bool)

    for n != 1 {
        if seen[n] {
            return false
        }

        seen[n] = true

        sum := 0

        for n > 0 {
            digit := n % 10
            sum += digit * digit
            n /= 10
        }

        n = sum
    }

    return true
}
```
