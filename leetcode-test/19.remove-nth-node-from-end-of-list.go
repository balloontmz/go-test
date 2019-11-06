/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (34.62%)
 * Likes:    2271
 * Dislikes: 169
 * Total Accepted:    474.9K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, remove the n-th node from the end of list and return
 * its head.
 *
 * Example:
 *
 *
 * Given linked list: 1->2->3->4->5, and n = 2.
 *
 * After removing the second node from the end, the linked list becomes
 * 1->2->3->5.
 *
 *
 * Note:
 *
 * Given n will always be valid.
 *
 * Follow up:
 *
 * Could you do this in one pass?
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
//删除链表中倒数第n的元素
// 先用一个fast标记n，然后用一个slow 和 fast进行同时遍历
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 208/208 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 71.43 % of golang submissions (2.2 MB)
	var fast = head
	for ; n > 0; n-- {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next // 代表需要删除的是第一个元素
	}
	var slow = head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

// @lc code=end

