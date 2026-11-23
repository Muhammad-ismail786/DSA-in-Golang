
func solveNQueens(n int) [][]string {

	ans := [][]string{}

	board := make([][]byte, n)

	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)

		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	var isSafe func(row, col int) bool

	isSafe = func(row, col int) bool {

		// Check upper column
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}

		// Check upper-left diagonal
		i := row - 1
		j := col - 1

		for i >= 0 && j >= 0 {
			if board[i][j] == 'Q' {
				return false
			}

			i--
			j--
		}

		// Check upper-right diagonal
		i = row - 1
		j = col + 1

		for i >= 0 && j < n {
			if board[i][j] == 'Q' {
				return false
			}

			i--
			j++
		}

		return true
	}

	var backtrack func(row int)

	backtrack = func(row int) {

		// All queens placed successfully
		if row == n {

			solution := make([]string, n)

			for i := 0; i < n; i++ {
				solution[i] = string(board[i])
			}

			ans = append(ans, solution)

			return
		}

		// Try every column
		for col := 0; col < n; col++ {

			if isSafe(row, col) {

				// Place queen
				board[row][col] = 'Q'

				// Move to next row
				backtrack(row + 1)

				// Backtrack: remove queen
				board[row][col] = '.'
			}
		}
	}

	backtrack(0)

	return ans
}
