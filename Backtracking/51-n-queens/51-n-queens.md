# LeetCode 51 - N-Queens (Golang)

## # Intuition

The idea is to place `n` queens on an `n x n` chessboard so that no two queens can attack each other.

A queen can attack horizontally, vertically, and diagonally. Therefore, every queen must be placed in a safe position.

We place one queen in each row and try different columns until we find all valid board configurations.

## # Approach

We use **Backtracking** to solve this problem.

1. Create an `n x n` chessboard filled with `.` (empty cells).
2. Start placing queens from the first row.
3. Use the `isSafe()` function to check whether a queen can be placed in a particular column.
4. Check the upper column, upper-left diagonal, and upper-right diagonal.
5. If the position is safe, place the queen (`Q`) and move to the next row.
6. If all queens are placed successfully, convert the board into strings and save it in the answer.
7. Remove the queen (`.`) and try another column. This process is called backtracking.
8. Return all valid board configurations.

## # Complexity

- **Time complexity:** `O(N! × N)` — We explore possible queen placements and check whether each position is safe.

- **Space complexity:** `O(N²)` — To store the chessboard and recursion-related space, excluding the output.

## # Code

```go
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

        for i := 0; i < row; i++ {
            if board[i][col] == 'Q' {
                return false
            }
        }

        i := row - 1
        j := col - 1

        for i >= 0 && j >= 0 {
            if board[i][j] == 'Q' {
                return false
            }
            i--
            j--
        }

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

        if row == n {

            solution := make([]string, n)

            for i := 0; i < n; i++ {
                solution[i] = string(board[i])
            }

            ans = append(ans, solution)
            return
        }

        for col := 0; col < n; col++ {

            if isSafe(row, col) {

                board[row][col] = 'Q'

                backtrack(row + 1)

                board[row][col] = '.'
            }
        }
    }

    backtrack(0)

    return ans
}
```

## # Example

**Input:**
```text
n = 4
```

**Output:**
```text
[
    [".Q..","...Q","Q...","..Q."],
    ["..Q.","Q...","...Q",".Q.."]
]
```

There are two distinct solutions for `n = 4`.
