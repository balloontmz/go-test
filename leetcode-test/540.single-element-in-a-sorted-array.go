/*
 * @lc app=leetcode id=540 lang=golang
 *
 * [540] Single Element in a Sorted Array
 *
 * https://leetcode.com/problems/single-element-in-a-sorted-array/description/
 *
 * algorithms
 * Medium (57.44%)
 * Likes:    812
 * Dislikes: 64
 * Total Accepted:    60.6K
 * Total Submissions: 105.4K
 * Testcase Example:  '[1,1,2,3,3,4,4,8,8]'
 *
 * Given a sorted array consisting of only integers where every element appears
 * exactly twice except for one element which appears exactly once. Find this
 * single element that appears only once.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: [1,1,2,3,3,4,4,8,8]
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [3,3,7,7,10,11,11]
 * Output: 10
 * 
 * 
 * 
 * 
 * Note: Your solution should run in O(log n) time and O(1) space.
 * 
 */
func singleNonDuplicate(nums []int) int {
	var l = 0
	var r = len(nums)-1
	// 令 index 为 Single Element 在数组中的位置。在 index 之后，数组中原来存在的成对状态被改变。如果 m 为偶数，并且 m + 1 < index，那么 nums[m] == nums[m + 1]；m + 1 >= index，那么 nums[m] != nums[m + 1]。
	// 从上面的规律可以知道，如果 nums[m] == nums[m + 1]，那么 index 所在的数组位置为 [m + 2, h]，此时令 l = m + 2；如果 nums[m] != nums[m + 1]，那么 index 所在的数组位置为 [l, m]，此时令 h = m。
	// ✔ 12/12 cases passed (4 ms)
	// ✔ Your runtime beats 96.59 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (4.1 MB)
	for l < r {
		var m = l + (r-l)/2
		if m%2==1 {
			m--   // 保证 l，m，r 都在偶数位，使区间大小一直未奇数
		}
		if nums[m]==nums[m+1] {
			l = m + 2
		} else {
			r = m
		}
	}
	return nums[l]
}

