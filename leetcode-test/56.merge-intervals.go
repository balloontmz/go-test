import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 *
 * https://leetcode.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (37.10%)
 * Likes:    2811
 * Dislikes: 222
 * Total Accepted:    448.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * Given a collection of intervals, merge all overlapping intervals.
 *
 * Example 1:
 *
 *
 * Input: [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into
 * [1,6].
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 *
 * NOTE: input types have been changed on April 15, 2019. Please reset to
 * default code definition to get new method signature.
 *
 */

// @lc code=start

func merge(intervals [][]int) [][]int {
	//拆分成两个数组分别进行排序
	// 169/169 cases passed (8 ms)
	// Your runtime beats 94.14 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (5.9 MB)
	if len(intervals) == 0 {
		return nil
	}
	var starts, ends []int
	var start, end = 0, 0
	for i := 0; i < len(intervals); i++ {
		starts = append(starts, intervals[i][0])
		ends = append(ends, intervals[i][1])
	}
	sort.Ints(starts)
	sort.Ints(ends)
	var res [][]int
	for i := 1; i < len(intervals); i++ {
		if starts[i] <= ends[i-1] {
			end = i
		} else {
			res = append(res, []int{starts[start], ends[end]})
			start = i
			end = i
		}
	}
	res = append(res, []int{starts[start], ends[end]})
	return res
}

// type Inter [][]int

// func merge(intervals [][]int) [][]int {
// 	//先排序,然后逐一遍历
// 	// 169/169 cases passed (8 ms)
// 	// Your runtime beats 94.22 % of golang submissions
// 	// Your memory usage beats 100 % of golang submissions (4.8 MB)
// 	if len(intervals) == 0 {
// 		return nil
// 	}
// 	// 根据第一个数大小升序排列
// 	var inters Inter = intervals
// 	sort.Sort(inters)
// 	var res = [][]int{inters[0]}
// 	for i := 1; i < inters.Len(); i++ {
// 		var n = len(res)
// 		if inters[i][0] <= res[n-1][1] {
// 			res[n-1][1] = max(res[n-1][1], inters[i][1])
// 		} else {
// 			res = append(res, inters[i])
// 		}
// 	}
// 	return res
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// //Len()
// func (t Inter) Len() int {
// 	return len(t)
// }

// //Less(): 成绩将有低到高排序
// func (t Inter) Less(i, j int) bool {
// 	return t[i][0] < t[j][0]
// }

// //Swap()
// func (t Inter) Swap(i, j int) {
// 	t[i], t[j] = t[j], t[i]
// }

// @lc code=end

