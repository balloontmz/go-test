/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (45.46%)
 * Likes:    1720
 * Dislikes: 1416
 * Total Accepted:    725.7K
 * Total Submissions: 1.6M
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 *
 * Example 1:
 *
 *
 * Input: 121
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 *
 *
 * Follow up:
 *
 * Coud you solve it without converting the integer to a string?
 *
 */

// @lc code=start
//判断一个整数是否是回文数
// 11509/11509 cases passed (8 ms)
// Your runtime beats 94.64 % of golang submissions
// Your memory usage beats 100 % of golang submissions (5 MB)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var right int
	var left = x
	for x > 0 {
		right = right*10 + x%10
		x /= 10
	}
	return left == right
}

// @lc code=end

