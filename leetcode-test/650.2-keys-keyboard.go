/*
 * @lc app=leetcode id=650 lang=golang
 *
 * [650] 2 Keys Keyboard
 *
 * https://leetcode.com/problems/2-keys-keyboard/description/
 *
 * algorithms
 * Medium (47.31%)
 * Likes:    774
 * Dislikes: 54
 * Total Accepted:    40.2K
 * Total Submissions: 84.7K
 * Testcase Example:  '3'
 *
 * Initially on a notepad only one character 'A' is present. You can perform
 * two operations on this notepad for each step:
 *
 *
 * Copy All: You can copy all the characters present on the notepad (partial
 * copy is not allowed).
 * Paste: You can paste the characters which are copied last time.
 *
 *
 *
 *
 * Given a number n. You have to get exactly n 'A' on the notepad by performing
 * the minimum number of steps permitted. Output the minimum number of steps to
 * get n 'A'.
 *
 * Example 1:
 *
 *
 * Input: 3
 * Output: 3
 * Explanation:
 * Intitally, we have one character 'A'.
 * In step 1, we use Copy All operation.
 * In step 2, we use Paste operation to get 'AA'.
 * In step 3, we use Paste operation to get 'AAA'.
 *
 *
 *
 *
 * Note:
 *
 *
 * The n will be in the range [1, 1000].
 *
 *
 *
 *
 */

// @lc code=start
func minSteps(n int) int {
	// 126/126 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2.3 MB)
	var dp = make([]int, n+1)
	s := mySqrt(n)
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 2; j <= s; j++ {
			if i%j == 0 {
				dp[i] = dp[j] + dp[i/j]
				break
			}
		}
	}
	return dp[n]
}

// func minSteps(n int) int {
// 	// 126/126 cases passed (0 ms)
// 	// Your runtime beats 100 % of golang submissions
// 	// Your memory usage beats 100 % of golang submissions (1.9 MB)
// 	if n == 1 {
// 		return 0
// 	}
// 	for i := 2; i <= mySqrt(n); i++ { // 从小到大遍历，
// 		if n%i == 0 {
// 			return i + minSteps(n/i)
// 		}
// 	}
// 	return n
// }

func mySqrt(x int) int {
	var l = 0
	var r = x
	for l <= r {
		m := l + (r-l)/2
		if m*m == x {
			return m
		} else if m*m < x {
			l = m + 1
		} else if m*m > x {
			r = m - 1
		}
	}
	return r
}

// @lc code=end

