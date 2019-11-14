/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 *
 * https://leetcode.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (54.06%)
 * Likes:    929
 * Dislikes: 121
 * Total Accepted:    424.8K
 * Total Submissions: 784.4K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * Given two strings s and t , write a function to determine if t is an anagram
 * of s.
 *
 * Example 1:
 *
 *
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "rat", t = "car"
 * Output: false
 *
 *
 * Note:
 * You may assume the string contains only lowercase alphabets.
 *
 * Follow up:
 * What if the inputs contain unicode characters? How would you adapt your
 * solution to such case?
 *
 */

// @lc code=start
//判断两个字符串的组成是否相同
// 34/34 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (3 MB)
func isAnagram(s string, t string) bool {
	var temp = make([]int, 26)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		temp[int(s[i]-'a')]++
	}
	for i := 0; i < len(t); i++ {
		temp[int(t[i]-'a')]--
	}
	for i := 0; i < 26; i++ {
		if temp[i] != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

