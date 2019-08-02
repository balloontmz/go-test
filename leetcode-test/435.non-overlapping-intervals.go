/*
 * @lc app=leetcode id=435 lang=golang
 *
 * [435] Non-overlapping Intervals
 *
 * https://leetcode.com/problems/non-overlapping-intervals/description/
 *
 * algorithms
 * Medium (41.58%)
 * Likes:    529
 * Dislikes: 22
 * Total Accepted:    42.5K
 * Total Submissions: 101.6K
 * Testcase Example:  '[[1,2]]'
 *
 * Given a collection of intervals, find the minimum number of intervals you
 * need to remove to make the rest of the intervals non-overlapping.
 * 
 * Note:
 * 
 * 
 * You may assume the interval's end point is always bigger than its start
 * point.
 * Intervals like [1,2] and [2,3] have borders "touching" but they don't
 * overlap each other.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: [ [1,2], [2,3], [3,4], [1,3] ]
 * 
 * Output: 1
 * 
 * Explanation: [1,3] can be removed and the rest of intervals are
 * non-overlapping.
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [ [1,2], [1,2], [1,2] ]
 * 
 * Output: 2
 * 
 * Explanation: You need to remove two [1,2] to make the rest of intervals
 * non-overlapping.
 * 
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: [ [1,2], [2,3] ]
 * 
 * Output: 0
 * 
 * Explanation: You don't need to remove any of the intervals since they're
 * already non-overlapping.
 * 
 * 
 * NOTE: input types have been changed on April 15, 2019. Please reset to
 * default code definition to get new method signature.
 * 
 */
func eraseOverlapIntervals(intervals [][]int) int {
	// ✔ 18/18 cases passed (4 ms)
	// ✔ Your runtime beats 100 % of golang submissions
	// ✔ Your memory usage beats 100 % of golang submissions (4.1 MB)
	if len(intervals) ==0 {
		return 0
	}
	intervals = arraySort(intervals)
	cnt, end := 1, intervals[0][1]
	for i := 1; i < len(intervals); i++ {
        if (intervals[i][0] < end) {
            continue;
        }
        end = intervals[i][1];
        cnt++;
    }
    return len(intervals) - cnt;
}

//按指定规则对nums进行排序(注：此firstIndex从0开始)
func arraySort(nums [][]int) [][]int {
    //检查
    if len(nums) <= 1 {
        return nums
    }

    //排序
    mIntArray := &IntArray{nums}
    sort.Sort(mIntArray)
    return mIntArray.mArr
}
 
type IntArray struct {
    mArr       [][]int
}
 
//IntArray实现sort.Interface接口
func (arr *IntArray) Len() int {
    return len(arr.mArr)
}
 
func (arr *IntArray) Swap(i, j int) {
    arr.mArr[i], arr.mArr[j] = arr.mArr[j], arr.mArr[i]
}
 
func (arr *IntArray) Less(i, j int) bool {
    arr1 := arr.mArr[i]
	arr2 := arr.mArr[j]
	
	if arr1[1] < arr2[1] {
		return true
	} 
	return false
}