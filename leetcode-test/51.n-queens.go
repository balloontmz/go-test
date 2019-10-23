/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 *
 * https://leetcode.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (41.74%)
 * Likes:    1222
 * Dislikes: 55
 * Total Accepted:    163.4K
 * Total Submissions: 389.4K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n×n chessboard
 * such that no two queens attack each other.
 *
 *
 *
 * Given an integer n, return all distinct solutions to the n-queens puzzle.
 *
 * Each solution contains a distinct board configuration of the n-queens'
 * placement, where 'Q' and '.' both indicate a queen and an empty space
 * respectively.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: [
 * ⁠[".Q..",  // Solution 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 *
 * ⁠["..Q.",  // Solution 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * Explanation: There exist two distinct solutions to the 4-queens puzzle as
 * shown above.
 *
 *
 */

// @lc code=start

type nQueens struct {
	solutions        [][]string
	nQueens          [][]byte
	colUsed          []bool
	diagonals45Used  []bool
	diagonals135Used []bool
	n                int
}

func solveNQueens(n int) [][]string {
	// 9/9 cases passed (4 ms)
	// Your runtime beats 97.88 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (8.9 MB)
	var solution nQueens
	solution.solutions = make([][]string, 0)
	solution.colUsed = make([]bool, n)
	solution.diagonals45Used = make([]bool, 2*n-1)
	solution.diagonals135Used = make([]bool, 2*n-1)
	solution.n = n

	var tempNQueen = make([][]byte, n)
	for i := 0; i < n; i++ {
		var temp = make([]byte, n)
		tempNQueen[i] = temp
		for j := 0; j < n; j++ {
			temp[j] = '.'
		}
	}

	solution.nQueens = tempNQueen

	solution.backtracking(0)

	return solution.solutions
}

func (s *nQueens) backtracking(row int) {
	if row == s.n {
		var tempList []string
		for _, v := range s.nQueens {
			tempList = append(tempList, string(v))
		}
		s.solutions = append(s.solutions, tempList)
		return
	}
	for col := 0; col < s.n; col++ {
		var diagonals45Idx = row + col
		var diagonals135Idx = s.n - 1 - (row - col)
		if s.colUsed[col] || s.diagonals45Used[diagonals45Idx] || s.diagonals135Used[diagonals135Idx] {
			continue
		}
		s.nQueens[row][col] = 'Q'
		s.colUsed[col] = true
		s.diagonals45Used[diagonals45Idx] = true
		s.diagonals135Used[diagonals135Idx] = true
		s.backtracking(row + 1)
		s.colUsed[col] = false
		s.diagonals45Used[diagonals45Idx] = false
		s.diagonals135Used[diagonals135Idx] = false
		s.nQueens[row][col] = '.'
	}
}

// @lc code=end

