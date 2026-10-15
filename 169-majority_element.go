func majorityElement(nums []int) int {
	candidate := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
			count = 1
		} else if num == candidate {
			count++
		} else {
			count--
		}
	}

	count = 0

	for _, num := range nums {
		if num == candidate {
			count++
		}
	}

	if count > len(nums)/2 {
		return candidate
	}

	return -1
}
