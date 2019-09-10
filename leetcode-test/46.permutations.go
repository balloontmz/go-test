/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 *
 * https://leetcode.com/problems/permutations/description/
 *
 * algorithms
 * Medium (56.56%)
 * Likes:    2382
 * Dislikes: 75
 * Total Accepted:    430.6K
 * Total Submissions: 752.2K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a collection of distinct integers, return all possible permutations.
 * 
 * Example:
 * 
 * 
 * Input: [1,2,3]
 * Output:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 * 
 * 
 */
var permutes [][]int
var permuteList []int

func permute(nums []int) [][]int {
	// ✔ 25/25 cases passed (4 ms)
	// ✔ Your runtime beats 91.97 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (7 MB)
	var hasVisitedOrigin = (make([]bool, len(nums)))
	var hasVisited = &hasVisitedOrigin
	// 此处注意第二个测试用例影响第一个
	permutes = make([][]int, 0)
	backtracking(hasVisited, nums)
	return permutes
}

func backtracking(visited *[]bool, nums []int)  {
	if len(permuteList) == len(nums) {
		// 此处注意如果采用 permuteList 。保存的是地址，那么所有值指向的都是最后一个值
		temp := make([]int, len(nums))
		copy(temp, permuteList)
		permutes = append(permutes, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if (*visited)[i] {
			continue
		}
		(*visited)[i] = true
		permuteList = append(permuteList, nums[i])
		backtracking(visited, nums)
		permuteList = permuteList[:len(permuteList)-1]
		(*visited)[i] = false
	}
}