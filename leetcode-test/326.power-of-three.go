/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 *
 * https://leetcode.com/problems/power-of-three/description/
 *
 * algorithms
 * Easy (41.88%)
 * Likes:    349
 * Dislikes: 1212
 * Total Accepted:    211.4K
 * Total Submissions: 504.8K
 * Testcase Example:  '27'
 *
 * Given an integer, write a function to determine if it is a power of three.
 *
 * Example 1:
 *
 *
 * Input: 27
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: 0
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: 9
 * Output: true
 *
 * Example 4:
 *
 *
 * Input: 45
 * Output: false
 *
 * Follow up:
 * Could you do it without using any loop / recursion?
 */

// @lc code=start
//求一个数是否为3的n次方
func isPowerOfThree(n int) bool {
	// 21038/21038 cases passed (36 ms)
	// Your runtime beats 23.7 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.1 MB)
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%3 != 0 {
			return false
		}
		n = n / 3
	}
	return true

}

// @lc code=end

