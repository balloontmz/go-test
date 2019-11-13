/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 *
 * https://leetcode.com/problems/implement-queue-using-stacks/description/
 *
 * algorithms
 * Easy (45.57%)
 * Likes:    709
 * Dislikes: 121
 * Total Accepted:    174.1K
 * Total Submissions: 380.9K
 * Testcase Example:  '["MyQueue","push","push","peek","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * Implement the following operations of a queue using stacks.
 *
 *
 * push(x) -- Push element x to the back of queue.
 * pop() -- Removes the element from in front of queue.
 * peek() -- Get the front element.
 * empty() -- Return whether the queue is empty.
 *
 *
 * Example:
 *
 *
 * MyQueue queue = new MyQueue();
 *
 * queue.push(1);
 * queue.push(2);
 * queue.peek();  // returns 1
 * queue.pop();   // returns 1
 * queue.empty(); // returns false
 *
 * Notes:
 *
 *
 * You must use only standard operations of a stack -- which means only push to
 * top, peek/pop from top, size, and is empty operations are valid.
 * Depending on your language, stack may not be supported natively. You may
 * simulate a stack by using a list or deque (double-ended queue), as long as
 * you use only standard operations of a stack.
 * You may assume that all operations are valid (for example, no pop or peek
 * operations will be called on an empty queue).
 *
 *
 */

// @lc code=start
//用两个栈实现队列。
// 17/17 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 25 % of golang submissions (1.9 MB)
type MyQueue struct {
	Stack1, Stack2 *stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		Stack1: newStack(),
		Stack2: newStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Stack1.push(x)
	return
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Stack2.isEmpty() {
		//优化: 栈a中留一个元素供pop,可以少一次操作
		for this.Stack1.len() > 1 {
			this.Stack2.push(this.Stack1.pop())
		}
		return this.Stack1.pop()
	}
	return this.Stack2.pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Stack2.isEmpty() {
		if this.Stack1.isEmpty() {
			return -1
		}
		return this.Stack1.nums[0]
	}
	return this.Stack2.nums[this.Stack2.len()-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.Stack1.isEmpty() && this.Stack2.isEmpty()
}

type stack struct {
	nums []int
}

func newStack() *stack {
	return &stack{
		nums: []int{},
	}
}

func (s *stack) push(n int) {
	s.nums = append(s.nums, n)
}

func (s *stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	res := s.nums[len(s.nums)-1]
	// fmt.Println("test")
	s.nums = (s.nums)[:len(s.nums)-1] // 此处总是报数组越界，是语法问题还是？ --如果加上面那行代码就好了
	return res
}

func (s *stack) len() int {
	return len(s.nums)
}
func (s *stack) isEmpty() bool {
	return s.len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

