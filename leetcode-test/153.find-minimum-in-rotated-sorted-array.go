/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 *
 * https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (43.36%)
 * Likes:    1122
 * Dislikes: 170
 * Total Accepted:    307.4K
 * Total Submissions: 708.5K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 * 
 * (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
 * 
 * Find the minimum element.
 * 
 * You may assume no duplicate exists in the array.
 * 
 * Example 1:
 * 
 * 
 * Input: [3,4,5,1,2] 
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [4,5,6,7,0,1,2]
 * Output: 0
 * 
 * 
 */
func findMin(nums []int) int {
	// ✔ 146/146 cases passed (0 ms)
	// ✔ Your runtime beats 100 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (2.5 MB)
	var l = 0
	var r = len(nums) - 1
	for l < r {
		var m = l + (r-l)/2
		if nums[m] <= nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}
	return nums[l]
}

