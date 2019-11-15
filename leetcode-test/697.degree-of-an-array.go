/*
 * @lc app=leetcode id=697 lang=golang
 *
 * [697] Degree of an Array
 *
 * https://leetcode.com/problems/degree-of-an-array/description/
 *
 * algorithms
 * Easy (51.97%)
 * Likes:    625
 * Dislikes: 556
 * Total Accepted:    63.7K
 * Total Submissions: 122.5K
 * Testcase Example:  '[1,2,2,3,1]'
 *
 * Given a non-empty array of non-negative integers nums, the degree of this
 * array is defined as the maximum frequency of any one of its elements.
 * Your task is to find the smallest possible length of a (contiguous) subarray
 * of nums, that has the same degree as nums.
 *
 * Example 1:
 *
 * Input: [1, 2, 2, 3, 1]
 * Output: 2
 * Explanation:
 * The input array has a degree of 2 because both elements 1 and 2 appear
 * twice.
 * Of the subarrays that have the same degree:
 * [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
 * The shortest length is 2. So return 2.
 *
 *
 *
 *
 * Example 2:
 *
 * Input: [1,2,2,3,1,4,2]
 * Output: 6
 *
 *
 *
 * Note:
 * nums.length will be between 1 and 50,000.
 * nums[i] will be an integer between 0 and 49,999.
 *
 */

// @lc code=start
//数组的度
func findShortestSubArray(nums []int) int {
	// 89/89 cases passed (36 ms)
	// Your runtime beats 21.74 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.9 MB)
	var numCnt = make(map[int]int, 0)
	var numStartIdx = make(map[int]int, 0)
	var numEndIdx = make(map[int]int, 0)
	//赋值中间变量map
	for k, num := range nums {
		numCnt[num] += 1
		numEndIdx[num] = k
		if _, ok := numStartIdx[num]; !ok {
			numStartIdx[num] = k
		}
	}
	var maxCnt int
	//找出最大值
	for _, cnt := range numCnt {
		maxCnt = max(maxCnt, cnt)
	}

	var res = len(nums)
	for _, num := range nums {
		if numCnt[num] != maxCnt {
			continue
		}
		res = min(res, numEndIdx[num]-numStartIdx[num]+1) // +1操作
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

