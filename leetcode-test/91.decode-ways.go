import "strconv"

/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 *
 * https://leetcode.com/problems/decode-ways/description/
 *
 * algorithms
 * Medium (23.00%)
 * Likes:    1791
 * Dislikes: 2032
 * Total Accepted:    307.6K
 * Total Submissions: 1.3M
 * Testcase Example:  '"12"'
 *
 * A message containing letters from A-Z is being encoded to numbers using the
 * following mapping:
 *
 *
 * 'A' -> 1
 * 'B' -> 2
 * ...
 * 'Z' -> 26
 *
 *
 * Given a non-empty string containing only digits, determine the total number
 * of ways to decode it.
 *
 * Example 1:
 *
 *
 * Input: "12"
 * Output: 2
 * Explanation: It could be decoded as "AB" (1 2) or "L" (12).
 *
 *
 * Example 2:
 *
 *
 * Input: "226"
 * Output: 3
 * Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2
 * 6).
 *
 */

// @lc code=start
func numDecodings(s string) int {
	// 258/258 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2.1 MB)
	if len(s) <= 0 {
		return 0
	}
	n := len(s)
	var dp = make([]int, n+1)
	dp[0] = 1
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}
	for i := 2; i <= n; i++ {
		one, _ := strconv.Atoi(s[i-1 : i])
		if one != 0 {
			dp[i] += dp[i-1]
		}
		if s[i-2] == '0' {
			continue
		}
		two, _ := strconv.Atoi(s[i-2 : i])
		if two <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

// @lc code=end

