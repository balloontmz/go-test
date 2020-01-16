/*
 * @lc app=leetcode id=400 lang=golang
 *
 * [400] Nth Digit
 *
 * https://leetcode.com/problems/nth-digit/description/
 *
 * algorithms
 * Medium (30.80%)
 * Likes:    299
 * Dislikes: 941
 * Total Accepted:    53K
 * Total Submissions: 171.7K
 * Testcase Example:  '3'
 *
 * Find the n^th digit of the infinite integer sequence 1, 2, 3, 4, 5, 6, 7, 8,
 * 9, 10, 11, ... 
 * 
 * Note:
 * n is positive and will fit within the range of a 32-bit signed integer (n <
 * 2^31).
 * 
 * 
 * Example 1:
 * 
 * Input:
 * 3
 * 
 * Output:
 * 3
 * 
 * 
 * 
 * Example 2:
 * 
 * Input:
 * 11
 * 
 * Output:
 * 0
 * 
 * Explanation:
 * The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a
 * 0, which is part of the number 10.
 * 
 * 
 */

// @lc code=start
//查找自然数序列的第 n 位的数字
// 70/70 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (1.9 MB)
func findNthDigit(n int) int {
	var length, cnt, start = 1, 9, 1
	for n > length * cnt {
		n -= length * cnt
		length++ // 每次遍历,区间的长度自加 1
		cnt *= 10
		start *= 10
	}
	start += (n-1)/length
	var t = strconv.Itoa(start)
	return int(t[(n-1)%length] - '0')
}
// @lc code=end

