
# 52. N-Queens II

**Difficulty:** Hard

**Topics:** Backtracking, Recursion, Hash Map

**Language:** Golang

## # Intuition

The main idea is to place `n` queens on an `n x n` chessboard so that no two queens attack each other.

We place one queen in each row and check whether its column and both diagonals are already occupied.

If a safe position is found, we place the queen and move to the next row using recursion.

Whenever all `n` rows have a queen, we increase the count by 1.

## # Approach

1. Create three maps:
   - `cols` to track occupied columns.
   - `diag1` to track occupied diagonals using `row - col`.
   - `diag2` to track occupied diagonals using `row + col`.

2. Start placing queens from row 0.

3. Check every column to find a safe position.

4. If the position is safe:
   - Mark the column and both diagonals as occupied.
   - Recursively move to the next row.
   - Backtrack by deleting the column and diagonals from their maps.

5. If `row == n`, all queens are successfully placed, so increment `count`.

6. Finally, return the total count.

## # Complexity

- Time complexity: O(N!) approximately, due to backtracking and trying different queen placements.

- Space complexity: O(N) auxiliary space for the three maps, excluding the recursion stack.

## # Code

```go
func totalNQueens(n int) int {

    count := 0

    cols := make(map[int]bool)
    diag1 := make(map[int]bool)
    diag2 := make(map[int]bool)

    var solve func(row int)

    solve = func(row int) {

        if row == n {
            count++
            return
        }

        for col := 0; col < n; col++ {

            d1 := row - col
            d2 := row + col

            if cols[col] || diag1[d1] || diag2[d2] {
                continue
            }

            cols[col] = true
            diag1[d1] = true
            diag2[d2] = true

            solve(row + 1)

            delete(cols, col)
            delete(diag1, d1)
            delete(diag2, d2)
        }
    }

    solve(0)

    return count
}
```

## # Example

**Input:**
```text
n = 4
```

**Output:**
```text
2
```

**Explanation:**

There are two distinct ways to place four queens on a 4 x 4 chessboard so that no two queens attack each other.

## # Tags

- Backtracking
- Recursion
- Hash Map
- LeetCode
```
