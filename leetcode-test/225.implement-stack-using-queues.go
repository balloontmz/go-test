import "fmt"

/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 *
 * https://leetcode.com/problems/implement-stack-using-queues/description/
 *
 * algorithms
 * Easy (41.45%)
 * Likes:    418
 * Dislikes: 485
 * Total Accepted:    148.9K
 * Total Submissions: 358.4K
 * Testcase Example:  '["MyStack","push","push","top","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * Implement the following operations of a stack using queues.
 *
 *
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * empty() -- Return whether the stack is empty.
 *
 *
 * Example:
 *
 *
 * MyStack stack = new MyStack();
 *
 * stack.push(1);
 * stack.push(2);
 * stack.top();   // returns 2
 * stack.pop();   // returns 2
 * stack.empty(); // returns false
 *
 * Notes:
 *
 *
 * You must use only standard operations of a queue -- which means only push to
 * back, peek/pop from front, size, and is empty operations are valid.
 * Depending on your language, queue may not be supported natively. You may
 * simulate a queue by using a list or deque (double-ended queue), as long as
 * you use only standard operations of a queue.
 * You may assume that all operations are valid (for example, no pop or top
 * operations will be called on an empty stack).
 *
 *
 */

// @lc code=start
//用队列实现栈 -- 比较粗糙，暂时搁置
// 16/16 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 50 % of golang submissions (2 MB)
type MyStack struct {
	Queue1, Queue2 *queue
	top            int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		Queue1: newQueue(),
		Queue2: newQueue(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.top = x
	this.Queue1.Push(x)
	return
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	fmt.Println("当前的q1队列为：", this.Queue1.nums)
	for len(this.Queue1.nums) > 1 {
		this.Queue2.Push(this.Queue1.Pop())
	}
	res := this.Queue1.Pop()
	this.Queue2, this.Queue1 = this.Queue1, this.Queue2
	if len(this.Queue1.nums) > 0 {
		this.top = this.Queue1.nums[len(this.Queue1.nums)-1]
	} else {
		this.top = 0
	}

	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.top
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Queue1.Empty() && this.Queue2.Empty()
}

type queue struct {
	nums []int
}

func newQueue() *queue {
	return &queue{
		nums: []int{},
	}
}

func (q *queue) Push(x int) {
	q.nums = append(q.nums, x)
	return
}

func (q *queue) Pop() int {
	if q.Empty() {
		return -1
	}
	res := q.nums[0]
	// fmt.Println(1)
	q.nums = q.nums[1:]
	return res
}

func (q *queue) Empty() bool {
	return len(q.nums) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

