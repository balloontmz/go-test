/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (28.76%)
 * Likes:    1033
 * Dislikes: 2472
 * Total Accepted:    379.7K
 * Total Submissions: 1.3M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (x^n).
 *
 * Example 1:
 *
 *
 * Input: 2.00000, 10
 * Output: 1024.00000
 *
 *
 * Example 2:
 *
 *
 * Input: 2.10000, 3
 * Output: 9.26100
 *
 *
 * Example 3:
 *
 *
 * Input: 2.00000, -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 * Note:
 *
 *
 * -100.0 < x < 100.0
 * n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]
 *
 *
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	// 304/304 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 75 % of golang submissions (2.1 MB)
	var res float64 = 1.0
	for i := n; i != 0; i /= 2 {
		if i%2 != 0 {
			res *= x
		}
		x *= x
	}
	if n > 0 {
		return res
	}
	return 1 / res
}

// func myPow(x float64, n int) float64 {
// 	//递归求 x 的 y 次方
// 	// 304/304 cases passed (4 ms)
// 	// Your runtime beats 20.35 % of golang submissions
// 	// Your memory usage beats 12.5 % of golang submissions (2.1 MB)
// 	if n == 0 {
// 		return 1
// 	}
// 	var half = myPow(x, n/2)
// 	if n%2 == 0 {
// 		return half * half
// 	}
// 	if n > 0 {
// 		return half * half * x
// 	}
// 	return half * half / x
// }

// @lc code=end

