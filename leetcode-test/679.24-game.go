/*
 * @lc app=leetcode id=679 lang=golang
 *
 * [679] 24 Game
 *
 * https://leetcode.com/problems/24-game/description/
 *
 * algorithms
 * Hard (45.64%)
 * Likes:    632
 * Dislikes: 133
 * Total Accepted:    36K
 * Total Submissions: 78.6K
 * Testcase Example:  '[4,1,8,7]'
 *
 * 
 * You have 4 cards each containing a number from 1 to 9.  You need to judge
 * whether they could operated through *, /, +, -, (, ) to get the value of
 * 24.
 * 
 * 
 * Example 1:
 * 
 * Input: [4, 1, 8, 7]
 * Output: True
 * Explanation: (8-4) * (7-1) = 24
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: [1, 2, 1, 2]
 * Output: False
 * 
 * 
 * 
 * Note:
 * 
 * The division operator / represents real division, not integer division.  For
 * example, 4 / (1 - 2/3) = 12.
 * Every operation done is between two numbers.  In particular, we cannot use -
 * as a unary operator.  For example, with [1, 1, 1, 1] as input, the
 * expression -1 - 1 - 1 - 1 is not allowed.
 * You cannot concatenate numbers together.  For example, if the input is [1,
 * 2, 1, 2], we cannot write this as 12 + 12.
 * 
 * 
 * 
 */

// @lc code=start
// 24 点,采用递归回溯,深度优先遍历
// 70/70 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (3 MB)
var exp = 0.000001

func judgePoint24(nums []int) bool {
	var res bool
	var numsF = make([]float64, 0)
	
	for _, num := range nums {
		numsF = append(numsF, float64(num))
	}
	helper(numsF, &res)
	return res
}

func helper(nums []float64, res *bool)  {
	if *res {
		return
	}
	if len(nums) == 1 {
		if (nums[0] - float64(24) > 0 && nums[0] - float64(24) < exp) || (float64(24) - nums[0] > 0 && float64(24) - nums[0] < exp) || float64(24) - nums[0] == 0 {
			*res = true
		}
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			var p = nums[i]
			var q = nums[j]
			var t = []float64{
				p+q,
				q-p,
				p-q,
				q*p,
			}
			if p > 0 {
				t = append(t, q/p)
			}

			if q > 0 {
				t = append(t, p/q)
			}

			//去除 q 和 p 对应的值
			nums = append(nums[:i], append(nums[i+1:j], nums[j+1:]...)...)
			for _, v := range t {
				nums = append(nums, v)
				helper(nums, res)
				nums = nums[:len(nums)-1] // 子函数不会改变数组,这边直接除去最后一位,并不会改变顺序
			}
			//位置不能发生变化 -- 否则会导致无法正确计算结果. 可能覆盖已有元素,或者重复遍历或空遍历元素!!!
			nums = append(nums[:i], append([]float64{p}, nums[i:]...)...)
			nums = append(nums[:j], append([]float64{q}, nums[j:]...)...)
		}
	}
}

// [3,3,8,8]

// @lc code=end

