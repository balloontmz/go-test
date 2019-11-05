/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (41.14%)
 * Likes:    1223
 * Dislikes: 229
 * Total Accepted:    355.4K
 * Total Submissions: 863.6K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 */

// @lc code=start
//二进制加法
func addBinary(a string, b string) string {
	// 294/294 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 50 % of golang submissions (2.6 MB)
	var i = len(a) - 1
	var j = len(b) - 1
	var temp = 0
	var res = ""
	for temp == 1 || i >= 0 || j >= 0 {
		if i >= 0 && a[i] == '1' {
			temp++
		}
		if j >= 0 && b[j] == '1' {
			temp++
		}
		i--
		j--
		res = string((temp%2)+'0') + res
		temp = temp / 2
	}
	return res
}

// @lc code=end

