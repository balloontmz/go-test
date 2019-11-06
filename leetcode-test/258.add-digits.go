/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 *
 * https://leetcode.com/problems/add-digits/description/
 *
 * algorithms
 * Easy (55.02%)
 * Likes:    536
 * Dislikes: 875
 * Total Accepted:    257.3K
 * Total Submissions: 467.7K
 * Testcase Example:  '38'
 *
 * Given a non-negative integer num, repeatedly add all its digits until the
 * result has only one digit.
 *
 * Example:
 *
 *
 * Input: 38
 * Output: 2
 * Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2.
 * Since 2 has only one digit, return it.
 *
 *
 * Follow up:
 * Could you do it without any loop/recursion in O(1) runtime?
 */

// @lc code=start
//将一个数的所有位数相加，只到最后变成个位数
func addDigits(num int) int {
	// 1101/1101 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 50 % of golang submissions (2.2 MB)
	// 将一个数的所有位数相加，只到最后变成个位数 -- -1 的目的是 类似于 90 的情况计算是9而不是0
	return (num-1)%9 + 1
}

// @lc code=end

