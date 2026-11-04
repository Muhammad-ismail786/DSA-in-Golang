package main

import (
	"encoding/json"
	"fmt"
)

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

func main() {
	nums := []int{2, 3, 4, 7, 9, 10}

	result := summaryRanges(nums)

	output, _ := json.Marshal(result)

	fmt.Println(string(output))
}
