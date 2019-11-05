import "strconv"

/*
 * @lc app=leetcode id=504 lang=golang
 *
 * [504] Base 7
 *
 * https://leetcode.com/problems/base-7/description/
 *
 * algorithms
 * Easy (45.36%)
 * Likes:    175
 * Dislikes: 131
 * Total Accepted:    46.8K
 * Total Submissions: 103.2K
 * Testcase Example:  '100'
 *
 * Given an integer, return its base 7 string representation.
 *
 * Example 1:
 *
 * Input: 100
 * Output: "202"
 *
 *
 *
 * Example 2:
 *
 * Input: -7
 * Output: "-10"
 *
 *
 *
 * Note:
 * The input will be in range of [-1e7, 1e7].
 *
 */

// @lc code=start
//7进制转换
func convertToBase7(num int) string {
	// 241/241 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2 MB)
	if num == 0 {
		return "0"
	}
	var isNavigate = false
	var ret = ""
	if num < 0 {
		isNavigate = true
		num = -num
	}
	for num > 0 {
		ret = strconv.Itoa(num%7) + ret
		num = num / 7
	}
	if isNavigate {
		ret = "-" + ret
	}
	return ret
}

// @lc code=end

