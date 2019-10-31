/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
 *
 * algorithms
 * Medium (44.95%)
 * Likes:    1646
 * Dislikes: 62
 * Total Accepted:    108.6K
 * Total Submissions: 241.2K
 * Testcase Example:  '[1,2,3,0,2]'
 *
 * Say you have an array for which the i^th element is the price of a given
 * stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete as many
 * transactions as you like (ie, buy one and sell one share of the stock
 * multiple times) with the following restrictions:
 *
 *
 * You may not engage in multiple transactions at the same time (ie, you must
 * sell the stock before you buy again).
 * After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1
 * day)
 *
 *
 * Example:
 *
 *
 * Input: [1,2,3,0,2]
 * Output: 3
 * Explanation: transactions = [buy, sell, cooldown, buy, sell]
 *
 */

// @lc code=start
//需要冷却期的股票交易 -- 维护 4 种状态的情况.空间复杂度为 o1
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func maxProfit(prices []int) int {
	// 211/211 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2.3 MB)
	var buy, preBuy, sell, preSell int
	buy = INT_MIN
	for _, price := range prices {
		preBuy = buy
		buy = max(preSell-price, preBuy)
		preSell = sell
		sell = max(preBuy+price, preSell)
	}
	return sell
}

//需要冷却期的股票交易 -- 维护 2种状态的情况,此种情况应该更优!!!
// func maxProfit(prices []int) int {
// 	// 211/211 cases passed (0 ms)
// 	// Your runtime beats 100 % of golang submissions
// 	// Your memory usage beats 100 % of golang submissions (2.4 MB)
// 	N := len(prices)
// 	if N == 0 {
// 		return 0
// 	}
// 	var hold = make([]int, N) // 某天持有股票的最大收益
// 	var sell = make([]int, N) // 某天空手的最大收益
// 	hold[0] = -prices[0]
// 	// sell[0] = 0

// 	for i := 1; i < N; i++ {
// 		if i < 2 {
// 			hold[i] = max(hold[i-1], -prices[i]) // 买入或者保持
// 		} else {
// 			hold[i] = max(hold[i-1], sell[i-2]-prices[i]) // 保持或者卖出
// 		}

// 		sell[i] = max(sell[i-1], hold[i-1]+prices[i])
// 	}
// 	return sell[N-1]
// }

//需要冷却期的股票交易 -- 维护 4 种状态的情况
// func maxProfit(prices []int) int {
// 	// 211/211 cases passed (0 ms)
// 	// Your runtime beats 100 % of golang submissions
// 	// Your memory usage beats 100 % of golang submissions (2.5 MB)
// 	N := len(prices)
// 	if N == 0 {
// 		return 0
// 	}
// 	var buy = make([]int, N)
// 	var s1 = make([]int, N)
// 	var sell = make([]int, N)
// 	var s2 = make([]int, N)
// 	s1[0] = -prices[0]
// 	buy[0] = -prices[0]
// 	// sell[0] = 0
// 	// s2[0] = 0
// 	// 从 1 开始,0 已经初始化
// 	for i := 1; i < N; i++ {
// 		buy[i] = s2[i-1] - prices[i] // 买入
// 		s1[i] = max(buy[i-1], s1[i-1])
// 		sell[i] = max(buy[i-1], s1[i-1]) + prices[i] // 卖出
// 		s2[i] = max(s2[i-1], sell[i-1])              // 冷却期
// 	}
// 	return max(sell[N-1], s2[N-1]) // 最后可能是售出或者冷却期
// }

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

