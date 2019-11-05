/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 *
 * https://leetcode.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (30.09%)
 * Likes:    1356
 * Dislikes: 460
 * Total Accepted:    280.9K
 * Total Submissions: 933.6K
 * Testcase Example:  '10'
 *
 * Count the number of prime numbers less than a non-negative number, n.
 *
 * Example:
 *
 *
 * Input: 10
 * Output: 4
 * Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
 *
 *
 */

// @lc code=start
//生成素数序列 -- 埃拉托斯特尼筛
func countPrimes(n int) int {
	// 20/20 cases passed (8 ms)
	// Your runtime beats 84.03 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (5 MB)
	var notPrime = make([]bool, n+1)
	var count = 0
	for i := 2; i < n; i++ {
		if notPrime[i] {
			continue
		}
		count++
		// 从 i * i 开始，因为如果 k < i，那么 k * i 在之前就已经被去除过了
		for j := i * i; j < n; j += i {
			notPrime[j] = true
		}
	}
	return count
}

// @lc code=end

