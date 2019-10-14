/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 *
 * https://leetcode.com/problems/subsets/description/
 *
 * algorithms
 * Medium (55.53%)
 * Likes:    2442
 * Dislikes: 58
 * Total Accepted:    427.9K
 * Total Submissions: 769.5K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a set of distinct integers, nums, return all possible subsets (the
 * power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: nums = [1,2,3]
 * Output:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */

// @lc code=start
func subsets(nums []int) [][]int {
	// 10/10 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 50 % of golang submissions (6.8 MB)
	var subsets [][]int
	var tempSubset []int
	for size := 0; size <= len(nums); size++ {
		backtracking(0, &tempSubset, &subsets, size, nums) // 不同的子集大小
	}
	return subsets
}

func backtracking(start int, tempSubset *[]int, subsets *[][]int, size int, nums []int) {
	if len(*tempSubset) == size {
		var temp = make([]int, len(*tempSubset))
		copy(temp, *tempSubset)
		*subsets = append(*subsets, temp)
		return
	}
	for i := start; i < len(nums); i++ {
		*tempSubset = append(*tempSubset, nums[i])
		backtracking(i+1, tempSubset, subsets, size, nums) // 此处用 i 而不是用 start,因为 i 为开始元素,而 start 只是开始元素的开始!!!
		*tempSubset = (*tempSubset)[:len(*tempSubset)-1]
	}
}

// @lc code=end

