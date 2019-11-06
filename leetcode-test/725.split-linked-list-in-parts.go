/*
 * @lc app=leetcode id=725 lang=golang
 *
 * [725] Split Linked List in Parts
 *
 * https://leetcode.com/problems/split-linked-list-in-parts/description/
 *
 * algorithms
 * Medium (50.05%)
 * Likes:    433
 * Dislikes: 89
 * Total Accepted:    34.2K
 * Total Submissions: 68.3K
 * Testcase Example:  '[1,2,3,4]\n5'
 *
 * Given a (singly) linked list with head node root, write a function to split
 * the linked list into k consecutive linked list "parts".
 *
 * The length of each part should be as equal as possible: no two parts should
 * have a size differing by more than 1.  This may lead to some parts being
 * null.
 *
 * The parts should be in order of occurrence in the input list, and parts
 * occurring earlier should always have a size greater than or equal parts
 * occurring later.
 *
 * Return a List of ListNode's representing the linked list parts that are
 * formed.
 *
 *
 * Examples
 * 1->2->3->4, k = 5 // 5 equal parts
 * [ [1],
 * [2],
 * [3],
 * [4],
 * null ]
 *
 * Example 1:
 *
 * Input:
 * root = [1, 2, 3], k = 5
 * Output: [[1],[2],[3],[],[]]
 * Explanation:
 * The input and each element of the output are ListNodes, not arrays.
 * For example, the input root has root.val = 1, root.next.val = 2,
 * \root.next.next.val = 3, and root.next.next.next = null.
 * The first element output[0] has output[0].val = 1, output[0].next = null.
 * The last element output[4] is null, but it's string representation as a
 * ListNode is [].
 *
 *
 *
 * Example 2:
 *
 * Input:
 * root = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], k = 3
 * Output: [[1, 2, 3, 4], [5, 6, 7], [8, 9, 10]]
 * Explanation:
 * The input has been split into consecutive parts with size difference at most
 * 1, and earlier parts are a larger size than the later parts.
 *
 *
 *
 * Note:
 * The length of root will be in the range [0, 1000].
 * Each value of a node in the input will be an integer in the range [0, 999].
 * k will be an integer in the range [1, 50].
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//分割链表成k份
func splitListToParts(root *ListNode, k int) []*ListNode {
	// 41/41 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (3.2 MB)
	var N int
	var cur = root
	for cur != nil {
		N++
		cur = cur.Next
	}
	var mod = N % k
	var size = N / k
	cur = root
	var res = make([]*ListNode, k)
	for i := 0; i < k && cur != nil; i++ {
		res[i] = cur
		var curSize int
		if mod > 0 {
			curSize = size + 1
			mod--
		} else {
			curSize = size
		}
		//此处减1是因为一个链表其实是相关两个元素的，所以天然多一个元素。。。
		for j := 0; j < curSize-1; j++ {
			cur = cur.Next
		}
		var next = cur.Next
		cur.Next = nil // 此处代表一个链表已经结束了
		cur = next     //此处指向新链表
	}
	return res
}

// @lc code=end

