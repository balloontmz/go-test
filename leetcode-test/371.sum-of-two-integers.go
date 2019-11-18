/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 *
 * https://leetcode.com/problems/sum-of-two-integers/description/
 *
 * algorithms
 * Easy (50.77%)
 * Likes:    903
 * Dislikes: 1705
 * Total Accepted:    154.7K
 * Total Submissions: 304.7K
 * Testcase Example:  '1\n2'
 *
 * Calculate the sum of two integers a and b, but you are not allowed to use
 * the operator + and -.
 *
 *
 * Example 1:
 *
 *
 * Input: a = 1, b = 2
 * Output: 3
 *
 *
 *
 * Example 2:
 *
 *
 * Input: a = -2, b = 3
 * Output: 1
 *
 *
 *
 *
 */

// @lc code=start
//整数的加法
func getSum(a int, b int) int {
	// 13/13 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (1.9 MB)
	if b == 0 {
		return a
	} else {
		return getSum(a^b, (a&b)<<1)
	}
}

// @lc code=end

