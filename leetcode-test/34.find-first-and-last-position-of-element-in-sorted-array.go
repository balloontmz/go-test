/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (33.97%)
 * Likes:    1833
 * Dislikes: 92
 * Total Accepted:    331.4K
 * Total Submissions: 974.9K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * Given an array of integers nums sorted in ascending order, find the starting
 * and ending position of a given target value.
 * 
 * Your algorithm's runtime complexity must be in the order of O(log n).
 * 
 * If the target is not found in the array, return [-1, -1].
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 * 
 */
func searchRange(nums []int, target int) []int {
	// ✔ 88/88 cases passed (8 ms)
	// ✔ Your runtime beats 83.93 % of golang submissions
	// ✔ Your memory usage beats 50 % of golang submissions (4.1 MB)
	var first = binarySearch(nums, target)
	var last = binarySearch(nums, target + 1) - 1
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	} else {
		return []int{first, max(first, last)}
	}
}

func binarySearch(nums []int, target int) int {
	var l = 0
	// h 的边界条件。如果取最后一个索引，则判断条件需要改为 l<=h, l 依旧可能超限 todo
	// 如果 middle 是奇数，l<=h 会导致死循环
	var h = len(nums)
	for l < h {
		var m = l + (h-l)/2
		if nums[m] >= target {
			h = m
		} else {
			l = m + 1
		}
	}
	return l

}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}