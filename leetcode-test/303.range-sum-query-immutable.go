/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 *
 * https://leetcode.com/problems/range-sum-query-immutable/description/
 *
 * algorithms
 * Easy (40.14%)
 * Likes:    575
 * Dislikes: 862
 * Total Accepted:    160.4K
 * Total Submissions: 397.5K
 * Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'
 *
 * Given an integer array nums, find the sum of the elements between indices i
 * and j (i ≤ j), inclusive.
 *
 * Example:
 *
 * Given nums = [-2, 0, 3, -5, 2, -1]
 *
 * sumRange(0, 2) -> 1
 * sumRange(2, 5) -> -1
 * sumRange(0, 5) -> -3
 *
 *
 *
 * Note:
 *
 * You may assume that the array does not change.
 * There are many calls to sumRange function.
 *
 *
 */

// @lc code=start
type NumArray struct {
	Dynamic []int
}

//动态规划解法,相当于提前把所有结果都计算出来!!!对于计算量要求少的觉得不划算!!!
func Constructor(nums []int) NumArray {
	// 16/16 cases passed (36 ms)
	// Your runtime beats 81.46 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (10.9 MB)
	var res = NumArray{
		make([]int, len(nums)),
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res.Dynamic[i] = nums[i]
		} else {
			res.Dynamic[i] = res.Dynamic[i-1] + nums[i]
		}
	}
	return res
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.Dynamic[j]
	}
	return this.Dynamic[j] - this.Dynamic[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end

