/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (54.67%)
 * Likes:    2091
 * Dislikes: 182
 * Total Accepted:    456.6K
 * Total Submissions: 835K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 *
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 */

// @lc code=start
//1. 多数投票算法 复杂度为 o(n)
//2. 先对数组排序，最中间那个数出现次数一定多于 n / 2。复杂度为 o(n2)
func majorityElement(nums []int) int {
	// 44/44 cases passed (16 ms)
	// Your runtime beats 93.66 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (5.9 MB)
	var cnt = 0
	var major = nums[0]
	for _, num := range nums {
		if cnt == 0 {
			major = num
		}
		if major == num {
			cnt += 1
		} else {
			cnt -= 1
		}
	}
	return major
	// 44/44 cases passed (16 ms)
	// Your runtime beats 93.66 % of golang submissions
	// Your memory usage beats 75 % of golang submissions (5.9 MB)
	// sort.Ints(nums)
	// return nums[len(nums)/2]
}

// @lc code=end

