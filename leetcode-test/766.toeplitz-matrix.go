/*
 * @lc app=leetcode id=766 lang=golang
 *
 * [766] Toeplitz Matrix
 *
 * https://leetcode.com/problems/toeplitz-matrix/description/
 *
 * algorithms
 * Easy (62.79%)
 * Likes:    732
 * Dislikes: 69
 * Total Accepted:    73.4K
 * Total Submissions: 116.8K
 * Testcase Example:  '[[1,2,3,4],[5,1,2,3],[9,5,1,2]]'
 *
 * A matrix is Toeplitz if every diagonal from top-left to bottom-right has the
 * same element.
 *
 * Now given an M x N matrix, return True if and only if the matrix is
 * Toeplitz.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * matrix = [
 * [1,2,3,4],
 * [5,1,2,3],
 * [9,5,1,2]
 * ]
 * Output: True
 * Explanation:
 * In the above grid, the diagonals are:
 * "[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]".
 * In each diagonal all elements are the same, so the answer is True.
 *
 *
 * Example 2:
 *
 *
 * Input:
 * matrix = [
 * [1,2],
 * [2,2]
 * ]
 * Output: False
 * Explanation:
 * The diagonal "[1, 2]" has different elements.
 *
 *
 *
 * Note:
 *
 *
 * matrix will be a 2D array of integers.
 * matrix will have a number of rows and columns in range [1, 20].
 * matrix[i][j] will be integers in range [0, 99].
 *
 *
 *
 * Follow up:
 *
 *
 * What if the matrix is stored on disk, and the memory is limited such that
 * you can only load at most one row of the matrix into the memory at once?
 * What if the matrix is so large that you can only load up a partial row into
 * the memory at once?
 *
 *
 */

// @lc code=start
//对角元素相等的矩阵
func isToeplitzMatrix(matrix [][]int) bool {
	// 482/482 cases passed (12 ms)
	// Your runtime beats 63.95 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (4.5 MB)
	for i := 0; i < len(matrix[0]); i++ {
		if !check(matrix, matrix[0][i], 0, i) {
			return false
		}
	}

	for i := 0; i < len(matrix); i++ {
		if !check(matrix, matrix[i][0], i, 0) {
			return false
		}
	}
	return true
}

func check(matrix [][]int, value, row, col int) bool {
	if row >= len(matrix) || col >= len(matrix[0]) {
		return true
	}
	if matrix[row][col] != value {
		return false
	}
	return check(matrix, value, row+1, col+1)
}

// @lc code=end

