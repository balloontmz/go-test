import (
	"strconv"
	"math"
)
/*
 * @lc app=leetcode id=483 lang=golang
 *
 * [483] Smallest Good Base
 *
 * https://leetcode.com/problems/smallest-good-base/description/
 *
 * algorithms
 * Hard (34.73%)
 * Likes:    125
 * Dislikes: 296
 * Total Accepted:    11.8K
 * Total Submissions: 33.5K
 * Testcase Example:  '"13"'
 *
 * For an integer n, we call k>=2 a good base of n, if all digits of n base k
 * are 1.
 * 
 * Now given a string representing n, you should return the smallest good base
 * of n in string format.
 * 
 * Example 1:
 * 
 * 
 * Input: "13"
 * Output: "3"
 * Explanation: 13 base 3 is 111.
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "4681"
 * Output: "8"
 * Explanation: 4681 base 8 is 11111.
 * 
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "1000000000000000000"
 * Output: "999999999999999999"
 * Explanation: 1000000000000000000 base 999999999999999999 is 11.
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * The range of n is [3, 10^18].
 * The string representing n is always valid and will not have leading
 * zeros.
 * 
 * 
 * 
 * 
 */

// @lc code=start
// 106/106 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (1.9 MB)
func smallestGoodBase(n string) string {
	snum, err := strconv.ParseInt(n, 10, 64)
	num := uint64(snum)
	if err != nil {
		fmt.Print("解析字符串成整数失败", err)
		return ""
	}

	//注意整数溢出...应该是边界条件的问题!!!
	for m := math.Log2(float64(num+1)); m >= 2.0; m = m - 1.0 {
		var left uint64 = 2
		var right = uint64(math.Ceil(math.Pow(float64(num), 1.0/(m-1.0)))) + 1
		// fmt.Print("当前遍历的 指数为:", m, "right 为:", right, "\n")
		
		for left < right {
			var sum uint64 = 0
			var mid = left + (right-left) / 2
			for j := 0.0; j < m; j++ {
				sum = sum * mid + 1
			}
			if sum == num {
				return strconv.FormatInt(int64(mid), 10)
			}
			if sum < num {
				left = mid+1
			} else {
				right = mid
			}
		}
	}
	return strconv.FormatInt(int64(num-1), 10)

}
// @lc code=end

