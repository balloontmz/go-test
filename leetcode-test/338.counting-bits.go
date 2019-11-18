/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 *
 * https://leetcode.com/problems/counting-bits/description/
 *
 * algorithms
 * Medium (65.95%)
 * Likes:    1746
 * Dislikes: 121
 * Total Accepted:    201.5K
 * Total Submissions: 305.1K
 * Testcase Example:  '2'
 *
 * Given a non negative integer number num. For every numbers i in the range 0
 * ≤ i ≤ num calculate the number of 1's in their binary representation and
 * return them as an array.
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: [0,1,1]
 *
 * Example 2:
 *
 *
 * Input: 5
 * Output: [0,1,1,2,1,2]
 *
 *
 * Follow up:
 *
 *
 * It is very easy to come up with a solution with run time
 * O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a
 * single pass?
 * Space complexity should be O(n).
 * Can you do it like a boss? Do it without using any builtin function like
 * __builtin_popcount in c++ or in any other language.
 *
 */

// @lc code=start
//计算数的二进制表示中1的个数
func countBits(num int) []int {
	//动态规划、利用i&(i-1)为去除i的最低位1的数的特性
	// 15/15 cases passed (860 ms)
	// Your runtime beats 42.57 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (8.1 MB)
	var dp = make([]int, num+1)
	for i := 1; i < num+1; i++ {
		dp[i] = dp[i&(i-1)] + 1
	}
	return dp
}

// @lc code=end

