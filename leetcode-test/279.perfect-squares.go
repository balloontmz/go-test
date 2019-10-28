/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 *
 * https://leetcode.com/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (42.69%)
 * Likes:    1627
 * Dislikes: 136
 * Total Accepted:    201.1K
 * Total Submissions: 469.9K
 * Testcase Example:  '12'
 *
 * Given a positive integer n, find the least number of perfect square numbers
 * (for example, 1, 4, 9, 16, ...) which sum to n.
 *
 * Example 1:
 *
 *
 * Input: n = 12
 * Output: 3
 * Explanation: 12 = 4 + 4 + 4.
 *
 * Example 2:
 *
 *
 * Input: n = 13
 * Output: 2
 * Explanation: 13 = 4 + 9.
 */

// 采用动态规划实现平方数拆解,性能和广度优先遍历差不多
func numSquares(n int) int {
	// 588/588 cases passed (16 ms)
	// Your runtime beats 88.12 % of golang submissions
	// Your memory usage beats 33.33 % of golang submissions (5.8 MB)
	squares := generateSquares(n)
	var dp = make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := n // 先赋予一个最大值
		for _, s := range squares {
			if s > i {
				break
			}
			min = minf(min, dp[i-s]+1) // 函数名和变量名重复
		}
		dp[i] = min
	}
	return dp[n]
}

func minf(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// 采用先进先出队列实现广度优先遍历
// func numSquares(n int) int {
// 	// ✔ 588/588 cases passed (4 ms)
// 	// ✔ Your runtime beats 97.63 % of golang submissions
// 	// ✔ Your memory usage beats 33.33 % of golang submissions (6 MB)
// 	var squares []int
// 	squares = generateSquares(n)
// 	var queue []int
// 	var marked = make([]bool, n+1)
// 	queue = append(queue, n)

// 	marked[n] = true
// 	var level = 0
// 	// 广度优先遍历 -- 当某次遍历成功时，该路径为最短路径
// 	for len(queue) != 0 {
// 		var size = len(queue)
// 		level ++
// 		for i := size; i > 0; i-- {
// 			var cur = queue[0] // 后入先出
// 			queue = queue[1:]
// 			for _, v := range squares {
// 				var next = cur - v
// 				if next < 0 {
// 					break
// 				}
// 				if next == 0 {
// 					return level
// 				}
// 				if marked[next] == true {
// 					continue
// 				}
// 				marked[next] = true
// 				queue = append(queue, next)
// 			}
// 		}
// 	}
// 	// 整个遍历未成功
// 	return n

// }

// 此处平方数生成的算法可根据数学归纳法推算出来
func generateSquares(n int) []int {
	var squares []int
	var square = 1
	var diff = 3
	for square <= n {
		squares = append(squares, square)
		square += diff
		diff += 2
	}
	return squares
}

