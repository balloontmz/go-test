/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 *
 * https://leetcode.com/problems/linked-list-cycle/description/
 *
 * algorithms
 * Easy (37.03%)
 * Likes:    1590
 * Dislikes: 173
 * Total Accepted:    413.9K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * Given a linked list, determine if it has a cycle in it.
 * 
 * To represent a cycle in the given linked list, we use an integer pos which
 * represents the position (0-indexed) in the linked list where tail connects
 * to. If pos is -1, then there is no cycle in the linked list.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: head = [3,2,0,-4], pos = 1
 * Output: true
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * second node.
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: head = [1,2], pos = 0
 * Output: true
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * first node.
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: head = [1], pos = -1
 * Output: false
 * Explanation: There is no cycle in the linked list.
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * Follow up:
 * 
 * Can you solve it using O(1) (i.e. constant) memory?
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// 判断链表是否存在环。
	// 使用双指针，一个指针每次移动一个节点，一个指针每次移动两个指针。如果存在环，那么两个指针一定会相遇
	if head == nil {
		return false
	}
	l1 := head
	l2 := head.Next
	for l1 != nil && l2 != nil && l2.Next != nil {
		if l1 == l2 {
			return true
		}
		l1 = l1.Next
		l2 = l2.Next.Next
	}
	// 如果存在环，不会到此处。
	return false
}

