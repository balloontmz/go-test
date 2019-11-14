/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (37.49%)
 * Likes:    3601
 * Dislikes: 175
 * Total Accepted:    752.2K
 * Total Submissions: 2M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */

// @lc code=start
//用栈实现括号匹配
// 76/76 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 90 % of golang submissions (2 MB)
func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			ele := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if s[i] == ')' && ele != '(' {
				return false
			}
			if s[i] == '}' && ele != '{' {
				return false
			}
			if s[i] == ']' && ele != '[' {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

// @lc code=end

