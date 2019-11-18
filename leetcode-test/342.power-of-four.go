/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 *
 * https://leetcode.com/problems/power-of-four/description/
 *
 * algorithms
 * Easy (40.84%)
 * Likes:    366
 * Dislikes: 167
 * Total Accepted:    128.1K
 * Total Submissions: 313.3K
 * Testcase Example:  '16'
 *
 * Given an integer (signed 32 bits), write a function to check whether it is a
 * power of 4.
 *
 * Example 1:
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
 * Input: 5
 * Output: false
 *
 *
 * Follow up: Could you solve it without loops/recursion?
 */

// @lc code=start
//判断一个数是否为4的n次方
func isPowerOfFour(num int) bool {
	//二进制表示只有一个1而且在奇数位上
	// 1060/1060 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2.1 MB)
	return num > 0 && (num&(num-1) == 0) && (num&0x55555555 != 0)
}

// @lc code=end

