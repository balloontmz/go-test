/*
 * @lc app=leetcode id=665 lang=golang
 *
 * [665] Non-decreasing Array
 *
 * https://leetcode.com/problems/non-decreasing-array/description/
 *
 * algorithms
 * Easy (19.48%)
 * Likes:    1143
 * Dislikes: 249
 * Total Accepted:    58.7K
 * Total Submissions: 301.5K
 * Testcase Example:  '[4,2,3]'
 *
 * 
 * Given an array with n integers, your task is to check if it could become
 * non-decreasing by modifying at most 1 element.
 * 
 * 
 * 
 * We define an array is non-decreasing if array[i]  holds for every i (1 
 * 
 * Example 1:
 * 
 * Input: [4,2,3]
 * Output: True
 * Explanation: You could modify the first 4 to 1 to get a non-decreasing
 * array.
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: [4,2,1]
 * Output: False
 * Explanation: You can't get a non-decreasing array by modify at most one
 * element.
 * 
 * 
 * 
 * Note:
 * The n belongs to [1, 10,000].
 * 
 */
func checkPossibility(nums []int) bool {
	// ✔ 325/325 cases passed (28 ms)
	// ✔ Your runtime beats 25.56 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (6.2 MB)
	var cnt = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[i-1] {
			continue
		}
		cnt ++
		if (i - 2 >=0) && (nums[i] < nums[i-2]) {
			nums[i] = nums[i-1]
		} else {
			nums[i-1] = nums[i]
		}
	}
	return cnt <= 1
}

