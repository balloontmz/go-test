/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (47.34%)
 * Likes:    2887
 * Dislikes: 136
 * Total Accepted:    547.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * Say you have an array for which the i^th element is the price of a given
 * stock on day i.
 * 
 * If you were only permitted to complete at most one transaction (i.e., buy
 * one and sell one share of the stock), design an algorithm to find the
 * maximum profit.
 * 
 * Note that you cannot sell a stock before you buy one.
 * 
 * Example 1:
 * 
 * 
 * Input: [7,1,5,3,6,4]
 * Output: 5
 * Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit
 * = 6-1 = 5.
 * Not 7-1 = 6, as selling price needs to be larger than buying price.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 * 
 * 
 */
func maxProfit(prices []int) int {
	// ✔ 200/200 cases passed (4 ms)
	// ✔ Your runtime beats 97.06 % of golang submissions
	// ✔ Your memory usage beats 73.58 % of golang submissions (3.1 MB)
    if len(prices) == 0 {
		return 0
	}
	var min int = prices[0]
	var max int = 0
	for index := 0; index < len(prices); index++ {
		if min > prices[index] {
			min = prices[index]
		} else {
			max = Max(max, prices[index]-min)
		}
	}
	return max
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
