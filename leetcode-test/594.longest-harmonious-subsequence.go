/*
 * @lc app=leetcode id=594 lang=golang
 *
 * [594] Longest Harmonious Subsequence
 *
 * https://leetcode.com/problems/longest-harmonious-subsequence/description/
 *
 * algorithms
 * Easy (44.81%)
 * Likes:    486
 * Dislikes: 68
 * Total Accepted:    43.4K
 * Total Submissions: 96.9K
 * Testcase Example:  '[1,3,2,2,5,2,3,7]'
 *
 * We define a harmounious array as an array where the difference between its
 * maximum value and its minimum value is exactly 1.
 *
 * Now, given an integer array, you need to find the length of its longest
 * harmonious subsequence among all its possible subsequences.
 *
 * Example 1:
 *
 *
 * Input: [1,3,2,2,5,2,3,7]
 * Output: 5
 * Explanation: The longest harmonious subsequence is [3,2,2,2,3].
 *
 *
 *
 *
 * Note: The length of the input array will not exceed 20,000.
 *
 */

// @lc code=start
//查找数组最长和谐序列的长度
// 201/201 cases passed (56 ms)
// Your runtime beats 74.19 % of golang submissions
// Your memory usage beats 100 % of golang submissions (7.6 MB)
func findLHS(nums []int) int {
	var long int
	var m = make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if m[k+1] != 0 {
			long = max(long, v+m[k+1])
		}
	}
	return long
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

