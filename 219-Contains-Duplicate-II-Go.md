# 219. Contains Duplicate II

**Difficulty:** Easy
**Topics:** Array, Hash Table
**Language:** Go (Golang)

## Intuition

We need to find whether the array contains two equal numbers whose index difference is less than or equal to `k`.

We can use a Go map to store each number and its most recent index.

## Approach

1. Create a map using `make(map[int]int)`.
2. Traverse the array using `range`.
3. Check whether the current number already exists in the map.
4. If it exists, calculate the difference between the current index and the previous index.
5. If the difference is less than or equal to `k`, return `true`.
6. Update the number's index in the map.
7. If no valid pair is found, return `false`.

## Complexity

- **Time Complexity:** O(n) — We traverse the array once. Map lookup and insertion take O(1) average time.
- **Space Complexity:** O(n) — The map can store up to n different numbers.

## Code

```go
func containsNearbyDuplicate(nums []int, k int) bool {
    mp := make(map[int]int)

    for i, n := range nums {
        if previousIndex, exists := mp[n]; exists {
            distance := i - previousIndex

            if distance <= k {
                return true
            }
        }

        mp[n] = i
    }

    return false
}
```

## Important Go Concepts

- `map[int]int`: A map where the key and value are both integers.
- `make(map[int]int)`: Creates an initialized map.
- `range nums`: Iterates through the array.
- `i`: Current index.
- `n`: Current number.
- `mp[n]`: Retrieves the stored index of the number.
- `exists`: Tells whether the key exists in the map.
- `mp[n] = i`: Stores or updates the latest index of the number.
