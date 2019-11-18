/*
 * @lc app=leetcode id=260 lang=golang
 *
 * [260] Single Number III
 *
 * https://leetcode.com/problems/single-number-iii/description/
 *
 * algorithms
 * Medium (58.59%)
 * Likes:    1057
 * Dislikes: 85
 * Total Accepted:    119.5K
 * Total Submissions: 203.5K
 * Testcase Example:  '[1,2,1,3,2,5]'
 *
 * Given an array of numbers nums, in which exactly two elements appear only
 * once and all the other elements appear exactly twice. Find the two elements
 * that appear only once.
 *
 * Example:
 *
 *
 * Input:  [1,2,1,3,2,5]
 * Output: [3,5]
 *
 * Note:
 *
 *
 * The order of the result is not important. So in the above example, [5, 3] is
 * also correct.
 * Your algorithm should run in linear runtime complexity. Could you implement
 * it using only constant space complexity?
 *
 */

// @lc code=start
//找出数组中不重复的两个数
func singleNumber(nums []int) []int {
	// 30/30 cases passed (4 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (4.2 MB)
	var diff int
	for _, num := range nums {
		diff ^= num // 不同的两个数的异或
	}
	diff &= -diff // 获取最低的不同的位
	var ret = make([]int, 2)
	for _, num := range nums {
		//除了有差异的这两个数，其余都是成对存在于0或者1的
		if (num & diff) == 0 {
			ret[0] ^= num
		} else {
			ret[1] ^= num
		}
	}
	return ret

}

// @lc code=end

