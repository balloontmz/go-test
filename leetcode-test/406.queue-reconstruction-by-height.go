/*
 * @lc app=leetcode id=406 lang=golang
 *
 * [406] Queue Reconstruction by Height
 *
 * https://leetcode.com/problems/queue-reconstruction-by-height/description/
 *
 * algorithms
 * Medium (59.91%)
 * Likes:    1693
 * Dislikes: 192
 * Total Accepted:    85K
 * Total Submissions: 140.7K
 * Testcase Example:  '[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]'
 *
 * Suppose you have a random list of people standing in a queue. Each person is
 * described by a pair of integers (h, k), where h is the height of the person
 * and k is the number of people in front of this person who have a height
 * greater than or equal to h. Write an algorithm to reconstruct the queue.
 * 
 * Note:
 * The number of people is less than 1,100.
 * 
 * 
 * Example
 * 
 * 
 * Input:
 * [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
 * 
 * Output:
 * [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
 * 
 * 
 * 
 * 
 */
func reconstructQueue(people [][]int) [][]int {
	// 题目描述：一个学生用两个分量 (h, k) 描述，h 表示身高，k 表示排在前面的有 k 个学生的身高比他高或者和他一样高。
	// 为了使插入操作不影响后续的操作，身高较高的学生应该先做插入操作，否则身高较小的学生原先正确插入的第 k 个位置可能会变成第 k+1 个位置。
	// 身高 h 降序、个数 k 值升序，然后将某个学生插入队列的第 k 个位置中。
	// ✔ 37/37 cases passed (44 ms)
	// ✔ Your runtime beats 44.32 % of golang submissions
	// ✔ Your memory usage beats 20 % of golang submissions (8.2 MB)
	if len(people) == 0 || len(people[0]) == 0 {
		return [][]int{}
	}
	people = arraySort(people)
	var res [][]int
	for _, v := range people {
		temp := append([][]int{}, res[v[1]:]...) 
		res = append(res[0:v[1]], v)
		res = append(res, temp...)
	}
	return res
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
	
	if arr1[0] == arr2[0] {
		if arr1[1] < arr2[1] {
			return true
		} 
	} else {
		if arr1[0] > arr2[0] {
			return true
		}
	}
	return false
}