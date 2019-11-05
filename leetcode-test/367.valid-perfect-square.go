/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 *
 * https://leetcode.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (40.51%)
 * Likes:    542
 * Dislikes: 123
 * Total Accepted:    131K
 * Total Submissions: 323.2K
 * Testcase Example:  '16'
 *
 * Given a positive integer num, write a function which returns True if num is
 * a perfect square else False.
 *
 * Note: Do not use any built-in library function such as sqrt.
 *
 * Example 1:
 *
 *
 *
 * Input: 16
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: 14
 * Output: false
 *
 *
 *
 */

// @lc code=start
//判断一个数是否为完全平方数
func isPerfectSquare(num int) bool {
	// 68/68 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (1.9 MB)
	var subSum = 1
	//平方数的差为等差数列
	for num > 0 {
		num = num - subSum
		subSum += 2
	}
	if num == 0 {
		return true
	}
	return false
}

// @lc code=end

