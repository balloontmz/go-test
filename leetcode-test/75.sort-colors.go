/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 *
 * https://leetcode.com/problems/sort-colors/description/
 *
 * algorithms
 * Medium (42.27%)
 * Likes:    1792
 * Dislikes: 163
 * Total Accepted:    337.6K
 * Total Submissions: 789.7K
 * Testcase Example:  '[2,0,2,1,1,0]'
 *
 * Given an array with n objects colored red, white or blue, sort them in-place
 * so that objects of the same color are adjacent, with the colors in the order
 * red, white and blue.
 * 
 * Here, we will use the integers 0, 1, and 2 to represent the color red,
 * white, and blue respectively.
 * 
 * Note: You are not suppose to use the library's sort function for this
 * problem.
 * 
 * Example:
 * 
 * 
 * Input: [2,0,2,1,1,0]
 * Output: [0,0,1,1,2,2]
 * 
 * Follow up:
 * 
 * 
 * A rather straight forward solution is a two-pass algorithm using counting
 * sort.
 * First, iterate the array counting number of 0's, 1's, and 2's, then
 * overwrite array with total number of 0's, then 1's and followed by 2's.
 * Could you come up with a one-pass algorithm using only constant space?
 * 
 * 
 */
func sortColors(nums []int)  {
	// ✔ 87/87 cases passed (0 ms)
	// ✔ Your runtime beats 100 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (2.3 MB)
	zero, one, two := -1, 0, len(nums)
	for one < two {
		if nums[one] == 0 {
			// 此处由于交换过来的值肯定是 1或者 0，所以one 自加。
			zero++
			swap(nums, zero, one)
			one++
		} else if nums[one] == 2 {
			// 此处由于交换过来的 two 的值 依旧未知，所以 one 并不自加
			two--
			swap(nums, two, one)
		} else {
			one++
		}
	}
	return 
}

func swap(nums []int, i, j int)  {
	// 其实可以直接赋值，此处理清流程，所以不简化
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
