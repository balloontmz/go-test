/*
 * @lc app=leetcode id=188 lang=golang
 *
 * [188] Best Time to Buy and Sell Stock IV
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/description/
 *
 * algorithms
 * Hard (26.97%)
 * Likes:    980
 * Dislikes: 61
 * Total Accepted:    102.6K
 * Total Submissions: 379.3K
 * Testcase Example:  '2\n[2,4,1]'
 *
 * Say you have an array for which the i-th element is the price of a given
 * stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete at most k
 * transactions.
 *
 * Note:
 * You may not engage in multiple transactions at the same time (ie, you must
 * sell the stock before you buy again).
 *
 * Example 1:
 *
 *
 * Input: [2,4,1], k = 2
 * Output: 2
 * Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit
 * = 4-2 = 2.
 *
 *
 * Example 2:
 *
 *
 * Input: [3,2,6,5,0,3], k = 2
 * Output: 7
 * Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit
 * = 6-2 = 4.
 * Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 =
 * 3.
 *
 *
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	//只能进行 k 次的股票交易 -- 来自拷贝
	// 211/211 cases passed (4 ms)
	// Your runtime beats 78.18 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (3.8 MB)
	var n = len(prices)
	if k > n/2 { // 交易次数不可能超出的情况
		var maxProfit int
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				maxProfit += prices[i] - prices[i-1]
			}
		}
		return maxProfit
	}
	var maxProfits = make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		maxProfits[i] = make([]int, n)
	}
	for i := 1; i <= k; i++ {
		var localMax = maxProfits[i-1][0] - prices[0] // 第0天购买i-1次的最大收入
		for j := 1; j < n; j++ {
			//第j天进行i次交易的最大收益为 第j-1天进行i次交易的最大值和第j-1天持有的最大收益
			maxProfits[i][j] = max(maxProfits[i][j-1], localMax+prices[j])
			//第j天持有股票的最大收益
			localMax = max(localMax, maxProfits[i-1][j]-prices[j])
		}
	}
	return maxProfits[k][n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

