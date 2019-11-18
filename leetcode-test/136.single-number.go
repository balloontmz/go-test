/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 *
 * https://leetcode.com/problems/single-number/description/
 *
 * algorithms
 * Easy (61.90%)
 * Likes:    3049
 * Dislikes: 119
 * Total Accepted:    566.5K
 * Total Submissions: 913.4K
 * Testcase Example:  '[2,2,1]'
 *
 * Given a non-empty array of integers, every element appears twice except for
 * one. Find that single one.
 *
 * Note:
 *
 * Your algorithm should have a linear runtime complexity. Could you implement
 * it without using extra memory?
 *
 * Example 1:
 *
 *
 * Input: [2,2,1]
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: [4,1,2,1,2]
 * Output: 4
 *
 *
 */

// @lc code=start
//找出数组中唯一不重复的数
func singleNumber(nums []int) int {
	// 16/16 cases passed (8 ms)
	// Your runtime beats 96.45 % of golang submissions
	// Your memory usage beats 57.14 % of golang submissions (4.7 MB)
	var ret int
	for _, num := range nums {
		ret ^= num
	}
	return ret
}

// @lc code=end

