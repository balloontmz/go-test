/*
 * @lc app=leetcode id=482 lang=golang
 *
 * [482] License Key Formatting
 *
 * https://leetcode.com/problems/license-key-formatting/description/
 *
 * algorithms
 * Easy (41.99%)
 * Likes:    397
 * Dislikes: 652
 * Total Accepted:    113.8K
 * Total Submissions: 267.5K
 * Testcase Example:  '"5F3Z-2e-9-w"\n4'
 *
 * You are given a license key represented as a string S which consists only
 * alphanumeric character and dashes. The string is separated into N+1 groups
 * by N dashes.
 * 
 * Given a number K, we would want to reformat the strings such that each group
 * contains exactly K characters, except for the first group which could be
 * shorter than K, but still must contain at least one character. Furthermore,
 * there must be a dash inserted between two groups and all lowercase letters
 * should be converted to uppercase.
 * 
 * Given a non-empty string S and a number K, format the string according to
 * the rules described above.
 * 
 * Example 1:
 * 
 * Input: S = "5F3Z-2e-9-w", K = 4
 * 
 * Output: "5F3Z-2E9W"
 * 
 * Explanation: The string S has been split into two parts, each part has 4
 * characters.
 * Note that the two extra dashes are not needed and can be removed.
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: S = "2-5g-3-J", K = 2
 * 
 * Output: "2-5G-3J"
 * 
 * Explanation: The string S has been split into three parts, each part has 2
 * characters except the first part as it could be shorter as mentioned
 * above.
 * 
 * 
 * 
 * Note:
 * 
 * The length of string S will not exceed 12,000, and K is a positive integer.
 * String S consists only of alphanumerical characters (a-z and/or A-Z and/or
 * 0-9) and dashes(-).
 * String S is non-empty.
 * 
 * 
 */

// @lc code=start
// 38/38 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (5 MB)
func licenseKeyFormatting(S string, K int) string {
	var res []byte
	var l = len(S)
	var tag int
	for i := l-1; i >= 0; i-- {
		if S[i] == '-' {
			continue
		} else {
			if S[i] >= 97 {
				res = append(res, S[i]-32)	
			} else {
				res = append(res, S[i])	
			}
			tag++
			if tag == K {
				tag = 0
				res = append(res, '-')
			}
		}
	}
	if len(res) > 0 && res[len(res)-1] == '-' {
		res = res[:len(res)-1]	
	}
	
	for from, to := 0, len(res)-1; from < to; from, to = from+1, to-1 {
		res[from], res[to] = res[to], res[from]
	}
	return string(res)
}
// @lc code=end

