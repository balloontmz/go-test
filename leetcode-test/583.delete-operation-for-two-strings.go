/*
 * @lc app=leetcode id=583 lang=golang
 *
 * [583] Delete Operation for Two Strings
 *
 * https://leetcode.com/problems/delete-operation-for-two-strings/description/
 *
 * algorithms
 * Medium (45.83%)
 * Likes:    800
 * Dislikes: 21
 * Total Accepted:    37.5K
 * Total Submissions: 81.4K
 * Testcase Example:  '"sea"\n"eat"'
 *
 *
 * Given two words word1 and word2, find the minimum number of steps required
 * to make word1 and word2 the same, where in each step you can delete one
 * character in either string.
 *
 *
 * Example 1:
 *
 * Input: "sea", "eat"
 * Output: 2
 * Explanation: You need one step to make "sea" to "ea" and another step to
 * make "eat" to "ea".
 *
 *
 *
 * Note:
 *
 * The length of given words won't exceed 500.
 * Characters in given words can only be lower-case letters.
 *
 *
 */

// @lc code=start
//动态规划解求两个字符串的最长子序列
func minDistance(word1 string, word2 string) int {
	// 1307/1307 cases passed (4 ms)
	// Your runtime beats 75 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.1 MB)
	m := len(word1)
	n := len(word2)
	var dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return m + n - 2*dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

