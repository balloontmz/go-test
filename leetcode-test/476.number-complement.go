/*
 * @lc app=leetcode id=476 lang=golang
 *
 * [476] Number Complement
 *
 * https://leetcode.com/problems/number-complement/description/
 *
 * algorithms
 * Easy (62.83%)
 * Likes:    608
 * Dislikes: 75
 * Total Accepted:    118.1K
 * Total Submissions: 187.8K
 * Testcase Example:  '5'
 *
 * Given a positive integer, output its complement number. The complement
 * strategy is to flip the bits of its binary representation.
 *
 * Note:
 *
 * The given integer is guaranteed to fit within the range of a 32-bit signed
 * integer.
 * You could assume no leading zero bit in the integer’s binary
 * representation.
 *
 *
 *
 * Example 1:
 *
 * Input: 5
 * Output: 2
 * Explanation: The binary representation of 5 is 101 (no leading zero bits),
 * and its complement is 010. So you need to output 2.
 *
 *
 *
 * Example 2:
 *
 * Input: 1
 * Output: 0
 * Explanation: The binary representation of 1 is 1 (no leading zero bits), and
 * its complement is 0. So you need to output 0.
 *
 *
 */

// @lc code=start
//求一个数的补码
func findComplement(num int) int {

	if num == 0 {
		return 1
	}
	var temp int
	for i := num; i > 0; i = i >> 1 {
		temp = temp << 1
		temp += 1
	}
	return num ^ temp
}

// @lc code=end

