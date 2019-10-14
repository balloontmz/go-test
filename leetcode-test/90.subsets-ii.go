import "sort"

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 *
 * https://leetcode.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (43.94%)
 * Likes:    1110
 * Dislikes: 54
 * Total Accepted:    228.2K
 * Total Submissions: 518.9K
 * Testcase Example:  '[1,2,2]'
 *
 * Given a collection of integers that might contain duplicates, nums, return
 * all possible subsets (the power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: [1,2,2]
 * Output:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 *
 *
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	// 19/19 cases passed (4 ms)
	// Your runtime beats 82.58 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.9 MB)
	sort.Ints(nums) // 排序 -- 用于对重复元素不重复取组合
	var subsets [][]int
	var tempSubset []int
	var hasVisited = make([]bool, len(nums))

	for i := 0; i <= len(nums); i++ {
		backtracking(0, &tempSubset, &subsets, hasVisited, i, nums)
	}
	return subsets
}

func backtracking(start int, tempSubset *[]int, subsets *[][]int, visited []bool, size int, nums []int) {
	if len(*tempSubset) == size {
		var temp = make([]int, len(*tempSubset))
		copy(temp, *tempSubset)
		*subsets = append(*subsets, temp)
	}
	for i := start; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}
		*tempSubset = append(*tempSubset, nums[i])
		visited[i] = true
		backtracking(i+1, tempSubset, subsets, visited, size, nums)
		visited[i] = false
		*tempSubset = (*tempSubset)[:len(*tempSubset)-1]
	}
}

// @lc code=end

