package main

import "fmt"

// Function to check whether a number is a power of two
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	for n%2 == 0 {
		n = n / 2
	}

	return n == 1
}

func main() {
	// Input data
	n := 16

	// Function call
	result := isPowerOfTwo(n)

	// Print output
	fmt.Println("Input:", n)
	fmt.Println("Is Power Of Two:", result)
}
