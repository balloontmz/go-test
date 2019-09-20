/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (48.95%)
 * Likes:    925
 * Dislikes: 52
 * Total Accepted:    226.9K
 * Total Submissions: 455.9K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of 1 ... n.
 * 
 * Example:
 * 
 * 
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 * 
 * 
 */
func combine(n int, k int) [][]int {
	var combinations [][]int
	var combineList []int
	backtracking(combineList, combinations, 1, k, n)
	return combinations
}

func backtracking(combineList []int, combinations [][]int, start, k, n int)  {
	
}

