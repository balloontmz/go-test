/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 *
 * https://leetcode.com/problems/search-a-2d-matrix-ii/description/
 *
 * algorithms
 * Medium (41.76%)
 * Likes:    2057
 * Dislikes: 60
 * Total Accepted:    229.2K
 * Total Submissions: 547.8K
 * Testcase Example:  '[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n' +
  '5'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 *
 *
 * Integers in each row are sorted in ascending from left to right.
 * Integers in each column are sorted in ascending from top to bottom.
 *
 *
 * Example:
 *
 * Consider the following matrix:
 *
 *
 * [
 * ⁠ [1,   4,  7, 11, 15],
 * ⁠ [2,   5,  8, 12, 19],
 * ⁠ [3,   6,  9, 16, 22],
 * ⁠ [10, 13, 14, 17, 24],
 * ⁠ [18, 21, 23, 26, 30]
 * ]
 *
 *
 * Given target = 5, return true.
 *
 * Given target = 20, return false.
 *
*/

// @lc code=start
//对角线递增 -- 无法采用二分法
func searchMatrix(matrix [][]int, target int) bool {
	// 129/129 cases passed (24 ms)
	// Your runtime beats 90.06 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.2 MB)
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	var m, n = len(matrix), len(matrix[0])
	var row, col = 0, n - 1
	for row < m && col >= 0 {
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			row++
		} else {
			col--
		}
	}
	return false
}

// @lc code=end

