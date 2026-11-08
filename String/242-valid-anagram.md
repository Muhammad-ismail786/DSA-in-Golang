
# 242. Valid Anagram

## Intuition

The main idea is to check whether both strings contain the same characters with the same frequency.

- If the lengths of both strings are different, return false.
- Create an array of size 26 to store character frequencies.
- Increase the frequency of each character in string `s`.
- Decrease the frequency of each character in string `t`.
- If any character's frequency becomes negative, return false.
- Otherwise, return true.

## Approach

1. First, check whether the lengths of both strings are equal.
2. Create an integer array `count[26]` initialized with zeros.
3. Traverse string `s` and increase the frequency of each character using `count[ch-'a']++`.
4. Traverse string `t` and decrease the frequency of each character using `count[ch-'a']--`.
5. If any frequency becomes negative, return false.
6. If all characters are processed successfully, return true.

## Complexity

- Time complexity: O(n)

  We traverse both strings once, where n is the length of the strings.

- Space complexity: O(1)

  We use a fixed-size array of 26 elements.

## Code

```go
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

    if isAnagram(s, t) {
        fmt.Println("Output: true")
    } else {
        fmt.Println("Output: false")
    }
}
```

## Example 1

Input:
```text
anagram
nagaram
```

Output:
```text
true
```

## Example 2

Input:
```text
rat
car
```

Output:
```text
false
```
