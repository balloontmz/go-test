/*
 * @lc app=leetcode id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 *
 * https://leetcode.com/problems/maximum-average-subarray-i/description/
 *
 * algorithms
 * Easy (41.13%)
 * Likes:    654
 * Dislikes: 108
 * Total Accepted:    72.9K
 * Total Submissions: 177.1K
 * Testcase Example:  '[1,12,-5,-6,50,3]\n4'
 *
 * Given an array consisting of n integers, find the contiguous subarray of
 * given length k that has the maximum average value. And you need to output
 * the maximum average value.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,12,-5,-6,50,3], k = 4
 * Output: 12.75
 * Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 1 <= k <= n <= 30,000.
 * Elements of the given array will be in the range [-10,000, 10,000].
 * 
 * 
 * 
 * 
 */

// @lc code=start
// 123/123 cases passed (112 ms)
// Your runtime beats 97.92 % of golang submissions
// Your memory usage beats 100 % of golang submissions (6.7 MB)
func findMaxAverage(nums []int, k int) float64 {
	var sum = 0
	var maxSUm = 0
    for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSUm = sum
	for i := k; i < len(nums); i++ {
		sum += nums[i]
		sum -= nums[i-k]
		maxSUm = max(maxSUm, sum)
	}
	return float64(maxSUm) / float64(k)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

