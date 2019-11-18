/*
 * @lc app=leetcode id=693 lang=golang
 *
 * [693] Binary Number with Alternating Bits
 *
 * https://leetcode.com/problems/binary-number-with-alternating-bits/description/
 *
 * algorithms
 * Easy (58.43%)
 * Likes:    339
 * Dislikes: 73
 * Total Accepted:    49.3K
 * Total Submissions: 84.3K
 * Testcase Example:  '5'
 *
 * Given a positive integer, check whether it has alternating bits: namely, if
 * two adjacent bits will always have different values.
 *
 * Example 1:
 *
 * Input: 5
 * Output: True
 * Explanation:
 * The binary representation of 5 is: 101
 *
 *
 *
 * Example 2:
 *
 * Input: 7
 * Output: False
 * Explanation:
 * The binary representation of 7 is: 111.
 *
 *
 *
 * Example 3:
 *
 * Input: 11
 * Output: False
 * Explanation:
 * The binary representation of 11 is: 1011.
 *
 *
 *
 * Example 4:
 *
 * Input: 10
 * Output: True
 * Explanation:
 * The binary representation of 10 is: 1010.
 *
 *
 */

// @lc code=start
//判断一个数的二进制表示是否全是不连续的01
func hasAlternatingBits(n int) bool {
	// 204/204 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (1.9 MB)
	var a = n ^ (n >> 1) // 此处用异或而不用或，可能存在同位两个都是1的情况
	// a 的二进制表示全是1，除了头部
	return a&(a+1) == 0 // 去除a+1二进制表示中最低的一位1，由于只有一位1，所以变成0
}

// @lc code=end

