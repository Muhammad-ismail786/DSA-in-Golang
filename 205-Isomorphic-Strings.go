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
