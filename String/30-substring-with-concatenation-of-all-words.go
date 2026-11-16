package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	ans := []int{}

	if len(words) == 0 {
		return ans
	}

	wordLength := len(words[0])
	totalWords := len(words)

	if len(s) < wordLength*totalWords {
		return ans
	}

	wordCount := make(map[string]int)

	// Step 1: Count words
	for _, word := range words {
		wordCount[word]++
	}

	// Step 2: Try different starting positions
	for i := 0; i < wordLength; i++ {

		left := i
		count := 0

		seen := make(map[string]int)

		// Step 3: Move through string
		for right := i; right+wordLength <= len(s); right += wordLength {

			word := s[right : right+wordLength]

			// Step 4: Check if word exists
			if _, exists := wordCount[word]; !exists {

				clear(seen)
				count = 0
				left = right + wordLength

				continue
			}

			seen[word]++
			count++

			// Step 5: Remove extra words
			for seen[word] > wordCount[word] {

				leftWord := s[left : left+wordLength]

				seen[leftWord]--
				count--

				left += wordLength
			}

			// Step 6: All words found
			if count == totalWords {

				ans = append(ans, left)

				leftWord := s[left : left+wordLength]

				seen[leftWord]--
				count--

				left += wordLength
			}
		}
	}

	return ans
}

func main1() {
	var s string
	var n int

	fmt.Print("Enter string: ")
	fmt.Scan(&s)

	fmt.Print("Enter number of words: ")
	fmt.Scan(&n)

	words := make([]string, n)

	fmt.Println("Enter words:")

	for i := 0; i < n; i++ {
		fmt.Scan(&words[i])
	}

	result := findSubstring(s, words)

	fmt.Println("Starting indices:", result)
}
