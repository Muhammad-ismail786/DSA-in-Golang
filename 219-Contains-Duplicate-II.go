
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
