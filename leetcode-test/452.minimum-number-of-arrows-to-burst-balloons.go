/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 *
 * https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/
 *
 * algorithms
 * Medium (46.47%)
 * Likes:    527
 * Dislikes: 26
 * Total Accepted:    41.6K
 * Total Submissions: 88.7K
 * Testcase Example:  '[[10,16],[2,8],[1,6],[7,12]]'
 *
 * There are a number of spherical balloons spread in two-dimensional space.
 * For each balloon, provided input is the start and end coordinates of the
 * horizontal diameter. Since it's horizontal, y-coordinates don't matter and
 * hence the x-coordinates of start and end of the diameter suffice. Start is
 * always smaller than end. There will be at most 10^4 balloons.
 * 
 * An arrow can be shot up exactly vertically from different points along the
 * x-axis. A balloon with xstart and xend bursts by an arrow shot at x if
 * xstart ≤ x ≤ xend. There is no limit to the number of arrows that can be
 * shot. An arrow once shot keeps travelling up infinitely. The problem is to
 * find the minimum number of arrows that must be shot to burst all balloons.
 * 
 * Example:
 * 
 * 
 * Input:
 * [[10,16], [2,8], [1,6], [7,12]]
 * 
 * Output:
 * 2
 * 
 * Explanation:
 * One way is to shoot one arrow for example at x = 6 (bursting the balloons
 * [2,8] and [1,6]) and another arrow at x = 11 (bursting the other two
 * balloons).
 * 
 * 
 * 
 * 
 */
func findMinArrowShots(points [][]int) int {
	// ✔ 43/43 cases passed (68 ms)
	// ✔ Your runtime beats 98.36 % of golang submissions
	// ✔ Your memory usage beats 50 % of golang submissions (6.9 MB)
    if len(points) ==0 {
		return 0
	}
	points = arraySort(points)
	cnt, end := 1, points[0][1]
	for i := 1; i < len(points); i++ {
        if (points[i][0] <= end) {
            continue;
        }
        end = points[i][1];
        cnt++;
    }
    return cnt;
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
