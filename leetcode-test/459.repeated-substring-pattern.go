/*
 * @lc app=leetcode id=459 lang=golang
 *
 * [459] Repeated Substring Pattern
 *
 * https://leetcode.com/problems/repeated-substring-pattern/description/
 *
 * algorithms
 * Easy (40.67%)
 * Likes:    1088
 * Dislikes: 120
 * Total Accepted:    103.1K
 * Total Submissions: 250K
 * Testcase Example:  '"abab"'
 *
 * Given a non-empty string check if it can be constructed by taking a
 * substring of it and appending multiple copies of the substring together. You
 * may assume the given string consists of lowercase English letters only and
 * its length will not exceed 10000.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "abab"
 * Output: True
 * Explanation: It's the substring "ab" twice.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "aba"
 * Output: False
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "abcabcabcabc"
 * Output: True
 * Explanation: It's the substring "abc" four times. (And the substring
 * "abcabc" twice.)
 * 
 * 
 */

// @lc code=start
//重复子字符串模式
// 120/120 cases passed (352 ms)
// Your runtime beats 11.92 % of golang submissions
// Your memory usage beats 50 % of golang submissions (8 MB)
func repeatedSubstringPattern(s string) bool {
	var n = len(s)
	for i := n / 2; i > 0; i-- {
		if n % i == 0 { // 只有 n 能整除 i 的情况下才可能组成长度相等的字符串
			c := n / i // 重复次数
			str := ""
			for j := 0; j < c; j++ {
				str += s[:i]
			}
			if str == s {
				return true
			}
		}
	}
	return false
}
// @lc code=end

