# 30. Substring with Concatenation of All Words

**LeetCode Problem:** [30. Substring with Concatenation of All Words](https://leetcode.com/problems/substring-with-concatenation-of-all-words/)

## # Intuition

We need to find all starting indices where the substring contains every word from the `words` array exactly once, in any order.

Since all words have the same length, we can use a Sliding Window and Go maps to count words efficiently.

## # Approach

1. First, calculate the length of each word and the total number of words.
2. Create a map called `wordCount` to store how many times each word appears in the `words` slice.
3. Try different starting positions from `0` to `wordLength - 1`.
4. Use two pointers, `left` and `right`, to maintain a sliding window.
5. Extract a word of `wordLength` from the current `right` position.
6. If the extracted word does not exist in `wordCount`, reset the `seen` map and move `left` forward.
7. If the word exists, increase its count in `seen`.
8. If a word appears more times than allowed, move the `left` pointer forward and remove extra words.
9. When the window contains all required words, add the `left` index to the answer.
10. Finally, return the answer slice.

## # Complexity

- **Time complexity:** \(O(n \times m)\)

  Where `n` is the length of string `s` and `m` is the length of each word. This accounts for substring hashing and map operations.

- **Space complexity:** \(O(k)\)

  Where `k` is the number of unique words stored in the hash maps, excluding the input strings.

## # Code

```go
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

                seen = make(map[string]int)
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
```
