import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=646 lang=golang
 *
 * [646] Maximum Length of Pair Chain
 *
 * https://leetcode.com/problems/maximum-length-of-pair-chain/description/
 *
 * algorithms
 * Medium (49.60%)
 * Likes:    647
 * Dislikes: 56
 * Total Accepted:    42.9K
 * Total Submissions: 86.2K
 * Testcase Example:  '[[1,2], [2,3], [3,4]]'
 *
 *
 * You are given n pairs of numbers. In every pair, the first number is always
 * smaller than the second number.
 *
 *
 *
 * Now, we define a pair (c, d) can follow another pair (a, b) if and only if b
 * < c. Chain of pairs can be formed in this fashion.
 *
 *
 *
 * Given a set of pairs, find the length longest chain which can be formed. You
 * needn't use up all the given pairs. You can select pairs in any order.
 *
 *
 *
 * Example 1:
 *
 * Input: [[1,2], [2,3], [3,4]]
 * Output: 2
 * Explanation: The longest chain is [1,2] -> [3,4]
 *
 *
 *
 * Note:
 *
 * The number of given pairs will be in the range [1, 1000].
 *
 *
 */

// @lc code=start
//动态规划解最长整数链--应该能优化排序算法进行优化
func findLongestChain(pairs [][]int) int {
	// 202/202 cases passed (60 ms)
	// Your runtime beats 12.82 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (6.3 MB)
	n := len(pairs)
	var dp = make([]int, n+1)
	temp := Temp{pairs}
	sort.Sort(temp)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if pairs[j][1] < pairs[i][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	fmt.Println("结果为:", dp)
	var ret int
	for i := 1; i <= n; i++ {
		ret = max(ret, dp[i])
	}
	return ret + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Temp struct {
	Data [][]int
}

func (t Temp) Len() int {
	return len(t.Data)
}

// 升序排列
func (t Temp) Less(x, y int) bool {
	return t.Data[x][0] < t.Data[y][0]
}

func (t Temp) Swap(x, y int) {
	t.Data[x], t.Data[y] = t.Data[y], t.Data[x]
}

// @lc code=end

