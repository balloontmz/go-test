/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 *
 * https://leetcode.com/problems/reverse-string-ii/description/
 *
 * algorithms
 * Easy (46.46%)
 * Likes:    366
 * Dislikes: 1080
 * Total Accepted:    80.7K
 * Total Submissions: 169.3K
 * Testcase Example:  '"abcdefg"\n2'
 *
 * 
 * Given a string and an integer k, you need to reverse the first k characters
 * for every 2k characters counting from the start of the string. If there are
 * less than k characters left, reverse all of them. If there are less than 2k
 * but greater than or equal to k characters, then reverse the first k
 * characters and left the other as original.
 * 
 * 
 * Example:
 * 
 * Input: s = "abcdefg", k = 2
 * Output: "bacdfeg"
 * 
 * 
 * 
 * Restrictions: 
 * 
 * ⁠The string consists of lower English letters only.
 * ⁠Length of the given string and k will in the range [1, 10000]
 * 
 */

// @lc code=start
// 60/60 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (5.8 MB)
func reverseStr(s string, k int) string {
	var l = len(s)
	var in = l / k
	if in == 0 {
		return reverse(s)
	}
	var o = ""
	for i := 1; i <= in; i++ {
		if i % 2 == 0 {
			o += s[(i-1)*k:i*k]
		} else {
			o += reverse(s[(i-1)*k:i*k])
		}
	}
	if in % 2 == 0 {
		return o + reverse(s[in*k:])
	}
	return o + s[in*k:]
}

func reverse(o string) string {
	var n []byte
	for i := len(o)-1; i >=0; i-- {
		n = append(n, o[i])
	}
	return string(n)
}
// @lc code=end

