import (
    "math/rand"
)
/*
 * @lc app=leetcode id=382 lang=golang
 *
 * [382] Linked List Random Node
 *
 * https://leetcode.com/problems/linked-list-random-node/description/
 *
 * algorithms
 * Medium (50.34%)
 * Likes:    475
 * Dislikes: 139
 * Total Accepted:    62.1K
 * Total Submissions: 122.4K
 * Testcase Example:  '["Solution","getRandom"]\n[[[1,2,3]],[]]'
 *
 * Given a singly linked list, return a random node's value from the linked
 * list. Each node must have the same probability of being chosen.
 * 
 * Follow up:
 * What if the linked list is extremely large and its length is unknown to you?
 * Could you solve this efficiently without using extra space?
 * 
 * 
 * Example:
 * 
 * // Init a singly linked list [1,2,3].
 * ListNode head = new ListNode(1);
 * head.next = new ListNode(2);
 * head.next.next = new ListNode(3);
 * Solution solution = new Solution(head);
 * 
 * // getRandom() should return either 1, 2, or 3 randomly. Each element should
 * have equal probability of returning.
 * solution.getRandom();
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

//  7/7 cases passed (176 ms)
//  Your runtime beats 5 % of golang submissions
//  Your memory usage beats 100 % of golang submissions (7.4 MB)
type Solution struct {
    cur *ListNode
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    var solution = Solution{
            cur: head,
    }
    return solution
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
    rand.Seed(time.Now().UnixNano())
    var i = 2
    var cur = this.cur
    var res = cur.Val
    //水塘抽样
    for cur.Next != nil {
        cur = cur.Next
        var r = rand.Intn(i)
        if r == 0 {
            res = cur.Val
        }
        i++
    }
    return res
    
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end

