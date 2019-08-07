/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 *
 * https://leetcode.com/problems/is-subsequence/description/
 *
 * algorithms
 * Medium (46.82%)
 * Likes:    601
 * Dislikes: 135
 * Total Accepted:    95.1K
 * Total Submissions: 201.2K
 * Testcase Example:  '"abc"\n"ahbgdc"'
 *
 * 
 * Given a string s and a string t, check if s is subsequence of t.
 * 
 * 
 * 
 * You may assume that there is only lower case English letters in both s and
 * t. t is potentially a very long (length ~= 500,000) string, and s is a short
 * string (
 * 
 * 
 * A subsequence of a string is a new string which is formed from the original
 * string by deleting some (can be none) of the characters without disturbing
 * the relative positions of the remaining characters. (ie, "ace" is a
 * subsequence of "abcde" while "aec" is not).
 * 
 * 
 * Example 1:
 * s = "abc", t = "ahbgdc"
 * 
 * 
 * Return true.
 * 
 * 
 * Example 2:
 * s = "axc", t = "ahbgdc"
 * 
 * 
 * Return false.
 * 
 * 
 * Follow up:
 * If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B, and you
 * want to check one by one to see if T has its subsequence. In this scenario,
 * how would you change your code?
 * 
 * Credits:Special thanks to @pbrother for adding this problem and creating all
 * test cases.
 */
func isSubsequence(s string, t string) bool {
	// ✔ 14/14 cases passed (16 ms)
	// ✔ Your runtime beats 32.08 % of golang submissions
	// ✔ Your memory usage beats 83.33 % of golang submissions (6.4 MB)
	var index int
Label:
    for _, v := range s {
		// 
		for index < len(t) {
			if t[index] == byte(v) {
				index++
				continue Label
			} else {
				index++
				continue
			}
		}
		if index >= len(t) {
			return false
		}
	}
	return true
}

