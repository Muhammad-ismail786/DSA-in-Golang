package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Always perform binary search on the smaller array
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

		// Find boundary values
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

		// Check correct partition
		if l1 <= r2 && l2 <= r1 {

			// Even total elements
			if (m+n)%2 == 0 {
				leftMax := max(l1, l2)
				rightMin := min(r1, r2)

				return (float64(leftMax) + float64(rightMin)) / 2.0
			}

			// Odd total elements
			return float64(max(l1, l2))
		}

		// Move binary search left
		if l1 > r2 {
			high = cut1 - 1
		} else {
			// Move binary search right
			low = cut1 + 1
		}
	}

	return 0.0
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	result := findMedianSortedArrays(nums1, nums2)

	fmt.Println("Median:", result)
}
