/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 *
 * https://leetcode.com/problems/factorial-trailing-zeroes/description/
 *
 * algorithms
 * Easy (37.64%)
 * Likes:    576
 * Dislikes: 822
 * Total Accepted:    176.2K
 * Total Submissions: 468.3K
 * Testcase Example:  '3'
 *
 * Given an integer n, return the number of trailing zeroes in n!.
 *
 * Example 1:
 *
 *
 * Input: 3
 * Output: 0
 * Explanation: 3! = 6, no trailing zero.
 *
 * Example 2:
 *
 *
 * Input: 5
 * Output: 1
 * Explanation: 5! = 120, one trailing zero.
 *
 * Note: Your solution should be in logarithmic time complexity.
 *
 */

// @lc code=start
//找出阶乘结果中0的个数
func trailingZeroes(n int) int {
	// 502/502 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 66.67 % of golang submissions (2 MB)
	var res = 0
	for n >= 5 {
		res += n / 5
		n = n / 5
	}
	return res
}

// @lc code=end

