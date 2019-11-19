/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 *
 * https://leetcode.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (31.29%)
 * Likes:    2387
 * Dislikes: 790
 * Total Accepted:    291.4K
 * Total Submissions: 929.2K
 * Testcase Example:  '[1,2,3]'
 *
 * Implement next permutation, which rearranges numbers into the
 * lexicographically next greater permutation of numbers.
 *
 * If such arrangement is not possible, it must rearrange it as the lowest
 * possible order (ie, sorted in ascending order).
 *
 * The replacement must be in-place and use only constant extra memory.
 *
 * Here are some examples. Inputs are in the left-hand column and its
 * corresponding outputs are in the right-hand column.
 *
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 *
 */

// @lc code=start
//下一个排序
func nextPermutation(nums []int) {
	// 那么是如何得到的呢，我们通过观察原数组可以发现，如果从末尾往前看，数字逐渐变大，到了2时才减小的，然后再从后往前找第一个比2大的数字，是3，那么我们交换2和3，再把此时3后面的所有数字转置一下即可，步骤如下
	//https://www.cnblogs.com/grandyang/p/4428207.html
	// 265/265 cases passed (4 ms)
	// Your runtime beats 75.99 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (3.4 MB)
	var n = len(nums)
	if n <= 1 {
		return
	}
	var idx int
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			idx = i
			break
		}
	}
	var temp int
	if idx != 0 {
		temp = nums[idx-1]
		for i := idx; i < n; i++ {
			if temp >= nums[i] {
				nums[idx-1] = nums[i-1]
				nums[i-1] = temp
				break
			}
			if i == n-1 {
				nums[idx-1] = nums[n-1]
				nums[n-1] = temp
			}
		}
	}
	for i, j := idx, n-1; i < j; i, j = i+1, j-1 {
		temp = nums[i]
		nums[i] = nums[j]
		nums[j] = temp
	}
	return
}

// @lc code=end

