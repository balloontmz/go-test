/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 *
 * https://leetcode.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (49.51%)
 * Likes:    1996
 * Dislikes: 146
 * Total Accepted:    346.7K
 * Total Submissions: 697.2K
 * Testcase Example:  '3\n2'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 *
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 *
 * How many possible unique paths are there?
 *
 *
 * Above is a 7 x 3 grid. How many possible unique paths are there?
 *
 * Note: m and n will be at most 100.
 *
 * Example 1:
 *
 *
 * Input: m = 3, n = 2
 * Output: 3
 * Explanation:
 * From the top-left corner, there are a total of 3 ways to reach the
 * bottom-right corner:
 * 1. Right -> Right -> Down
 * 2. Right -> Down -> Right
 * 3. Down -> Right -> Right
 *
 *
 * Example 2:
 *
 *
 * Input: m = 7, n = 3
 * Output: 28
 *
 */

// @lc code=start
//动态规划解法 --
func uniquePaths(m int, n int) int {
	// 62/62 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2 MB)
	var dp = make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[m-1]
}

// 组合解法,此处需要再去看一下排列组合!!! -- 以此推断,组合是否能用动态规划求解!!!
// func uniquePaths(m int, n int) int {
// 	// 62/62 cases passed (0 ms)
// 	// Your runtime beats 100 % of golang submissions
// 	// Your memory usage beats 100 % of golang submissions (1.9 MB)
// 	s := m + n - 2
// 	d := n - 1
// 	// h := 1
// 	// l := 1
// 	r := 1
// 	for i := 1; i <= d; i++ {
// 		// h = h * (s - d + i)
// 		// l = l * i
// 		r = r * (s - d + i) / i // c(s, d+1) d的数学索引从 1 开始,程序索引从 0 开始!!!
// 	}
// 	// return h / l
// 	return r
// }
// @lc code=end