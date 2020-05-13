/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 *
 * https://leetcode.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (43.76%)
 * Likes:    4106
 * Dislikes: 130
 * Total Accepted:    280.6K
 * Total Submissions: 641.4K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * Given an array of integers and an integer k, you need to find the total
 * number of continuous subarrays whose sum equals to k.
 * 
 * Example 1:
 * 
 * 
 * Input:nums = [1,1,1], k = 2
 * Output: 2
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The length of the array is in range [1, 20,000].
 * The range of numbers in the array is [-1000, 1000] and the range of the
 * integer k is [-1e7, 1e7].
 * 
 * 
 */

// @lc code=start
//哈希表
// 81/81 cases passed (20 ms)
// Your runtime beats 51.72 % of golang submissions
// Your memory usage beats 100 % of golang submissions (6.3 MB)
func subarraySum(nums []int, k int) int {
	var M = make(map[int]int, 0)
	var sum = 0
	var res = 0
	M[0] = 1 // 对于所有sum=k 都自加一次
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res += M[sum-k]
		M[sum] += 1 // sum 数自加 1
	}
	return res
}
// @lc code=end

