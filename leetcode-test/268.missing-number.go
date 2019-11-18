/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 *
 * https://leetcode.com/problems/missing-number/description/
 *
 * algorithms
 * Easy (49.52%)
 * Likes:    1161
 * Dislikes: 1585
 * Total Accepted:    341.3K
 * Total Submissions: 688.2K
 * Testcase Example:  '[3,0,1]'
 *
 * Given an array containing n distinct numbers taken from 0, 1, 2, ..., n,
 * find the one that is missing from the array.
 *
 * Example 1:
 *
 *
 * Input: [3,0,1]
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [9,6,4,2,3,5,7,0,1]
 * Output: 8
 *
 *
 * Note:
 * Your algorithm should run in linear runtime complexity. Could you implement
 * it using only constant extra space complexity?
 */

// @lc code=start
//找出数组中缺失的数
func missingNumber(nums []int) int {
	// 122/122 cases passed (20 ms)
	// Your runtime beats 28.97 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (5.9 MB)
	var ret int
	for k, v := range nums {
		ret = ret ^ (k + 1) ^ v
	}
	return ret
}

// @lc code=end

