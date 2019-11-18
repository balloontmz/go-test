/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 *
 * https://leetcode.com/problems/power-of-two/description/
 *
 * algorithms
 * Easy (42.48%)
 * Likes:    545
 * Dislikes: 145
 * Total Accepted:    257.1K
 * Total Submissions: 604.5K
 * Testcase Example:  '1'
 *
 * Given an integer, write a function to determine if it is a power of two.
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: true
 * Explanation: 2^0 = 1
 *
 *
 * Example 2:
 *
 *
 * Input: 16
 * Output: true
 * Explanation: 2^4 = 16
 *
 * Example 3:
 *
 *
 * Input: 218
 * Output: false
 *
 */

// @lc code=start
//判断一个数是否为2的n次方
func isPowerOfTwo(n int) bool {
	//代表二进制表示只存在一个1
	// 1108/1108 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 75 % of golang submissions (2.2 MB)
	return n > 0 && (n&(n-1) == 0)
}

// @lc code=end

