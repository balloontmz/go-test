/*
 * @lc app=leetcode id=503 lang=golang
 *
 * [503] Next Greater Element II
 *
 * https://leetcode.com/problems/next-greater-element-ii/description/
 *
 * algorithms
 * Medium (52.62%)
 * Likes:    879
 * Dislikes: 52
 * Total Accepted:    64K
 * Total Submissions: 121.4K
 * Testcase Example:  '[1,2,1]'
 *
 *
 * Given a circular array (the next element of the last element is the first
 * element of the array), print the Next Greater Number for every element. The
 * Next Greater Number of a number x is the first greater number to its
 * traversing-order next in the array, which means you could search circularly
 * to find its next greater number. If it doesn't exist, output -1 for this
 * number.
 *
 *
 * Example 1:
 *
 * Input: [1,2,1]
 * Output: [2,-1,2]
 * Explanation: The first 1's next greater number is 2; The number 2 can't find
 * next greater number; The second 1's next greater number needs to search
 * circularly, which is also 2.
 *
 *
 *
 * Note:
 * The length of given array won't exceed 10000.
 *
 */

// @lc code=start
//循环数组中比当前元素大的下一个元素
//和普通数组类似，重复一次元素数组
func nextGreaterElements(nums []int) []int {
	// 224/224 cases passed (332 ms)
	// Your runtime beats 62.96 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (8.2 MB)
	var n = len(nums)
	var stack []int
	var res = make([]int, n) // 用来保存遍历的索引
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	for i := 0; i < n*2; i++ {
		num := nums[i%n]
		for len(stack) != 0 && num > nums[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		if i < n {
			stack = append(stack, i)
		}
	}
	return res
}

// @lc code=end

