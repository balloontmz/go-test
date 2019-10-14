/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 *
 * https://leetcode.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (53.27%)
 * Likes:    736
 * Dislikes: 39
 * Total Accepted:    137.4K
 * Total Submissions: 257.8K
 * Testcase Example:  '3\n7'
 *
 *
 * Find all possible combinations of k numbers that add up to a number n, given
 * that only numbers from 1 to 9 can be used and each combination should be a
 * unique set of numbers.
 *
 * Note:
 *
 *
 * All numbers will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: k = 3, n = 7
 * Output: [[1,2,4]]
 *
 *
 * Example 2:
 *
 *
 * Input: k = 3, n = 9
 * Output: [[1,2,6], [1,3,5], [2,3,4]]
 *
 *
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	// 18/18 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (1.9 MB)
	var combinations [][]int
	backtracking(k, n, 1, &[]int{}, &combinations)
	return combinations
}

func backtracking(k, n, start int, tempCombination *[]int, combinations *[][]int) {
	if k == 0 && n == 0 {
		var temp = make([]int, len(*tempCombination))
		copy(temp, *tempCombination)
		*combinations = append(*combinations, temp)
		return
	}
	if k == 0 || n == 0 {
		return
	}
	for i := start; i < 10; i++ {
		*tempCombination = append(*tempCombination, i)
		backtracking(k-1, n-i, i+1, tempCombination, combinations)
		*tempCombination = (*tempCombination)[:len(*tempCombination)-1]
	}
}

// @lc code=end

