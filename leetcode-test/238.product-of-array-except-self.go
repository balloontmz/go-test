/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 *
 * https://leetcode.com/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (57.30%)
 * Likes:    3005
 * Dislikes: 259
 * Total Accepted:    332.2K
 * Total Submissions: 579.7K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an array nums of n integers where n > 1,  return an array output such
 * that output[i] is equal to the product of all the elements of nums except
 * nums[i].
 *
 * Example:
 *
 *
 * Input:  [1,2,3,4]
 * Output: [24,12,8,6]
 *
 *
 * Note: Please solve it without division and in O(n).
 *
 * Follow up:
 * Could you solve it with constant space complexity? (The output array does
 * not count as extra space for the purpose of space complexity analysis.)
 *
 */

// @lc code=start
//o(n)复杂度计算除了自己的乘积数组
func productExceptSelf(nums []int) []int {
	// 17/17 cases passed (1440 ms)
	// Your runtime beats 99.77 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (9.1 MB)
	var n = len(nums)
	var res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	var left = 1
	for i := 1; i < n; i++ {
		left *= nums[i-1]
		res[i] *= left
	}
	var right = 1
	for i := n - 2; i >= 0; i-- {
		right *= nums[i+1]
		res[i] *= right
	}
	return res
}

// @lc code=end

