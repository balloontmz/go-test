/*
 * @lc app=leetcode id=397 lang=golang
 *
 * [397] Integer Replacement
 *
 * https://leetcode.com/problems/integer-replacement/description/
 *
 * algorithms
 * Medium (32.04%)
 * Likes:    283
 * Dislikes: 277
 * Total Accepted:    48K
 * Total Submissions: 148.9K
 * Testcase Example:  '8'
 *
 * 
 * Given a positive integer n and you can do operations as follow:
 * 
 * 
 * 
 * 
 * If n is even, replace n with n/2.
 * If n is odd, you can replace n with either n + 1 or n - 1.
 * 
 * 
 * 
 * 
 * What is the minimum number of replacements needed for n to become 1?
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * Input:
 * 8
 * 
 * Output:
 * 3
 * 
 * Explanation:
 * 8 -> 4 -> 2 -> 1
 * 
 * 
 * 
 * Example 2:
 * 
 * Input:
 * 7
 * 
 * Output:
 * 4
 * 
 * Explanation:
 * 7 -> 8 -> 4 -> 2 -> 1
 * or
 * 7 -> 6 -> 3 -> 2 -> 1
 * 
 * 
 */

// @lc code=start
//整数替换,采用递归
// 47/47 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (1.9 MB)
func integerReplacement(n int) int {
	if n == 2 {
		return 1
	}
	var res int
	for n % 2 == 0 {
		n = n / 2
		res++
	}
    if n == 1 {
        return res
    }
	return min(integerReplacement(n+1), integerReplacement(n-1)) + res + 1
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
// @lc code=end

