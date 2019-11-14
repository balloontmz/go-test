/*
 * @lc app=leetcode id=696 lang=golang
 *
 * [696] Count Binary Substrings
 *
 * https://leetcode.com/problems/count-binary-substrings/description/
 *
 * algorithms
 * Easy (54.55%)
 * Likes:    739
 * Dislikes: 133
 * Total Accepted:    36.6K
 * Total Submissions: 67.1K
 * Testcase Example:  '"00110"'
 *
 * Give a string s, count the number of non-empty (contiguous) substrings that
 * have the same number of 0's and 1's, and all the 0's and all the 1's in
 * these substrings are grouped consecutively.
 *
 * Substrings that occur multiple times are counted the number of times they
 * occur.
 *
 * Example 1:
 *
 * Input: "00110011"
 * Output: 6
 * Explanation: There are 6 substrings that have equal number of consecutive
 * 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
 * Notice that some of these substrings repeat and are counted the number of
 * times they occur.
 * Also, "00110011" is not a valid substring because all the 0's (and 1's) are
 * not grouped together.
 *
 *
 *
 * Example 2:
 *
 * Input: "10101"
 * Output: 4
 * Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal
 * number of consecutive 1's and 0's.
 *
 *
 *
 * Note:
 * s.length will be between 1 and 50,000.
 * s will only consist of "0" or "1" characters.
 *
 */

// @lc code=start
//统计二进制字符串中连续 1 和连续 0 数量相同的子字符串个数
func countBinarySubstrings(s string) int {
	// 90/90 cases passed (8 ms)
	// Your runtime beats 94.44 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (5.9 MB)
	//采用两个变量分别记录当前最大连续和前一个最大连续。
	var preLen, curLen, count = 0, 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curLen++
		} else {
			preLen = curLen
			curLen = 1
		}
		if preLen >= curLen {
			count++
		}
	}
	return count
}

// @lc code=end

