/*
 * @lc app=leetcode id=565 lang=golang
 *
 * [565] Array Nesting
 *
 * https://leetcode.com/problems/array-nesting/description/
 *
 * algorithms
 * Medium (53.88%)
 * Likes:    662
 * Dislikes: 84
 * Total Accepted:    42.6K
 * Total Submissions: 78.9K
 * Testcase Example:  '[5,4,0,3,1,6,2]'
 *
 * A zero-indexed array A of length N contains all integers from 0 to N-1. Find
 * and return the longest length of set S, where S[i] = {A[i], A[A[i]],
 * A[A[A[i]]], ... } subjected to the rule below.
 *
 * Suppose the first element in S starts with the selection of element A[i] of
 * index = i, the next element in S should be A[A[i]], and then A[A[A[i]]]… By
 * that analogy, we stop adding right before a duplicate element occurs in
 * S.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: A = [5,4,0,3,1,6,2]
 * Output: 4
 * Explanation:
 * A[0] = 5, A[1] = 4, A[2] = 0, A[3] = 3, A[4] = 1, A[5] = 6, A[6] = 2.
 *
 * One of the longest S[K]:
 * S[0] = {A[0], A[5], A[6], A[2]} = {5, 6, 2, 0}
 *
 *
 *
 *
 * Note:
 *
 *
 * N is an integer within the range [1, 20,000].
 * The elements of A are all distinct.
 * Each element of A is an integer within the range [0, N-1].
 *
 *
 */

// @lc code=start
//嵌套数组
func arrayNesting(nums []int) int {
	// 856/856 cases passed (8 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (5.9 MB)
	var max int
	for k, _ := range nums {
		var cnt int
		for j := k; nums[j] != -1; {
			//此处代码必须 数组中t是唯一的，也就是说链不存在分叉。但是可能存在多个链。但是多个链互不影响
			cnt++
			t := nums[j]
			nums[j] = -1
			j = t
		}
		max = mF(max, cnt)
	}
	return max
}

func mF(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

