/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 *
 * https://leetcode.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (38.54%)
 * Likes:    964
 * Dislikes: 281
 * Total Accepted:    242.3K
 * Total Submissions: 628.1K
 * Testcase Example:  '"egg"\n"add"'
 *
 * Given two strings s and t, determine if they are isomorphic.
 *
 * Two strings are isomorphic if the characters in s can be replaced to get t.
 *
 * All occurrences of a character must be replaced with another character while
 * preserving the order of characters. No two characters may map to the same
 * character but a character may map to itself.
 *
 * Example 1:
 *
 *
 * Input: s = "egg", t = "add"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "foo", t = "bar"
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: s = "paper", t = "title"
 * Output: true
 *
 * Note:
 * You may assume both s and t have the same length.
 *
 */

// @lc code=start
//判断两个字符串是否结构相同
// 30/30 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.8 MB)
func isIsomorphic(s string, t string) bool {
	var temp1 = make([]int, 256)
	var temp2 = make([]int, 256)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if temp1[s[i]] != temp2[t[i]] {
			return false
		}
		//记录对应的上一次出现时间，而不是出现总数。+1是防止0
		temp1[s[i]] = i + 1
		temp2[t[i]] = i + 1
	}
	return true
}

// @lc code=end

