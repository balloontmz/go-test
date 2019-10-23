/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 *
 * https://leetcode.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (39.07%)
 * Likes:    1146
 * Dislikes: 75
 * Total Accepted:    148.9K
 * Total Submissions: 379.6K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * Write a program to solve a Sudoku puzzle by filling the empty cells.
 *
 * A sudoku solution must satisfy all of the following rules:
 *
 *
 * Each of the digits 1-9 must occur exactly once in each row.
 * Each of the digits 1-9 must occur exactly once in each column.
 * Each of the the digits 1-9 must occur exactly once in each of the 9 3x3
 * sub-boxes of the grid.
 *
 *
 * Empty cells are indicated by the character '.'.
 *
 *
 * A sudoku puzzle...
 *
 *
 * ...and its solution numbers marked in red.
 *
 * Note:
 *
 *
 * The given board contain only digits 1-9 and the character '.'.
 * You may assume that the given Sudoku puzzle will have a single unique
 * solution.
 * The given board size is always 9x9.
 *
 *
 */

// @lc code=start
// 长度为 0 的原因为索引 0 为预留
var rowsUsed [9][10]bool
var colsUsed [9][10]bool
var cubesUsed [9][10]bool

func solveSudoku(board [][]byte) {
	// 6/6 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2 MB)
	// 先看存在哪些数字已经存在
	// 此处循环条件出过问题
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			var num = board[i][j] - '0'
			rowsUsed[i][num] = true
			colsUsed[j][num] = true
			cubesUsed[cubeNum(i, j)][num] = true
		}
	}
	backtracking(0, 0, board)
	// 重新初始化数组,防止变量污染
	var tempRow [9][10]bool
	var tempCol [9][10]bool
	var tempCube [9][10]bool
	rowsUsed = tempRow
	colsUsed = tempCol
	cubesUsed = tempCube
	return
}

func backtracking(row, col int, board [][]byte) bool {
	// 找到下一个未填写的数
	for row < 9 && board[row][col] != '.' {
		if col == 8 {
			row = row + 1
			col = 0
		} else {
			col = col + 1
		}
	}
	if row == 9 {
		return true
	}

	for num := 1; num <= 9; num++ {
		if rowsUsed[row][num] || colsUsed[col][num] || cubesUsed[cubeNum(row, col)][num] {
			continue
		}
		// 此处赋值曾忘记修改
		rowsUsed[row][num] = true
		colsUsed[col][num] = true
		cubesUsed[cubeNum(row, col)][num] = true
		board[row][col] = byte(num + 48)
		if backtracking(row, col, board) {
			return true
		}
		board[row][col] = '.'
		rowsUsed[row][num] = false
		colsUsed[col][num] = false
		cubesUsed[cubeNum(row, col)][num] = false
	}
	return false //
}

func cubeNum(i, j int) int {
	// 这个地方需要抽象下
	r := i / 3 // 行
	c := j / 3 // 列
	return r*3 + c
}

// @lc code=end

