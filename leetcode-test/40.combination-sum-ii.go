import "sort"

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 *
 * https://leetcode.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (43.81%)
 * Likes:    1119
 * Dislikes: 51
 * Total Accepted:    256.3K
 * Total Submissions: 584.2K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * Given a collection of candidate numbers (candidates) and a target number
 * (target), find all unique combinations in candidates where the candidate
 * numbers sums to target.
 *
 * Each number in candidates may only be used once in the combination.
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
 * Input: candidates = [10,1,2,7,6,1,5], target = 8,
 * A solution set is:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,5,2,1,2], target = 5,
 * A solution set is:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	var combinations [][]int
	var hasVisited = make([]bool, len(candidates))
	sort.Ints(candidates)
	backtracking(&[]int{}, &combinations, hasVisited, 0, target, candidates)
	return combinations
}

func backtracking(
	tempCombination *[]int,
	combinations *[][]int,
	visited []bool,
	start, target int,
	candidates []int) {
	if target == 0 {
		var temp = make([]int, len(*tempCombination))
		copy(temp, *tempCombination)
		*combinations = append(*combinations, temp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if i != 0 && candidates[i] == candidates[i-1] && !visited[i-1] {
			continue // 重复数字的组合只会出现一次
		}
		if candidates[i] <= target {
			*tempCombination = append(*tempCombination, candidates[i])
			visited[i] = true
			backtracking(tempCombination, combinations, visited, i+1, target-candidates[i], candidates)
			visited[i] = false
			*tempCombination = (*tempCombination)[:len(*tempCombination)-1]
		}
	}
}

// @lc code=end

