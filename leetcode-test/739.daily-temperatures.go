/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 *
 * https://leetcode.com/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (61.05%)
 * Likes:    1752
 * Dislikes: 54
 * Total Accepted:    101.2K
 * Total Submissions: 165.4K
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 *
 * Given a list of daily temperatures T, return a list such that, for each day
 * in the input, tells you how many days you would have to wait until a warmer
 * temperature.  If there is no future day for which this is possible, put 0
 * instead.
 *
 * For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76,
 * 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].
 *
 *
 * Note:
 * The length of temperatures will be in the range [1, 30000].
 * Each temperature will be an integer in the range [30, 100].
 *
 */

// @lc code=start
//数组中元素与下一个比它大的元素之间的距离
//最近的下一个温暖的天气
//采用栈存储遍历过的元素，如果当前遍历比当前栈顶元素大，那么当前遍历就是栈顶元素的最近温暖天气
//同时栈内的温度从栈顶往下必定是升序的
// 37/37 cases passed (900 ms)
// Your runtime beats 59.77 % of golang submissions
// Your memory usage beats 100 % of golang submissions (8.5 MB)
func dailyTemperatures(T []int) []int {
	var n = len(T)
	var stack []int
	var res = make([]int, n) // 用来保存遍历的索引
	for i := 0; i < n; i++ {
		for len(stack) != 0 && T[i] > T[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

// @lc code=end

