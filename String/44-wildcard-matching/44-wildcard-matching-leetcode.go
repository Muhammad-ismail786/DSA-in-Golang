
func isMatch(s string, p string) bool {

	i := 0
	j := 0

	star := -1
	match := 0

	for i < len(s) {

		// Case 1: Characters match or '?' found
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {

			i++
			j++

		} else if j < len(p) && p[j] == '*' {

			// Case 2: '*' found
			star = j
			match = i
			j++

		} else if star != -1 {

			// Case 3: Previous '*' exists
			j = star + 1
			match++
			i = match

		} else {

			// Case 4: No match possible
			return false
		}
	}

	// Ignore remaining '*' characters
	for j < len(p) && p[j] == '*' {
		j++
	}

	return j == len(p)
}
