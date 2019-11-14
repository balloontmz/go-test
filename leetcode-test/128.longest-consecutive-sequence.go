/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 *
 * https://leetcode.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Hard (43.07%)
 * Likes:    2339
 * Dislikes: 133
 * Total Accepted:    242.3K
 * Total Submissions: 561.8K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * Given an unsorted array of integers, find the length of the longest
 * consecutive elements sequence.
 *
 * Your algorithm should run in O(n) complexity.
 *
 * Example:
 *
 *
 * Input: [100, 4, 200, 1, 3, 2]
 * Output: 4
 * Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
 * Therefore its length is 4.
 *
 *
 */

// @lc code=start
//计算最长连续序列的长度
//核心函数为forword，前提是数组元素不重复，也就是map的值初始为0或者1，然后forword计算并赋值大于1的数
//重复元素也只取一个
// 68/68 cases passed (12 ms)
// Your runtime beats 21.62 % of golang submissions
// Your memory usage beats 14.29 % of golang submissions (4.3 MB)
var m = make(map[int]int, 0)

func longestConsecutive(nums []int) int {
	defer func() {
		m = make(map[int]int, 0)
	}()
	for _, v := range nums {
		if m[v] == 0 {
			m[v]++
		}
	}
	for k, _ := range m {
		forword(k)
	}
	return max(m)
}

func forword(val int) int {
	if m[val] == 0 {
		return 0
	}
	cnt := m[val]
	if cnt > 1 { // 代表已经计算过，直接返回
		return cnt
	}
	cnt = forword(val+1) + 1
	m[val] = cnt
	return cnt
}

func max(m map[int]int) int {
	var mV int
	for _, v := range m {
		if v > mV {
			mV = v
		}
	}
	return mV
}

// @lc code=end

