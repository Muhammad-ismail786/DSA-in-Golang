package main

import "fmt"

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)

	// Create DP table
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, m+1)

		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	// Recursive helper function
	var isMatchRec func(n, m int) bool

	isMatchRec = func(n, m int) bool {
		// Base case
		if m == 0 {
			return n == 0
		}

		// Already calculated?
		if dp[n][m] != -1 {
			return dp[n][m] == 1
		}

		ans := false

		// Handle '*'
		if p[m-1] == '*' {

			// Case 1: Zero occurrences
			zero := isMatchRec(n, m-2)

			// Case 2: One or more occurrences
			oneOrMore := false

			if n > 0 && (p[m-2] == s[n-1] || p[m-2] == '.') {
				oneOrMore = isMatchRec(n-1, m)
			}

			ans = zero || oneOrMore

		} else {
			// Handle normal character or '.'
			if n > 0 && (s[n-1] == p[m-1] || p[m-1] == '.') {
				ans = isMatchRec(n-1, m-1)
			}
		}

		// Save answer
		if ans {
			dp[n][m] = 1
		} else {
			dp[n][m] = 0
		}

		return ans
	}

	return isMatchRec(n, m)
}

func main() {
	var s, p string

	fmt.Print("Enter string: ")
	fmt.Scan(&s)

	fmt.Print("Enter pattern: ")
	fmt.Scan(&p)

	result := isMatch(s, p)

	fmt.Println("Result:", result)
}
