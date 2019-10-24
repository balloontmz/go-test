/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 *
 * https://leetcode.com/problems/house-robber-ii/description/
 *
 * algorithms
 * Medium (35.64%)
 * Likes:    1110
 * Dislikes: 39
 * Total Accepted:    134.5K
 * Total Submissions: 377.2K
 * Testcase Example:  '[2,3,2]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed. All houses at this place are
 * arranged in a circle. That means the first house is the neighbor of the last
 * one. Meanwhile, adjacent houses have security system connected and it will
 * automatically contact the police if two adjacent houses were broken into on
 * the same night.
 *
 * Given a list of non-negative integers representing the amount of money of
 * each house, determine the maximum amount of money you can rob tonight
 * without alerting the police.
 *
 * Example 1:
 *
 *
 * Input: [2,3,2]
 * Output: 3
 * Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money
 * = 2),
 * because they are adjacent houses.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money =
 * 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 */

// @lc code=start
func rob(nums []int) int {
	// 74/74 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2 MB)
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(subRob(nums, 0, n-1), subRob(nums, 1, n))
}

func subRob(nums []int, first, end int) int {
	var p1 = 0
	var p2 = 0
	var cur = 0
	for i := first; i < end; i++ {
		cur = max(p2+nums[i], p1)
		p2 = p1
		p1 = cur
	}
	return cur
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

