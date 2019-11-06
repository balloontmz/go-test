/*
 * @lc app=leetcode id=628 lang=golang
 *
 * [628] Maximum Product of Three Numbers
 *
 * https://leetcode.com/problems/maximum-product-of-three-numbers/description/
 *
 * algorithms
 * Easy (46.73%)
 * Likes:    804
 * Dislikes: 293
 * Total Accepted:    86.5K
 * Total Submissions: 185K
 * Testcase Example:  '[1,2,3]'
 *
 * Given an integer array, find three numbers whose product is maximum and
 * output the maximum product.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 * Output: 6
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4]
 * Output: 24
 *
 *
 *
 *
 * Note:
 *
 *
 * The length of the given array will be in range [3,10^4] and all elements are
 * in the range [-1000, 1000].
 * Multiplication of any three numbers in the input won't exceed the range of
 * 32-bit signed integer.
 *
 *
 *
 *
 */

// @lc code=start
//计算数组中三个数相乘的最大乘积
//^按位取反
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func maximumProduct(nums []int) int {
	// 83/83 cases passed (36 ms)
	// Your runtime beats 94.87 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.3 MB)
	var max1, max2, max3 = INT_MIN, INT_MIN, INT_MIN
	var min1, min2 = INT_MAX, INT_MAX
	for _, num := range nums {
		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num > max2 {
			max3 = max2
			max2 = num
		} else if num > max3 {
			max3 = num
		}

		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}
	}
	return max(max1*max2*max3, max1*min1*min2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

