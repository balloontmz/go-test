/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 *
 * https://leetcode.com/problems/longest-palindrome/description/
 *
 * algorithms
 * Easy (48.93%)
 * Likes:    666
 * Dislikes: 67
 * Total Accepted:    115.2K
 * Total Submissions: 235.3K
 * Testcase Example:  '"abccccdd"'
 *
 * Given a string which consists of lowercase or uppercase letters, find the
 * length of the longest palindromes that can be built with those letters.
 *
 * This is case sensitive, for example "Aa" is not considered a palindrome
 * here.
 *
 * Note:
 * Assume the length of given string will not exceed 1,010.
 *
 *
 * Example:
 *
 * Input:
 * "abccccdd"
 *
 * Output:
 * 7
 *
 * Explanation:
 * One longest palindrome that can be built is "dccaccd", whose length is 7.
 *
 *
 */

// @lc code=start
//字符串中能构成回文字符串的最多字符个数
//
// 95/95 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.1 MB)
func longestPalindrome(s string) int {
	var temp = make([]int, 256)
	for i := 0; i < len(s); i++ {
		temp[s[i]-'a']++
	}
	var p = 0
	for _, v := range temp {
		p += (v / 2) * 2
	}
	if p < len(s) {
		p++
	}
	return p
}

// @lc code=end

