/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 *
 * https://leetcode.com/problems/wildcard-matching/description/
 *
 * algorithms
 * Hard (23.59%)
 * Likes:    1420
 * Dislikes: 82
 * Total Accepted:    205.7K
 * Total Submissions: 869.7K
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement wildcard pattern
 * matching with support for '?' and '*'.
 *
 *
 * '?' Matches any single character.
 * '*' Matches any sequence of characters (including the empty sequence).
 *
 *
 * The matching should cover the entire input string (not partial).
 *
 * Note:
 *
 *
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters
 * like ? or *.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * s = "aa"
 * p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 *
 *
 * Example 2:
 *
 *
 * Input:
 * s = "aa"
 * p = "*"
 * Output: true
 * Explanation: '*' matches any sequence.
 *
 *
 * Example 3:
 *
 *
 * Input:
 * s = "cb"
 * p = "?a"
 * Output: false
 * Explanation: '?' matches 'c', but the second letter is 'a', which does not
 * match 'b'.
 *
 *
 * Example 4:
 *
 *
 * Input:
 * s = "adceb"
 * p = "*a*b"
 * Output: true
 * Explanation: The first '*' matches the empty sequence, while the second '*'
 * matches the substring "dce".
 *
 *
 * Example 5:
 *
 *
 * Input:
 * s = "acdcb"
 * p = "a*c?b"
 * Output: false
 *
 *
 */

// @lc code=start
func isMatch(s string, p string) bool {
	//[外卡匹配](https://www.cnblogs.com/grandyang/p/4401196.html)
	//动态规划外卡匹配
	// 1809/1809 cases passed (8 ms)
	// Your runtime beats 81.18 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.1 MB)
	var m, n = len(s), len(p)
	var dp = make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i < n+1; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else {
				dp[i][j] = (s[i-1] == p[j-1] || p[j-1] == '?') && dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

// func isMatch(s string, p string) bool {
// 	// 1809/1809 cases passed (8 ms)
// 	// Your runtime beats 80.95 % of golang submissions
// 	// Your memory usage beats 100 % of golang submissions (3 MB)
// 	var i, j = 0, 0           // 分别记录s和p的位置
// 	var iStar, jStar = -1, -1 // 分别记录s 和 j 的*位置
// 	var m, n = len(s), len(p)
// 	for i < m {
// 		if j < n && (s[i] == p[j] || p[j] == '?') { // 如果当前遍历的字符能够匹配,则分别自增 1
// 			i++
// 			j++
// 		} else if j < n && p[j] == '*' {
// 			iStar = i
// 			jStar = j
// 			j++
// 		} else if iStar >= 0 { // 此处可用于回退不正确的分支到星号所在位置并前移 s 位置重新开始!!!
// 			iStar++
// 			i = iStar
// 			j = jStar + 1
// 		} else {
// 			return false
// 		}
// 	}
// 	for j < n && p[j] == '*' {
// 		j++
// 	}
// 	return j == n
// }

// @lc code=end

