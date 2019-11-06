/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 *
 * https://leetcode.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (37.31%)
 * Likes:    2071
 * Dislikes: 295
 * Total Accepted:    315.6K
 * Total Submissions: 845.9K
 * Testcase Example:  '[1,2]'
 *
 * Given a singly linked list, determine if it is a palindrome.
 *
 * Example 1:
 *
 *
 * Input: 1->2
 * Output: false
 *
 * Example 2:
 *
 *
 * Input: 1->2->2->1
 * Output: true
 *
 * Follow up:
 * Could you do it in O(n) time and O(1) space?
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
//判断一个链表是否为回文链表
//用到了多步长遍历、反转链表。
func isPalindrome(head *ListNode) bool {
	// 26/26 cases passed (20 ms)
	// Your runtime beats 21.69 % of golang submissions
	// Your memory usage beats 75 % of golang submissions (5.9 MB)
	if head == nil || head.Next == nil {
		return true
	}
	var slow, fast = head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil { // 代表是偶数个元素
		slow = slow.Next
	}
	cut(head, slow)
	slow = reverseList(slow)
	return isE(head, slow)
}

func cut(head, slow *ListNode) {
	for head.Next != slow {
		head = head.Next
	}
	head.Next = nil
}

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

func isE(l1, l2 *ListNode) bool {
	// l1 和 l2 可能不相等，后面的可能比前面的长度大1--当链表长度为奇数的时候
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

// @lc code=end

