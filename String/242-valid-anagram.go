package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count [26]int

	// Count characters in s
	for _, ch := range s {
		count[ch-'a']++
	}

	// Decrease characters in t
	for _, ch := range t {
		count[ch-'a']--

		if count[ch-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {
	var s, t string

	fmt.Print("Enter first string: ")
	fmt.Scan(&s)

	fmt.Print("Enter second string: ")
	fmt.Scan(&t)

	if isAnagram(s, t) {
		fmt.Println("Output: true")
	} else {
		fmt.Println("Output: false")
	}
}
