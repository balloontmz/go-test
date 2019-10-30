/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 *
 * https://leetcode.com/problems/word-break/description/
 *
 * algorithms
 * Medium (36.83%)
 * Likes:    2848
 * Dislikes: 155
 * Total Accepted:    405.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * Given a non-empty string s and a dictionary wordDict containing a list of
 * non-empty words, determine if s can be segmented into a space-separated
 * sequence of one or more dictionary words.
 *
 * Note:
 *
 *
 * The same word in the dictionary may be reused multiple times in the
 * segmentation.
 * You may assume the dictionary does not contain duplicate words.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "leetcode", wordDict = ["leet", "code"]
 * Output: true
 * Explanation: Return true because "leetcode" can be segmented as "leet
 * code".
 *
 *
 * Example 2:
 *
 *
 * Input: s = "applepenapple", wordDict = ["apple", "pen"]
 * Output: true
 * Explanation: Return true because "applepenapple" can be segmented as "apple
 * pen apple".
 * Note that you are allowed to reuse a dictionary word.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * Output: false
 *
 *
 */

// @lc code=start
//动态规划解字符串按单词列表分割 -- 暂时未完全理解,下次看.
func wordBreak(s string, wordDict []string) bool {
	// 36/36 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2.1 MB)
	sL := len(s)
	var dp = make([]bool, sL+1)
	dp[0] = true
	for i := 0; i <= sL; i++ {
		for _, word := range wordDict {
			wL := len(word)
			if i-wL >= 0 && s[i-wL:i] == word {
				dp[i] = dp[i] || dp[i-wL]
			}
		}
	}
	return dp[sL]
}

// @lc code=end

