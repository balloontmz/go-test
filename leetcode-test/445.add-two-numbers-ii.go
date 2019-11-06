/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
 *
 * https://leetcode.com/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (51.76%)
 * Likes:    928
 * Dislikes: 123
 * Total Accepted:    114.9K
 * Total Submissions: 221.8K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The most significant digit comes first and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 * Follow up:
 * What if you cannot modify the input lists? In other words, reversing the
 * lists is not allowed.
 *
 *
 *
 * Example:
 *
 * Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 8 -> 0 -> 7
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
//采用链表反转以及删除链表的方法解两个链表的和--此处耗时较多，调试不够熟练。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 1563/1563 cases passed (12 ms)
	// Your runtime beats 69.23 % of golang submissions
	// Your memory usage beats 100 % of golang submissions (5 MB)

	// 涉及到 206 链表反转 以及 19 删除链表的倒数第n个节点
	var res = &ListNode{}
	t1, t2, tail := getTail(l1, l2)
	head, tailNum := addTwoNum(t1, t2, res)
	tailR := reverseList(tail)

	tempHead := tailR
	//TODO: 垃圾硬编码！！！
	for tempHead != nil || tailNum != 0 {
		if tempHead == nil { // tailNum 不为0的情况下，只有初次
			tempHead := &ListNode{}
			tempHead.Val = tailNum
			tailR = tempHead
			tailNum = 0
		} else {
			if tailNum != 0 {
				tempHead.Val += tailNum
				if tempHead.Val >= 10 {
					tailNum = tempHead.Val / 10
					tempHead.Val %= 10

					// 只有这种情况需要判断如果没有下一个则加一个下一个节点
					if tempHead.Next == nil {
						temp := &ListNode{}
						tempHead.Next = temp
					}
				} else {
					tailNum = 0
				}
			}
			tempHead = tempHead.Next
		}
	}

	head.Next = tailR

	res = reverseList(res.Next)

	return res
}

func addTwoNum(l1 *ListNode, l2 *ListNode, res *ListNode) (*ListNode, int) {
	if l1 == nil {
		return res, 0
	}
	r, tailNum := addTwoNum(l1.Next, l2.Next, res)
	var head = &ListNode{
		(l1.Val + l2.Val + tailNum) % 10,
		nil,
	}
	r.Next = head
	num := (l1.Val + l2.Val + tailNum) / 10
	return r.Next, num
}

func getTail(l1 *ListNode, l2 *ListNode) (*ListNode, *ListNode, *ListNode) {
	var t1, t2 = l1, l2
	for t1 != nil && t2 != nil {
		t1 = t1.Next
		t2 = t2.Next
	}
	if t1 != nil {
		r1, r2 := removeNFromEnd(l1, t1)
		return l2, r1, r2
	}
	if t2 != nil {
		r1, r2 := removeNFromEnd(l2, t2)
		return l1, r1, r2
	}
	return l1, l2, nil
}

func removeNFromEnd(head *ListNode, fast *ListNode) (*ListNode, *ListNode) {
	var slow = head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	culList := slow.Next
	slow.Next = nil
	return culList, head
}

func reverseList(head *ListNode) *ListNode {
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

// @lc code=end

