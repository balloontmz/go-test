/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (46.97%)
 * Likes:    1493
 * Dislikes: 136
 * Total Accepted:    372.7K
 * Total Submissions: 793.2K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head.
 *
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 *
 *
 *
 * Example:
 *
 *
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
 *
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
//链表中交换相邻的元素
func swapPairs(head *ListNode) *ListNode {
	// 55/55 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 66.67 % of golang submissions (2.1 MB)
	var node = &ListNode{}
	node.Next = head
	var pre = node
	for pre.Next != nil && pre.Next.Next != nil {
		var l1, l2 = pre.Next, pre.Next.Next
		var next = l2.Next // 此处如果定义为l1的next则会使链表变为环 l1 指向l2  l2指向l1
		l1.Next = next
		l2.Next = l1
		pre.Next = l2

		pre = l1
	}
	return node.Next
}

// @lc code=end

