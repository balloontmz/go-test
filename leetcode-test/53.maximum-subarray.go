/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (43.68%)
 * Likes:    4745
 * Dislikes: 176
 * Total Accepted:    586.6K
 * Total Submissions: 1.3M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 * 
 * Example:
 * 
 * 
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 * 
 * 
 * Follow up:
 * 
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 * 
 */
func maxSubArray(nums []int) int {
	// ✔ 202/202 cases passed (4 ms)
	// ✔ Your runtime beats 97.18 % of golang submissions
	// ✔ Your memory usage beats 59.26 % of golang submissions (3.3 MB)
	// 计算子序列的最大和
    if nums == nil || len(nums) == 0 {
		return 0
	}
	var preSum = nums[0]
	var maxSum = nums[0]
	for i := 1; i < len(nums); i++ {
		if preSum > 0  {
			preSum = preSum + nums[i]
		} else {
			preSum = nums[i]
		}
		maxSum = max(maxSum, preSum)
	}
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
