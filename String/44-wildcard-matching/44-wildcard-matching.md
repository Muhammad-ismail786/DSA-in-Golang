
# 44. Wildcard Matching

**Difficulty:** Hard

**Language:** Go (Golang)

**Approach:** Greedy Algorithm

## Intuition

We need to check whether the given pattern `p` matches the entire string `s`.

There are two special characters:

- `?` matches any single character.
- `*` matches any sequence of characters, including an empty sequence.

We use a Greedy Approach to solve this problem.

We keep track of the last `*` position and use it to match additional characters whenever a mismatch occurs.

## Approach

1. Initialize two pointers `i` and `j` to traverse strings `s` and `p`.
2. Initialize `star = -1` to store the last position of `*`.
3. Initialize `match = 0` to store the position in string `s` where `*` starts matching.
4. If characters match or the pattern contains `?`, increment both pointers.
5. If the pattern contains `*`, save its position in `star` and the current string position in `match`.
6. If a mismatch occurs and a previous `*` exists, let `*` match one additional character and continue checking.
7. If no previous `*` exists, return `false`.
8. After processing the string, skip any remaining `*` characters in the pattern.
9. Return `true` if the entire pattern has been matched; otherwise, return `false`.

## Complexity

- **Time Complexity:** O(n + m) — Linear amortized time using greedy matching.
- **Space Complexity:** O(1) — Only a constant number of variables are used.

## Code

### Go (Golang)

\`\`\`go
func isMatch(s string, p string) bool {

    i := 0
    j := 0

    star := -1
    match := 0

    for i < len(s) {

        if j < len(p) && (s[i] == p[j] || p[j] == '?') {

            i++
            j++

        } else if j < len(p) && p[j] == '*' {

            star = j
            match = i
            j++

        } else if star != -1 {

            j = star + 1
            match++
            i = match

        } else {

            return false
        }
    }

    for j < len(p) && p[j] == '*' {
        j++
    }

    return j == len(p)
}
\`\`\`
