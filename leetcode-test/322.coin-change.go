/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 *
 * https://leetcode.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (32.11%)
 * Likes:    2396
 * Dislikes: 86
 * Total Accepted:    262.1K
 * Total Submissions: 811.3K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * You are given coins of different denominations and a total amount of money
 * amount. Write a function to compute the fewest number of coins that you need
 * to make up that amount. If that amount of money cannot be made up by any
 * combination of the coins, return -1.
 *
 * Example 1:
 *
 *
 * Input: coins = [1, 2, 5], amount = 11
 * Output: 3
 * Explanation: 11 = 5 + 5 + 1
 *
 * Example 2:
 *
 *
 * Input: coins = [2], amount = 3
 * Output: -1
 *
 *
 * Note:
 * You may assume that you have an infinite number of each kind of coin.
 *
 */

// @lc code=start
//动态规划解最少找零硬币数
func coinChange(coins []int, amount int) int {
	// 182/182 cases passed (12 ms)
	// Your runtime beats 59.83 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (5.9 MB)
	if amount == 0 {
		return 0
	}
	var dp = make([]int, amount+1)
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if i == coin {
				dp[i] = 1
			} else if dp[i] == 0 && dp[i-coin] != 0 {
				dp[i] = dp[i-coin] + 1
			} else if dp[i-coin] != 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
			// 还差一种 dp[i] == 0 && dp[i-icon] == 0
			// 还差一种 dp[i] != 0 && dp[i-icon] == 0
		}
	}
	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

