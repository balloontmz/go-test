/*
 * @lc app=leetcode id=714 lang=golang
 *
 * [714] Best Time to Buy and Sell Stock with Transaction Fee
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
 *
 * algorithms
 * Medium (51.90%)
 * Likes:    1100
 * Dislikes: 33
 * Total Accepted:    49.7K
 * Total Submissions: 95.4K
 * Testcase Example:  '[1,3,2,8,4,9]\n2'
 *
 * Your are given an array of integers prices, for which the i-th element is
 * the price of a given stock on day i; and a non-negative integer fee
 * representing a transaction fee.
 * You may complete as many transactions as you like, but you need to pay the
 * transaction fee for each transaction.  You may not buy more than 1 share of
 * a stock at a time (ie. you must sell the stock share before you buy again.)
 * Return the maximum profit you can make.
 *
 * Example 1:
 *
 * Input: prices = [1, 3, 2, 8, 4, 9], fee = 2
 * Output: 8
 * Explanation: The maximum profit can be achieved by:
 * Buying at prices[0] = 1Selling at prices[3] = 8Buying at prices[4] =
 * 4Selling at prices[5] = 9The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) =
 * 8.
 *
 *
 *
 * Note:
 * 0 < prices.length .
 * 0 < prices[i] < 50000.
 * 0 .
 *
 */

// @lc code=start
//需要支付费用无冷却期的股票交易 -- 维护 2 个状态数组,性能更优
//应该还可以改造成 o1 空间复杂度,暂时搁置
func maxProfit(prices []int, fee int) int {
	// 44/44 cases passed (88 ms)
	// Your runtime beats 97.14 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.6 MB)
	N := len(prices)
	var hold = make([]int, N) // 持有股票的情况
	var sell = make([]int, N) // 未持有股票的情况
	hold[0] = -prices[0]
	sell[0] = 0
	for i := 1; i < N; i++ {
		hold[i] = max(hold[i-1], sell[i-1]-prices[i])
		sell[i] = max(sell[i-1], hold[i-1]+prices[i]-fee)
	}
	return sell[N-1]
}

//需要支付费用无冷却期的股票交易 -- 维护四个状态数组
// func maxProfit(prices []int, fee int) int {
// 	// 44/44 cases passed (92 ms)
// 	// Your runtime beats 82.86 % of golang submissions
// 	// Your memory usage beats 100 % of golang submissions (7.6 MB)
// 	N := len(prices)
// 	var buy = make([]int, N)
// 	var s1 = make([]int, N)
// 	var sell = make([]int, N)
// 	var s2 = make([]int, N)
// 	buy[0] = -prices[0]
// 	s1[0] = -prices[0]
// 	// sell[0] = 0
// 	// s2[0] = 0
// 	for i := 1; i < N; i++ {
// 		buy[i] = max(s2[i-1], sell[i-1]) - prices[i]
// 		s1[i] = max(buy[i-1], s1[i-1])
// 		sell[i] = max(s1[i-1], buy[i-1]) + prices[i] - fee
// 		s2[i] = max(s2[i-1], sell[i-1])
// 	}
// 	return max(s2[N-1], sell[N-1])
// }

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

