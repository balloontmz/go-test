/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 *
 * https://leetcode.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (57.72%)
 * Likes:    3010
 * Dislikes: 75
 * Total Accepted:    730.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Reverse a singly linked list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 *
 *
 * Follow up:
 *
 * A linked list can be reversed either iteratively or recursively. Could you
 * implement both?
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
//循环方法的链表反转 -- 头插法，性能更好
func reverseList(head *ListNode) *ListNode {
	// 27/27 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 83.33 % of golang submissions (2.6 MB)
	var temp *ListNode
	var next *ListNode
	for head != nil {
		next = head.Next
		head.Next = temp
		temp = head
		head = next
	}
	return temp

}

//递归方法的链表反转
// func reverseList(head *ListNode) *ListNode {
// 	// 27/27 cases passed (4 ms)
// 	// Your runtime beats 10.81 % of golang submissions
// 	// Your memory usage beats 16.67 % of golang submissions (2.9 MB)
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	var next = head.Next
// 	var newHead = reverseList(next)
// 	next.Next = head
// 	head.Next = nil
// 	return newHead
// }

// @lc code=end

