/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 *
 * https://leetcode.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (41.50%)
 * Likes:    3151
 * Dislikes: 73
 * Total Accepted:    275K
 * Total Submissions: 661.2K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * Given an unsorted array of integers, find the length of longest increasing
 * subsequence.
 *
 * Example:
 *
 *
 * Input: [10,9,2,5,3,7,101,18]
 * Output: 4
 * Explanation: The longest increasing subsequence is [2,3,7,101], therefore
 * the length is 4.
 *
 * Note:
 *
 *
 * There may be more than one LIS combination, it is only necessary for you to
 * return the length.
 * Your algorithm should run in O(n^2) complexity.
 *
 *
 * Follow up: Could you improve it to O(n log n) time complexity?
 *
 */

// @lc code=start

// 二分法解最长递增子序列 -- 更高性能
func lengthOfLIS(nums []int) int {
	// 24/24 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2.4 MB)
	n := len(nums)
	var tails = make([]int, n)
	len := 0
	for _, num := range nums {
		index := binarySearch(tails, len, num)
		tails[index] = num
		if index == len {
			len++
		}
	}
	return len
}

func binarySearch(tails []int, len, key int) int {
	l := 0
	r := len
	for l < r {
		mid := (l + r) / 2
		if tails[mid] == key {
			return mid
		} else if tails[mid] < key {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// 动态规划解最长递增子序列
// func lengthOfLIS(nums []int) int {
// 	// 24/24 cases passed (16 ms)
// 	// Your runtime beats 15.42 % of golang submissions
// 	// Your memory usage beats 60 % of golang submissions (2.4 MB)
// 	n := len(nums)
// 	if n <= 1 {
// 		return n
// 	}
// 	var dp = make([]int, n)
// 	for i := 0; i < n; i++ {
// 		max := 1
// 		for j := 0; j < i; j++ {
// 			if nums[i] > nums[j] {
// 				max = maxf(max, dp[j]+1)
// 			}
// 		}
// 		dp[i] = max
// 	}
// 	ret := 0
// 	for i := 0; i < n; i++ {
// 		ret = maxf(ret, dp[i])
// 	}
// 	return ret
// }

// func maxf(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// @lc code=end

