/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 *
 * https://leetcode.com/problems/partition-equal-subset-sum/description/
 *
 * algorithms
 * Medium (41.80%)
 * Likes:    1597
 * Dislikes: 49
 * Total Accepted:    116K
 * Total Submissions: 278.4K
 * Testcase Example:  '[1,5,11,5]'
 *
 * Given a non-empty array containing only positive integers, find if the array
 * can be partitioned into two subsets such that the sum of elements in both
 * subsets is equal.
 *
 * Note:
 *
 *
 * Each of the array element will not exceed 100.
 * The array size will not exceed 200.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1, 5, 11, 5]
 *
 * Output: true
 *
 * Explanation: The array can be partitioned as [1, 5, 5] and [11].
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1, 2, 3, 5]
 *
 * Output: false
 *
 * Explanation: The array cannot be partitioned into equal sum subsets.
 *
 *
 *
 *
 */

// @lc code=start
//动态规划解是否能取出子数组和为父数组的一半
func canPartition(nums []int) bool {
	// 105/105 cases passed (8 ms)
	// Your runtime beats 86.92 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2.5 MB)
	sum := calSum(nums)
	if sum%2 != 0 {
		return false
	}
	S := sum / 2
	var dp = make([]bool, S+1)
	dp[0] = true
	for _, v := range nums {
		for i := S; i >= v; i-- {
			dp[i] = (dp[i] || dp[i-v])
		}
	}
	return dp[S]
}

func calSum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// @lc code=end

