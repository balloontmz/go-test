/*
 * @lc app=leetcode id=668 lang=golang
 *
 * [668] Kth Smallest Number in Multiplication Table
 *
 * https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/description/
 *
 * algorithms
 * Hard (45.00%)
 * Likes:    473
 * Dislikes: 18
 * Total Accepted:    19.2K
 * Total Submissions: 42.6K
 * Testcase Example:  '3\n3\n5'
 *
 * 
 * Nearly every one have used the Multiplication Table. But could you find out
 * the k-th smallest number quickly from the multiplication table?
 * 
 * 
 * 
 * Given the height m and the length n of a m * n Multiplication Table, and a
 * positive integer k, you need to return the k-th smallest number in this
 * table.
 * 
 * 
 * Example 1:
 * 
 * Input: m = 3, n = 3, k = 5
 * Output: 
 * Explanation: 
 * The Multiplication Table:
 * 1    2    3
 * 2    4    6
 * 3    6    9
 * 
 * The 5-th smallest number is 3 (1, 2, 2, 3, 3).
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: m = 2, n = 3, k = 6
 * Output: 
 * Explanation: 
 * The Multiplication Table:
 * 1    2    3
 * 2    4    6
 * 
 * The 6-th smallest number is 6 (1, 2, 2, 3, 4, 6).
 * 
 * 
 * 
 * 
 * Note:
 * 
 * The m and n will be in the range [1, 30000].
 * The k will be in the range [1, m * n]
 * 
 * 
 */

// @lc code=start
//二分搜索查询 第 k 个最小元素 table 中
// 69/69 cases passed (28 ms)
// Your runtime beats 66.67 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2 MB)
func findKthNumber(m int, n int, k int) int {
	//二分搜索法
	var left = 1
	var right = m * n
	for left < right {
		var mid = left + (right - left) / 2 // 查询 table 中小于 mid 的数字的个数
		var cnt = 0
		for i := 1; i <= m; i++ {
			if mid >= i * n  {
				cnt += n
			} else {
				cnt += mid / i
			}
		}
		if cnt < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}
// @lc code=end

