/*
 * @lc app=leetcode id=658 lang=golang
 *
 * [658] Find K Closest Elements
 *
 * https://leetcode.com/problems/find-k-closest-elements/description/
 *
 * algorithms
 * Medium (40.08%)
 * Likes:    1169
 * Dislikes: 229
 * Total Accepted:    92.3K
 * Total Submissions: 229.9K
 * Testcase Example:  '[1,2,3,4,5]\n4\n3'
 *
 * Given a sorted array arr, two integers k and x, find the k closest elements
 * to x in the array. The result should also be sorted in ascending order. If
 * there is a tie, the smaller elements are always preferred.
 * 
 * 
 * Example 1:
 * Input: arr = [1,2,3,4,5], k = 4, x = 3
 * Output: [1,2,3,4]
 * Example 2:
 * Input: arr = [1,2,3,4,5], k = 4, x = -1
 * Output: [1,2,3,4]
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= k <= arr.length
 * 1 <= arr.length <= 10^4
 * Absolute value of elements in the array and x will not exceed 10^4
 * 
 * 
 */

// @lc code=start
// 59/59 cases passed (36 ms)
// Your runtime beats 95.71 % of golang submissions
// Your memory usage beats 100 % of golang submissions (6.4 MB)
func findClosestElements(arr []int, k int, x int) []int {
	//二分搜索 寻找左边节点
	var left = 0
	var right = len(arr)-k
	//二分法
	for left < right {
		var mid = left + (right-left)/2
		if x - arr[mid] > arr[mid+k] - x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	
	return arr[left:left+k]
}
// [0,1,2,2,2,3,6,8,8,9]
// 5
// 9
// @lc code=end

