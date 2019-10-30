/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 *
 * https://leetcode.com/problems/target-sum/description/
 *
 * algorithms
 * Medium (45.83%)
 * Likes:    1698
 * Dislikes: 81
 * Total Accepted:    117.2K
 * Total Submissions: 255.3K
 * Testcase Example:  '[1,1,1,1,1]\n3'
 *
 *
 * You are given a list of non-negative integers, a1, a2, ..., an, and a
 * target, S. Now you have 2 symbols + and -. For each integer, you should
 * choose one from + and - as its new symbol.
 * ⁠
 *
 * Find out how many ways to assign symbols to make sum of integers equal to
 * target S.
 *
 *
 * Example 1:
 *
 * Input: nums is [1, 1, 1, 1, 1], S is 3.
 * Output: 5
 * Explanation:
 *
 * -1+1+1+1+1 = 3
 * +1-1+1+1+1 = 3
 * +1+1-1+1+1 = 3
 * +1+1+1-1+1 = 3
 * +1+1+1+1-1 = 3
 *
 * There are 5 ways to assign symbols to make the sum of nums be target 3.
 *
 *
 *
 * Note:
 *
 * The length of the given array is positive and will not exceed 20.
 * The sum of elements in the given array will not exceed 1000.
 * Your output answer is guaranteed to be fitted in a 32-bit integer.
 *
 *
 */

// @lc code=start
//dfs(深度优先遍历)解将一部分数变为负数,求和能组成目标数的序列个数
// 没有复用计算结果,性能低!!!
// type Dfs struct {
// 	Nums []int
// 	n    int
// }

var Nums []int

func findTargetSumWays(nums []int, S int) int {
	// 139/139 cases passed (496 ms)
	// Your runtime beats 34.93 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2.1 MB)
	// 139/139 cases passed (644 ms)
	// Your runtime beats 20.55 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2.1 MB)
	// ret := dfs(0, S, nums)
	Nums = nums
	// d := Dfs{nums, len(nums)}
	return dfs(0, S)
	// ret := dfs(0, S)
	// return ret
}

func dfs(index, target int) int {
	if index == len(Nums) {
		if target == 0 {
			return 1
		}
		return 0
	}
	return dfs(index+1, target-Nums[index]) + dfs(index+1, target+Nums[index])
	// ret := d.dfs(index+1, target-Nums[index]) + d.dfs(index+1, target+Nums[index])
	// return ret
}

//动态规划解将一部分数变为负数,求和能组成目标数的序列个数
// func findTargetSumWays(nums []int, S int) int {
// 	// 139/139 cases passed (0 ms)
// 	// Your runtime beats 100 % of golang submissions
// 	// Your memory usage beats 100 % of golang submissions (2.4 MB)
// 	sum := calSum(nums)
// 	// 和不为偶数或者和小于目标都不可能存在解
// 	if sum < S || (sum+S)%2 != 0 {
// 		return 0
// 	}
// 	T := (sum + S) / 2 // 实际问题就是存在子序列之和为 T
// 	var dp = make([]int, T+1)
// 	dp[0] = 1
// 	for _, num := range nums {
// 		// 倒序,因为可能存在小的数已经使用过一次的情况!!! -- 防止多次使用
// 		for j := T; j >= num; j-- {
// 			dp[j] = dp[j] + dp[j-num]
// 		}
// 	}
// 	return dp[T]
// }

// func calSum(nums []int) int {
// 	ret := 0
// 	for _, v := range nums {
// 		ret += v
// 	}
// 	return ret
// }

// @lc code=end

