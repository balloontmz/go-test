/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 *
 * https://leetcode.com/problems/max-consecutive-ones/description/
 *
 * algorithms
 * Easy (55.74%)
 * Likes:    431
 * Dislikes: 336
 * Total Accepted:    158.9K
 * Total Submissions: 285K
 * Testcase Example:  '[1,0,1,1,0,1]'
 *
 * Given a binary array, find the maximum number of consecutive 1s in this
 * array.
 *
 * Example 1:
 *
 * Input: [1,1,0,1,1,1]
 * Output: 3
 * Explanation: The first two digits or the last three digits are consecutive
 * 1s.
 * ⁠   The maximum number of consecutive 1s is 3.
 *
 *
 *
 * Note:
 *
 * The input array will only contain 0 and 1.
 * The length of input array is a positive integer and will not exceed 10,000
 *
 *
 */

// @lc code=start
//找出数组中最长连续的1
func findMaxConsecutiveOnes(nums []int) int {
	// 41/41 cases passed (40 ms)
	// Your runtime beats 52.54 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.3 MB)
	var cur, max = 0, 0
	for _, num := range nums {
		if num == 0 {
			cur = 0
		} else {
			cur += 1
		}
		max = maxF(max, cur)
	}
	return max
}

func maxF(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

