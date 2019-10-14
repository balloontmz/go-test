import "sort"

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 *
 * https://leetcode.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (41.54%)
 * Likes:    1230
 * Dislikes: 47
 * Total Accepted:    272.1K
 * Total Submissions: 645.6K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a collection of numbers that might contain duplicates, return all
 * possible unique permutations.
 *
 * Example:
 *
 *
 * Input: [1,1,2]
 * Output:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 *
 *
 */
func permuteUnique(nums []int) [][]int {
	// 30/30 cases passed (8 ms)
	// Your runtime beats 95.28 % of golang submissions
	// Your memory usage beats 50 % of golang submissions (8 MB)
	var permutes [][]int
	var permuteList []int
	sort.Ints(nums)
	var hasVisited = make([]bool, len(nums))
	backtracking(&permuteList, &permutes, hasVisited, nums)
	return permutes
}

func backtracking(permuteList *[]int, permutes *[][]int, visited []bool, nums []int) {
	if len(*permuteList) == len(nums) {
		var temp = make([]int, len(*permuteList))
		copy(temp, *permuteList)
		*permutes = append(*permutes, temp)

		return
	}
	for i := 0; i < len(visited); i++ {
		if i != 0 && nums[i] == nums[i-1] && !(visited)[i-1] {
			continue
		}
		if (visited)[i] {
			continue
		}
		(visited)[i] = true
		*permuteList = append(*permuteList, nums[i])
		backtracking(permuteList, permutes, visited, nums)
		*permuteList = (*permuteList)[:len(*permuteList)-1]
		(visited)[i] = false
	}
}

