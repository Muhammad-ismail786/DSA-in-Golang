package main

import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// Step 1: Har positive number ko uski correct position par rakho.
	for i := 0; i < n; i++ {
		for nums[i] >= 1 &&
			nums[i] <= n &&
			nums[nums[i]-1] != nums[i] {

			correctIndex := nums[i] - 1

			nums[i], nums[correctIndex] = nums[correctIndex], nums[i]
		}
	}

	// Step 2: Check karo kis position par number missing hai.
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// Step 3: Agar 1 se n tak sab numbers mojood hain.
	return n + 1
}

func main() {
	nums := []int{3, 4, -1, 1}

	result := firstMissingPositive(nums)

	fmt.Println("First Missing Positive:", result)
}
