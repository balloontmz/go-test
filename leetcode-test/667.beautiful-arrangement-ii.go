/*
 * @lc app=leetcode id=667 lang=golang
 *
 * [667] Beautiful Arrangement II
 *
 * https://leetcode.com/problems/beautiful-arrangement-ii/description/
 *
 * algorithms
 * Medium (52.78%)
 * Likes:    261
 * Dislikes: 554
 * Total Accepted:    20.5K
 * Total Submissions: 38.8K
 * Testcase Example:  '3\n2'
 *
 *
 * Given two integers n and k, you need to construct a list which contains n
 * different positive integers ranging from 1 to n and obeys the following
 * requirement:
 *
 * Suppose this list is [a1, a2, a3, ... , an], then the list [|a1 - a2|, |a2 -
 * a3|, |a3 - a4|, ... , |an-1 - an|] has exactly k distinct integers.
 *
 *
 *
 * If there are multiple answers, print any of them.
 *
 *
 * Example 1:
 *
 * Input: n = 3, k = 1
 * Output: [1, 2, 3]
 * Explanation: The [1, 2, 3] has three different positive integers ranging
 * from 1 to 3, and the [1, 1] has exactly 1 distinct integer: 1.
 *
 *
 *
 * Example 2:
 *
 * Input: n = 3, k = 2
 * Output: [1, 3, 2]
 * Explanation: The [1, 3, 2] has three different positive integers ranging
 * from 1 to 3, and the [2, 1] has exactly 2 distinct integers: 1 and 2.
 *
 *
 *
 * Note:
 *
 * The n and k are in the range 1
 *
 *
 */

// @lc code=start
//数组相邻差值的个数
func constructArray(n int, k int) []int {
	// 68/68 cases passed (80 ms)
	// Your runtime beats 93.75 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (8 MB)
	var res = make([]int, n)
	res[0] = 1
	var iter = k
	for i := 1; i <= k; i++ {
		if i%2 == 1 {
			res[i] = res[i-1] + iter
		} else {
			res[i] = res[i-1] - iter
		}
		iter--
	}
	for i := k + 1; i < n; i++ {
		res[i] = i + 1
	}
	return res
}

// @lc code=end

