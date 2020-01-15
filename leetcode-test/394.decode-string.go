/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 *
 * https://leetcode.com/problems/decode-string/description/
 *
 * algorithms
 * Medium (47.02%)
 * Likes:    2308
 * Dislikes: 121
 * Total Accepted:    157.7K
 * Total Submissions: 330.9K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * Given an encoded string, return its decoded string.
 * 
 * The encoding rule is: k[encoded_string], where the encoded_string inside the
 * square brackets is being repeated exactly k times. Note that k is guaranteed
 * to be a positive integer.
 * 
 * You may assume that the input string is always valid; No extra white spaces,
 * square brackets are well-formed, etc.
 * 
 * Furthermore, you may assume that the original data does not contain any
 * digits and that digits are only for those repeat numbers, k. For example,
 * there won't be input like 3a or 2[4].
 * 
 * Examples:
 * 
 * 
 * s = "3[a]2[bc]", return "aaabcbc".
 * s = "3[a2[c]]", return "accaccacc".
 * s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
 * 
 * 
 * 
 * 
 */

// @lc code=start
//解码字符串 -- 采用双 stack 缓存前缀和前缀乘数
// 29/29 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.2 MB)
func decodeString(s string) string {
	var stack []int // 用于缓存乘数
	var prefixStack []string
	// var res string
	var temp int
	var tempPrefix string

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			//此处加入前缀和乘数
			prefixStack = append(prefixStack, tempPrefix)
			stack = append(stack, temp)
			tempPrefix = ""
			temp = 0
		} else if s[i] == ']' {
			var t = ""
			for i := 0; i < stack[len(stack)-1]; i++ {
				t += tempPrefix
			}
			stack = stack[:len(stack)-1]
			tempPrefix = prefixStack[len(prefixStack)-1] + t
			prefixStack = prefixStack[:len(prefixStack)-1]
		} else if s[i] >= 65 { // 是字符的情况
			tempPrefix += string(s[i])
		} else {
			temp = temp * 10 + int(s[i]-48)
		}
	}
	return tempPrefix
}
// @lc code=end

