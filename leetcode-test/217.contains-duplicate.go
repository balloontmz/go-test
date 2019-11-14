/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 *
 * https://leetcode.com/problems/contains-duplicate/description/
 *
 * algorithms
 * Easy (53.80%)
 * Likes:    510
 * Dislikes: 588
 * Total Accepted:    418K
 * Total Submissions: 775.7K
 * Testcase Example:  '[1,2,3,1]'
 *
 * Given an array of integers, find if the array contains any duplicates.
 *
 * Your function should return true if any value appears at least twice in the
 * array, and it should return false if every element is distinct.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,1]
 * Output: true
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4]
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: [1,1,1,3,3,4,3,2,4,2]
 * Output: true
 *
 */

// @lc code=start
//判断数组中是否存在重复元素
// 18/18 cases passed (24 ms)
// Your runtime beats 50.29 % of golang submissions
// Your memory usage beats 25 % of golang submissions (10.9 MB)
func containsDuplicate(nums []int) bool {
	var m = make(map[int]int, 0)
	for _, v := range nums {
		if m[v] != 0 {
			return true
		}
		m[v]++
	}
	return false
}

// @lc code=end

