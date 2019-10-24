/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 *
 * https://leetcode.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (48.95%)
 * Likes:    1710
 * Dislikes: 44
 * Total Accepted:    271.5K
 * Total Submissions: 552.5K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * Given a m x n grid filled with non-negative numbers, find a path from top
 * left to bottom right which minimizes the sum of all numbers along its path.
 *
 * Note: You can only move either down or right at any point in time.
 *
 * Example:
 *
 *
 * Input:
 * [
 * [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * Output: 7
 * Explanation: Because the path 1→3→1→1→1 minimizes the sum.
 *
 *
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	// 61/61 cases passed (8 ms)
	// Your runtime beats 88.63 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (3.9 MB)
	var m = len(grid[0])
	var n = len(grid)
	var p = make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 0 {
				p[j] = p[j]
			} else if i == 0 {
				p[j] = p[j-1]
			} else {
				p[j] = min(p[j], p[j-1])
			}
			p[j] += grid[i][j]
		}
	}
	return p[m-1]

}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

