/*
 * @lc app=leetcode id=1422 lang=golang
 *
 * [1422] Maximum Score After Splitting a String
 *
 * https://leetcode.com/problems/maximum-score-after-splitting-a-string/description/
 *
 * algorithms
 * Easy (53.97%)
 * Likes:    182
 * Dislikes: 10
 * Total Accepted:    18.1K
 * Total Submissions: 33.1K
 * Testcase Example:  '"011101"'
 *
 * Given a string s of zeros and ones, return the maximum score after splitting
 * the string into two non-empty substrings (i.e. left substring and right
 * substring).
 * 
 * The score after splitting a string is the number of zeros in the left
 * substring plus the number of ones in the right substring.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s = "011101"
 * Output: 5 
 * Explanation: 
 * All possible ways of splitting s into two non-empty substrings are:
 * left = "0" and right = "11101", score = 1 + 4 = 5 
 * left = "01" and right = "1101", score = 1 + 3 = 4 
 * left = "011" and right = "101", score = 1 + 2 = 3 
 * left = "0111" and right = "01", score = 1 + 1 = 2 
 * left = "01110" and right = "1", score = 2 + 1 = 3
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "00111"
 * Output: 5
 * Explanation: When left = "00" and right = "111", we get the maximum score =
 * 2 + 3 = 5
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: s = "1111"
 * Output: 3
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 2 <= s.length <= 500
 * The string s consists of characters '0' and '1' only.
 * 
 * 
 */

// @lc code=start
//计算最佳切分后字符串个数统计
// 103/103 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2 MB)
func maxScore(s string) int {
	var zCnt int
	for _, b := range s {
		if b == '0' {
			zCnt++
		}
	}
	var cZ int
	if s[0] == '0' {
		cZ = 1
	}
	var max = len(s) - 1 - zCnt + 2 * cZ
	for k, v := range s[1:len(s)-1] {
		if v == '0' {
			cZ++
			if len(s) - k - 2 - zCnt + 2 * cZ > max {
				max = len(s) - k - 2 - zCnt + 2 * cZ
			}
		}
	}
	return max
}
// @lc code=end

