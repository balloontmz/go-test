/*
 * @lc app=leetcode id=378 lang=golang
 *
 * [378] Kth Smallest Element in a Sorted Matrix
 *
 * https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/description/
 *
 * algorithms
 * Medium (51.19%)
 * Likes:    1578
 * Dislikes: 96
 * Total Accepted:    137.8K
 * Total Submissions: 268.5K
 * Testcase Example:  '[[1,5,9],[10,11,13],[12,13,15]]\n8'
 *
 * Given a n x n matrix where each of the rows and columns are sorted in
 * ascending order, find the kth smallest element in the matrix.
 *
 *
 * Note that it is the kth smallest element in the sorted order, not the kth
 * distinct element.
 *
 *
 * Example:
 *
 * matrix = [
 * ⁠  [ 1,  5,  9],
 * ⁠  [10, 11, 13],
 * ⁠  [12, 13, 15]
 * ],
 * k = 8,
 *
 * return 13.
 *
 *
 *
 * Note:
 * You may assume k is always valid, 1 ≤ k ≤ n^2.
 */

// @lc code=start
//查找二维有序矩阵中排第k位的元素
func kthSmallest(matrix [][]int, k int) int {
	//-- 二分法
	// 85/85 cases passed (40 ms)
	// Your runtime beats 41.84 % of golang submissions
	// Your memory usage beats 50 % of golang submissions (6.4 MB)
	var m, n = len(matrix), len(matrix[0])
	var lo, hi = matrix[0][0], matrix[m-1][n-1]
	for lo <= hi {
		var mid = lo + (hi-lo)/2 // 防止溢出
		var cnt = 0
		for i := 0; i < m; i++ {
			for j := 0; j < n && matrix[i][j] <= mid; j++ {
				cnt++
			}
		}
		if cnt < k { // 这个条件不可互换，有个等号的情况
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

// @lc code=end

