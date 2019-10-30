/*
 * @lc app=leetcode id=474 lang=golang
 *
 * [474] Ones and Zeroes
 *
 * https://leetcode.com/problems/ones-and-zeroes/description/
 *
 * algorithms
 * Medium (40.76%)
 * Likes:    667
 * Dislikes: 160
 * Total Accepted:    36.3K
 * Total Submissions: 88.9K
 * Testcase Example:  '["10","0001","111001","1","0"]\n5\n3'
 *
 * In the computer world, use restricted resource you have to generate maximum
 * benefit is what we always want to pursue.
 *
 * For now, suppose you are a dominator of m 0s and n 1s respectively. On the
 * other hand, there is an array with strings consisting of only 0s and 1s.
 *
 * Now your task is to find the maximum number of strings that you can form
 * with given m 0s and n 1s. Each 0 and 1 can be used at most once.
 *
 * Note:
 *
 *
 * The given numbers of 0s and 1s will both not exceed 100
 * The size of given string array won't exceed 600.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
 * Output: 4
 *
 * Explanation: This are totally 4 strings can be formed by the using of 5 0s
 * and 3 1s, which are “10,”0001”,”1”,”0”
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: Array = {"10", "0", "1"}, m = 1, n = 1
 * Output: 2
 *
 * Explanation: You could form "10", but then you'd have nothing left. Better
 * form "0" and "1".
 *
 *
 *
 *
 */

// @lc code=start
//动态规划解01 字符构成最多的字符串
func findMaxForm(strs []string, m int, n int) int {
	// 63/63 cases passed (36 ms)
	// Your runtime beats 12.07 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (3.4 MB)
	var dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		zeros := 0
		ones := 0
		for i := 0; i < len(str); i++ {
			if str[i] == '0' {
				zeros++
			} else {
				ones++
			}
		}
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}
	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

