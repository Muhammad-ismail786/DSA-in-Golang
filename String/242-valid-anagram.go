package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count [26]int

	for _, ch := range s {
		count[ch-'a']++
	}

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

	result := isAnagram(s, t)

	fmt.Println("Output:", result)
}
