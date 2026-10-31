# 217. Contains Duplicate

## Intuition

We need to check whether the given array contains any duplicate numbers.

We can use a Go map to keep track of numbers we have already visited. If a number is already present in the map, it means a duplicate exists, so we return `true`.

## Approach

1. Create an empty map named `seen` using `make(map[int]bool)`.
2. Iterate through each number in the slice using a for loop and range.
3. Use `_` to ignore the index because we only need the number.
4. Check whether the current number already exists in the map using `seen[n]`.
5. If its value is `true`, return `true` because a duplicate is found.
6. Otherwise, store the number in the map using `seen[n] = true`.
7. If the loop finishes without finding any duplicate, return `false`.

## Complexity

- **Time Complexity:** O(n) average — We traverse the slice once. Map lookup and insertion take O(1) average time.
- **Space Complexity:** O(n) — In the worst case, the map stores all unique numbers.

## Code

```go
func containsDuplicate(nums []int) bool {
    seen := make(map[int]bool)

    for _, n := range nums {
        if seen[n] {
            return true
        }

        seen[n] = true
    }

    return false
}
```
