# 205. Isomorphic Strings

## Intuition
We need to check whether characters in string `s` follow the same pattern as characters in string `t`. Each character must consistently map to one character in the other string.

## Approach
1. If the lengths of both strings are different, return `false`.
2. Create two arrays, `mapS` and `mapT`, each containing 256 integers initialized to zero.
3. Loop through both strings using index `i`.
4. Compare the last-seen positions of the current characters. If they differ, return `false`.
5. Update both arrays with `i + 1` to record the current position.
6. If all characters match the same pattern, return `true`.

## Complexity
- Time complexity: `O(n)`
- Space complexity: `O(1)`

## Code

```go
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    mapS := [256]int{}
    mapT := [256]int{}

    for i := 0; i < len(s); i++ {
        if mapS[s[i]] != mapT[t[i]] {
            return false
        }

        mapS[s[i]] = i + 1
        mapT[t[i]] = i + 1
    }

    return true
}
```
