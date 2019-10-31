/*
 * @lc app=leetcode id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/description/
 *
 * algorithms
 * Hard (35.01%)
 * Likes:    1428
 * Dislikes: 59
 * Total Accepted:    173K
 * Total Submissions: 491.6K
 * Testcase Example:  '[3,3,5,0,0,3,1,4]'
 *
 * Say you have an array for which the i^th element is the price of a given
 * stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete at most two
 * transactions.
 *
 * Note: You may not engage in multiple transactions at the same time (i.e.,
 * you must sell the stock before you buy again).
 *
 * Example 1:
 *
 *
 * Input: [3,3,5,0,0,3,1,4]
 * Output: 6
 * Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit
 * = 3-0 = 3.
 * Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 =
 * 3.
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4,5]
 * Output: 4
 * Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit
 * = 5-1 = 4.
 * Note that you cannot buy on day 1, buy on day 2 and sell them later, as you
 * are
 * engaging multiple transactions at the same time. You must sell before buying
 * again.
 *
 *
 * Example 3:
 *
 *
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 *
 */

// @lc code=start
//动态规划解只能交易 2 次的股票交易
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func maxProfit(prices []int) int {
	// 200/200 cases passed (4 ms)
	// Your runtime beats 98.17 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (3.1 MB)
	var firstBuy, firstSell, secondBuy, secondSell int
	firstBuy = INT_MIN
	secondBuy = INT_MIN
	for _, price := range prices {
		if firstBuy < -price {
			// 此处一直为负数!!!
			firstBuy = -price // 如果第一次购买价格大于当前遍历价格,改为当前遍历
		}
		if firstSell < price+firstBuy {
			firstSell = price + firstBuy
		}
		// 参数会被第一次的参数初始化!!!仔细理解流程的话
		if secondBuy < firstSell-price {
			secondBuy = firstSell - price
		}
		if secondSell < secondBuy+price {
			secondSell = secondBuy + price
		}
	}
	// if len(prices) <= 2 {
	// 	fmt.Println("计算的结果为:", firstBuy, firstSell, secondBuy, secondSell)
	// }
	return secondSell
}

// @lc code=end

