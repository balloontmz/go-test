/*
 * @lc app=leetcode id=645 lang=golang
 *
 * [645] Set Mismatch
 *
 * https://leetcode.com/problems/set-mismatch/description/
 *
 * algorithms
 * Easy (41.23%)
 * Likes:    483
 * Dislikes: 260
 * Total Accepted:    57.2K
 * Total Submissions: 138.8K
 * Testcase Example:  '[1,2,2,4]'
 *
 *
 * The set S originally contains numbers from 1 to n. But unfortunately, due to
 * the data error, one of the numbers in the set got duplicated to another
 * number in the set, which results in repetition of one number and loss of
 * another number.
 *
 *
 *
 * Given an array nums representing the data status of this set after the
 * error. Your task is to firstly find the number occurs twice and then find
 * the number that is missing. Return them in the form of an array.
 *
 *
 *
 * Example 1:
 *
 * Input: nums = [1,2,2,4]
 * Output: [2,3]
 *
 *
 *
 * Note:
 *
 * The given array size will in the range [2, 10000].
 * The given array's numbers won't have any order.
 *
 *
 */

// @lc code=start
//找出错误的数
func findErrorNums(nums []int) []int {
	// 49/49 cases passed (24 ms)
	// Your runtime beats 98.46 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.4 MB)
	var n = len(nums)
	//此种排序只适用于当前环境！！！
	for i := 0; i < n; i++ {
		for nums[i] != i+1 && nums[nums[i]-1] != nums[i] { // 当两个位置的数位置都不对时
			swap(nums, i, nums[i]-1)
		}
	}
	for i := 0; i < n; i++ {
		if nums[i]-1 != i {
			return []int{nums[i], i + 1}
		}
	}
	return nil
}

func swap(nums []int, i, j int) {
	var temp = nums[i]
	nums[i] = nums[j]
	nums[j] = temp
	return
}

// @lc code=end

