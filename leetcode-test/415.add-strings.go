/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 *
 * https://leetcode.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (45.08%)
 * Likes:    541
 * Dislikes: 182
 * Total Accepted:    122.4K
 * Total Submissions: 271.4K
 * Testcase Example:  '"0"\n"0"'
 *
 * Given two non-negative integers num1 and num2 represented as string, return
 * the sum of num1 and num2.
 *
 * Note:
 *
 * The length of both num1 and num2 is < 5100.
 * Both num1 and num2 contains only digits 0-9.
 * Both num1 and num2 does not contain any leading zero.
 * You must not use any built-in BigInteger library or convert the inputs to
 * integer directly.
 *
 *
 */

// @lc code=start
//字符串加法
func addStrings(num1 string, num2 string) string {
	// 315/315 cases passed (4 ms)
	// Your runtime beats 54.46 % of golang submissions
	// Your memory usage beats 40 % of golang submissions (7.1 MB)
	var i = len(num1) - 1
	var j = len(num2) - 1
	var temp = 0
	var res = ""
	for temp != 0 || i >= 0 || j >= 0 {
		if i >= 0 && num1[i] != '0' {
			temp += int(num1[i] - '0')
		}
		if j >= 0 && num2[j] != '0' {
			temp += int(num2[j] - '0')
		}
		i--
		j--
		res = string(temp%10+'0') + res
		temp = temp / 10
	}
	return res
}

// @lc code=end

