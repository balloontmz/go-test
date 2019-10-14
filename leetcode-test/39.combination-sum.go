/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 *
 * https://leetcode.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (51.13%)
 * Likes:    2525
 * Dislikes: 75
 * Total Accepted:    405.3K
 * Total Submissions: 792.6K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * Given a set of candidate numbers (candidates) (without duplicates) and a
 * target number (target), find all unique combinations in candidates where the
 * candidate numbers sums to target.
 *
 * The same repeated number may be chosen from candidates unlimited number of
 * times.
 *
 * Note:
 *
 *
 * All numbers (including target) will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [2,3,6,7], target = 7,
 * A solution set is:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,3,5], target = 8,
 * A solution set is:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 *
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	// 168/168 cases passed (4 ms)
	// Your runtime beats 86.57 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (4 MB)
	var combinations [][]int
	backtracking(&[]int{}, &combinations, 0, target, candidates)
	return combinations
}

func backtracking(tempCombination *[]int, combinations *[][]int, start, target int, candidates []int) {
	if target == 0 {
		var temp = make([]int, len(*tempCombination))
		copy(temp, *tempCombination)
		*combinations = append(*combinations, temp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] <= target {
			*tempCombination = append(*tempCombination, candidates[i])
			backtracking(tempCombination, combinations, i, target-candidates[i], candidates)
			*tempCombination = (*tempCombination)[:len(*tempCombination)-1]
		}
	}
}

// @lc code=end

