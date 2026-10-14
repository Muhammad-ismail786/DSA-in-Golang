package dsaingolang

func convertToTitle(columnNumber int) string {
	result := ""

	for columnNumber > 0 {
		columnNumber--

		remainder := columnNumber % 26

		result += string(rune('A' + remainder))

		columnNumber = columnNumber / 26
	}

	letters := []rune(result)

	for left, right := 0, len(letters)-1; left < right; left, right = left+1, right-1 {
		letters[left], letters[right] = letters[right], letters[left]
	}

	result = string(letters)

	return result
}
