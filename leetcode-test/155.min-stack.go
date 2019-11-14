/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (39.47%)
 * Likes:    2275
 * Dislikes: 244
 * Total Accepted:    370.7K
 * Total Submissions: 936.2K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 *
 *
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * getMin() -- Retrieve the minimum element in the stack.
 *
 *
 *
 *
 * Example:
 *
 *
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> Returns -3.
 * minStack.pop();
 * minStack.top();      --> Returns 0.
 * minStack.getMin();   --> Returns -2.
 *
 *
 *
 *
*/

// @lc code=start
//最小值栈
//维护两个栈，一个保存元素，另一个保存当前最小值得索引
// 18/18 cases passed (20 ms)
// Your runtime beats 79.16 % of golang submissions
// Your memory usage beats 100 % of golang submissions (8.5 MB)
type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.min) == 0 {
		this.min = append(this.min, 0) // 第一次压入肯定是第一个元素
	} else {
		// 如果传入的值最小
		if this.data[this.min[len(this.min)-1]] > x {
			this.min = append(this.min, len(this.data)-1)
		}
	}
	return
}

func (this *MinStack) Pop() {
	length := len(this.data)
	minIndex := length - 1
	if minIndex == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.data = this.data[:length-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.data[this.min[len(this.min)-1]]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

