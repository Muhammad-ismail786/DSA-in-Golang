package main

import "fmt"

func solveSudoku(board [][]byte) {
	var solve func() bool

	solve = func() bool {
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if board[row][col] == '.' {

					for num := byte('1'); num <= '9'; num++ {

						valid := true

						for i := 0; i < 9; i++ {

							if board[row][i] == num ||
								board[i][col] == num {

								valid = false
								break
							}

							r := 3*(row/3) + i/3
							c := 3*(col/3) + i%3

							if board[r][c] == num {
								valid = false
								break
							}
						}

						if valid {

							board[row][col] = num

							if solve() {
								return true
							}

							board[row][col] = '.'
						}
					}

					return false
				}
			}
		}

		return true
	}

	solve()
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board)

	fmt.Println("Solved Sudoku:")
	fmt.Println()

	for _, row := range board {

		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}

		fmt.Println()
	}
}
