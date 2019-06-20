/*
 * @lc app=leetcode id=633 lang=golang
 *
 * [633] Sum of Square Numbers
 *
 * https://leetcode.com/problems/sum-of-square-numbers/description/
 *
 * algorithms
 * Easy (32.72%)
 * Likes:    339
 * Dislikes: 227
 * Total Accepted:    45K
 * Total Submissions: 137.7K
 * Testcase Example:  '5'
 *
 * Given a non-negative integer c, your task is to decide whether there're two
 * integers a and b such that a^2 + b^2 = c.
 * 
 * Example 1:
 * 
 * 
 * Input: 5
 * Output: True
 * Explanation: 1 * 1 + 2 * 2 = 5
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 3
 * Output: False
 * 
 * 
 * 
 * 
 */
 import (
	 "math"
 )

func judgeSquareSum(c int) bool {
	i := 0
	j := int(math.Sqrt(float64(c)))
	for i <= j { // j 等于 j 的情况。
		sum := i*i + j*j
		if sum == c {
			return true
		} else if sum > c {
			j--
		} else {
			i++
		}
	}
	return false
}

