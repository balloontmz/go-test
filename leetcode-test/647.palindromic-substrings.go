/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 *
 * https://leetcode.com/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (58.33%)
 * Likes:    1760
 * Dislikes: 89
 * Total Accepted:    130.9K
 * Total Submissions: 224K
 * Testcase Example:  '"abc"'
 *
 * Given a string, your task is to count how many palindromic substrings in
 * this string.
 *
 * The substrings with different start indexes or end indexes are counted as
 * different substrings even they consist of same characters.
 *
 * Example 1:
 *
 *
 * Input: "abc"
 * Output: 3
 * Explanation: Three palindromic strings: "a", "b", "c".
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "aaa"
 * Output: 6
 * Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
 *
 *
 *
 *
 * Note:
 *
 *
 * The input string length won't exceed 1000.
 *
 *
 *
 */

// @lc code=start
//字符串回文子串的个数
func countSubstrings(s string) int {
	// 130/130 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2 MB)
	var cnt int
	for i := 0; i < len(s); i++ {
		//奇数回文
		var start1, end1 = i, i
		for start1 >= 0 && end1 < len(s) && s[start1] == s[end1] {
			start1--
			end1++
			cnt++
		}
		//偶数回文
		var start2, end2 = i, i + 1
		for start2 >= 0 && end2 < len(s) && s[start2] == s[end2] {
			start2--
			end2++
			cnt++
		}
	}
	return cnt
}

// @lc code=end

