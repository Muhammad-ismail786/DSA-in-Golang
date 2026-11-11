
# 4. Median of Two Sorted Arrays

**Difficulty:** Hard

**Topics:** Array, Binary Search, Divide and Conquer

## Intuition

We are given two sorted arrays.

We need to find the median without merging both arrays.

The main idea is to divide both arrays into two parts:

- Left half
- Right half

The left half should contain half of the total elements.

All elements in the left half must be smaller than or equal to all elements in the right half.

We use Binary Search on the smaller array to find the correct partition.

## Approach

1. First, make sure nums1 is the smaller array.
2. Calculate the sizes of both arrays.
3. Initialize low = 0 and high = m.
4. Use Binary Search to calculate cut1.
5. Calculate cut2 so that the left partition contains half of the total elements.
6. Find four boundary values:
   - l1: Last element on the left side of nums1.
   - r1: First element on the right side of nums1.
   - l2: Last element on the left side of nums2.
   - r2: First element on the right side of nums2.
7. Check whether the partition is correct:
   l1 <= r2 and l2 <= r1
8. If the total number of elements is even, return the average of the two middle elements.
9. If the total number of elements is odd, return the maximum element from the left partition.
10. If the partition is incorrect, adjust the Binary Search boundaries.

## Complexity

- Time complexity: O(log(min(m, n)))
- Space complexity: O(1)

## Code

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

    if len(nums1) > len(nums2) {
        return findMedianSortedArrays(nums2, nums1)
    }

    m := len(nums1)
    n := len(nums2)

    low := 0
    high := m

    for low <= high {

        cut1 := low + (high-low)/2
        cut2 := (m+n+1)/2 - cut1

        l1 := -int(^uint(0)>>1) - 1
        if cut1 > 0 {
            l1 = nums1[cut1-1]
        }

        r1 := int(^uint(0) >> 1)
        if cut1 < m {
            r1 = nums1[cut1]
        }

        l2 := -int(^uint(0)>>1) - 1
        if cut2 > 0 {
            l2 = nums2[cut2-1]
        }

        r2 := int(^uint(0) >> 1)
        if cut2 < n {
            r2 = nums2[cut2]
        }

        if l1 <= r2 && l2 <= r1 {

            if (m+n)%2 == 0 {
                leftMax := max(l1, l2)
                rightMin := min(r1, r2)

                return (float64(leftMax) + float64(rightMin)) / 2.0
            }

            return float64(max(l1, l2))
        }

        if l1 > r2 {
            high = cut1 - 1
        } else {
            low = cut1 + 1
        }
    }

    return 0.0
}
```
