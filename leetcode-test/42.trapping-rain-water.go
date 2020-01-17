/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 *
 * https://leetcode.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (45.46%)
 * Likes:    5338
 * Dislikes: 104
 * Total Accepted:    410.1K
 * Total Submissions: 885.1K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * Given n non-negative integers representing an elevation map where the width
 * of each bar is 1, compute how much water it is able to trap after raining.
 * 
 * 
 * The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
 * In this case, 6 units of rain water (blue section) are being trapped. Thanks
 * Marcos for contributing this image!
 * 
 * Example:
 * 
 * 
 * Input: [0,1,0,2,1,0,1,3,2,1,2,1]
 * Output: 6
 * 
 */

// @lc code=start
//雨水槽的蓄水量 -- 动态规划解法
// 315/315 cases passed (4 ms)
// Your runtime beats 83.27 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.9 MB)
// func trap(height []int) int {
// 	var n = len(height)
// 	var dp = make([]int, n)
// 	var res int
//     if n == 0 {
//         return res
//     }
// 	//寻找当前点左边的最大值放入数组
// 	for i := 1; i < n; i++ {
// 		dp[i] = max(dp[i-1], height[i-1])
// 	}
// 	dp[n-1] = 0// 
// 	for i := n-2; i >= 0; i-- {
// 		dp[i] = min(dp[i], max(height[i+1], dp[i+1]))
// 		if dp[i] > height[i] {
// 			res += dp[i] - height[i]
// 		}
// 	}
// 	return res
// }

//双指针计算雨水槽的蓄水量
// 315/315 cases passed (4 ms)
// Your runtime beats 83.85 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.8 MB)
// func trap(height []int) int {
// 	var n = len(height)
// 	var l = 0
// 	var r = n - 1
// 	var res = 0
// 	for l < r {
// 		var mn = min(height[l], height[r])
// 		if mn == height[l] {
// 			l++
// 			for l < r && mn > height[l] {
// 				res += mn - height[l]
// 				l++
// 			}
// 		} else {
// 			r--
// 			for l < r && mn > height[r] {
// 				res += mn - height[r]
// 				r--
// 			}
// 		}
// 	}
// 	return res
// }

//单调递增栈计算雨水槽的蓄水量 -- dp 是计算每个的高度,双指针也是,单调递增栈则是计算每个横向的宽度
// 315/315 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.9 MB)
func trap(height []int) int {
	var stack []int // 保存高度单调减的索引
	var i = 0
	var n = len(height)
	var res = 0
	for i < n {
		if len(stack) == 0 || height[stack[len(stack)-1]] >= height[i] {
			stack = append(stack, i)
			i++
		} else {
			if len(stack) == 1 {
				stack = []int{}
				continue
			}			
			res += (min(height[i], height[stack[len(stack)-2]]) - height[stack[len(stack)-1]]) * (i - stack[len(stack)-2] - 1)
			stack = stack[:len(stack)-1]
		}
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
	if x > y {
		return y
	}
	return x
}
// @lc code=end

