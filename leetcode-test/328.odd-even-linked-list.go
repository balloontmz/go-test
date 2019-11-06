/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 *
 * https://leetcode.com/problems/odd-even-linked-list/description/
 *
 * algorithms
 * Medium (50.90%)
 * Likes:    988
 * Dislikes: 246
 * Total Accepted:    177.1K
 * Total Submissions: 347.7K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Given a singly linked list, group all odd nodes together followed by the
 * even nodes. Please note here we are talking about the node number and not
 * the value in the nodes.
 *
 * You should try to do it in place. The program should run in O(1) space
 * complexity and O(nodes) time complexity.
 *
 * Example 1:
 *
 *
 * Input: 1->2->3->4->5->NULL
 * Output: 1->3->5->2->4->NULL
 *
 *
 * Example 2:
 *
 *
 * Input: 2->1->3->5->6->4->7->NULL
 * Output: 2->3->6->7->1->5->4->NULL
 *
 *
 * Note:
 *
 *
 * The relative order inside both the even and odd groups should remain as it
 * was in the input.
 * The first node is considered odd, the second node even and so on ...
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
//链表按序号奇偶重拍 -- 没想到！！！
func oddEvenList(head *ListNode) *ListNode {
	// 71/71 cases passed (4 ms)
	// Your runtime beats 88.72 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (3.3 MB)
	if head == nil || head.Next == nil {
		return head
	}
	var odd, even = head, head.Next
	var evenHead = even // 用于保存偶数链表的头部
	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		odd = odd.Next
		even.Next = even.Next.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

// @lc code=end

