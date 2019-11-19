/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (57.88%)
 * Likes:    3588
 * Dislikes: 208
 * Total Accepted:    423K
 * Total Submissions: 727.5K
 * Testcase Example:  '3'
 *
 *
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 *
 *
 *
 * For example, given n = 3, a solution set is:
 *
 *
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 *
 */

// @lc code=start
//生成括号
func generateParenthesis(n int) []string {
	// 8/8 cases passed (8 ms)
	// Your runtime beats 42.94 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (7.6 MB)
	var res = &([]string{})
	dfs(res, n, 0, "")
	return *res
}

func dfs(res *[]string, start, end int, s string) {
	if start == 0 {
		for i := 0; i < end; i++ {
			s += ")"
		}
		*res = append(*res, s)
		return
	}
	//添加start
	dfs(res, start-1, end+1, s+"(")
	if end != 0 {
		//如果end不为0，则从end栈弹出一个
		dfs(res, start, end-1, s+")")
	}
	return
}

// @lc code=end

