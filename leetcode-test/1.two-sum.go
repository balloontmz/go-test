/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (44.71%)
 * Likes:    12444
 * Dislikes: 434
 * Total Accepted:    2.3M
 * Total Submissions: 5.1M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * Example:
 *
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 *
 */

// @lc code=start
//o(n)时间复杂度，o(n)空间复杂度 -- 双指针 o(nlogn)时间复杂度 o(1)空间复杂度
// 29/29 cases passed (4 ms)
// Your runtime beats 95.05 % of golang submissions
// Your memory usage beats 13.46 % of golang submissions (3.8 MB)
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, 0)
	for k, num := range nums {
		m[num] = k
	}
	for i := 0; i < len(nums); i++ {
		if m[target-nums[i]] != 0 && m[target-nums[i]] != i {
			return []int{i, m[target-nums[i]]}
		}
	}
	return nil
}

// @lc code=end

