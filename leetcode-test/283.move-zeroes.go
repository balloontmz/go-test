/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 *
 * https://leetcode.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (55.51%)
 * Likes:    2548
 * Dislikes: 90
 * Total Accepted:    560.1K
 * Total Submissions: 1M
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * Given an array nums, write a function to move all 0's to the end of it while
 * maintaining the relative order of the non-zero elements.
 *
 * Example:
 *
 *
 * Input: [0,1,0,3,12]
 * Output: [1,3,12,0,0]
 *
 * Note:
 *
 *
 * You must do this in-place without making a copy of the array.
 * Minimize the total number of operations.
 *
 */

// @lc code=start
//将数组中的0移到数组尾部
func moveZeroes(nums []int) {
	// 21/21 cases passed (64 ms)
	// Your runtime beats 80.45 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (7.7 MB)
	var n = len(nums)
	var idx = 0
	for _, v := range nums {
		if v != 0 {
			nums[idx] = v
			idx++
		}
	}
	for idx < n {
		nums[idx] = 0
		idx++
	}
	return
}

// @lc code=end

