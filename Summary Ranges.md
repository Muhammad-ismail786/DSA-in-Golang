# 228. Summary Ranges – Go Solution

## Intuition

The main idea is to find consecutive numbers in the sorted array and group them into ranges.

If numbers are consecutive, we keep them in the same range. If there is a gap between numbers, we finish the current range and start a new one.

Example:

`nums = [2, 3, 4, 7, 9, 10]`

The consecutive ranges are:

- `2, 3, 4` → `"2->4"`
- `7` → `"7"`
- `9, 10` → `"9->10"`

## Approach

1. Create an empty `result` slice to store the final ranges.
2. If the input array is empty, return the empty result.
3. Initialize `start = 0` to keep track of the starting index of the current range.
4. Traverse the array starting from index `1`.
5. Check whether the current number is consecutive with the previous number using `nums[i] != nums[i-1]+1`.
6. If there is a gap:
   - If `start == i-1`, only one number exists in the range, so add that number to the result.
   - Otherwise, add the range in `"start->end"` format.
   - Update `start = i` to begin a new range.
7. After the loop, add the last remaining range to the result.
8. Return the result slice.

## Complexity

- **Time complexity:** `O(n)` — We traverse the input array once.
- **Space complexity:** `O(1)` auxiliary space, excluding the output slice. Including the output slice, total space can be `O(n)`.

## Code

```go
func summaryRanges(nums []int) []string {
    result := []string{}

    if len(nums) == 0 {
        return result
    }

    start := 0

    for i := 1; i < len(nums); i++ {

        if nums[i] != nums[i-1]+1 {

            if start == i-1 {
                result = append(result, fmt.Sprintf("%d", nums[start]))
            } else {
                result = append(
                    result,
                    fmt.Sprintf("%d->%d", nums[start], nums[i-1]),
                )
            }

            start = i
        }
    }

    if start == len(nums)-1 {
        result = append(result, fmt.Sprintf("%d", nums[start]))
    } else {
        result = append(
            result,
            fmt.Sprintf("%d->%d", nums[start], nums[len(nums)-1]),
        )
    }

    return result
}
```
