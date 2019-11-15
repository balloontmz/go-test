/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 *
 * https://leetcode.com/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (51.50%)
 * Likes:    3096
 * Dislikes: 361
 * Total Accepted:    235.4K
 * Total Submissions: 456.1K
 * Testcase Example:  '[1,3,4,2,2]'
 *
 * Given an array nums containing n + 1 integers where each integer is between
 * 1 and n (inclusive), prove that at least one duplicate number must exist.
 * Assume that there is only one duplicate number, find the duplicate one.
 *
 * Example 1:
 *
 *
 * Input: [1,3,4,2,2]
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [3,1,3,4,2]
 * Output: 3
 *
 * Note:
 *
 *
 * You must not modify the array (assume the array is read only).
 * You must use only constant, O(1) extra space.
 * Your runtime complexity should be less than O(n^2).
 * There is only one duplicate number in the array, but it could be repeated
 * more than once.
 *
 *
 */

// @lc code=start
func findDuplicate(nums []int) int {
	// 二分法
	// 53/53 cases passed (4 ms)
	// Your runtime beats 98.11 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (3.8 MB)
	// var low, high = 1, len(nums) - 1
	// for low <= high {
	// 	var cnt int
	// 	var mid = low + (high-low)/2
	// 	for i := 0; i < len(nums); i++ {
	// 		if nums[i] <= mid {
	// 			cnt++
	// 		}
	// 	}
	// 	if cnt > mid {
	// 		high = mid - 1
	// 	} else {
	// 		low = mid + 1
	// 	}
	// }
	// return low

	//双指针法
	//查找有环链表的入口
	// 53/53 cases passed (4 ms)
	// Your runtime beats 98.11 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (3.8 MB)
	var slow, fast = nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// @lc code=end

