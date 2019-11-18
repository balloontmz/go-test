/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 *
 * https://leetcode.com/problems/hamming-distance/description/
 *
 * algorithms
 * Easy (70.77%)
 * Likes:    1437
 * Dislikes: 138
 * Total Accepted:    267.6K
 * Total Submissions: 378K
 * Testcase Example:  '1\n4'
 *
 * The Hamming distance between two integers is the number of positions at
 * which the corresponding bits are different.
 *
 * Given two integers x and y, calculate the Hamming distance.
 *
 * Note:
 * 0 ≤ x, y < 2^31.
 *
 *
 * Example:
 *
 * Input: x = 1, y = 4
 *
 * Output: 2
 *
 * Explanation:
 * 1   (0 0 0 1)
 * 4   (0 1 0 0)
 * ⁠      ↑   ↑
 *
 * The above arrows point to positions where the corresponding bits are
 * different.
 *
 *
 */

// @lc code=start
//统计两个数不同的位
func hammingDistance(x int, y int) int {
	// 149/149 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (2 MB)
	var cnt int
	var z = x ^ y // 异或，不相等的位才是1
	for z != 0 {
		if z&1 == 1 {
			cnt++
		}
		z = z >> 1
	}
	return cnt
}

// @lc code=end

