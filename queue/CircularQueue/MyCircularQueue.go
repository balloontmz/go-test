package circularqueue

import (
	"errors"
	"log"
	"time"
)

// MyCircularQueue
type MyCircularQueue struct {
	element  []string
	Capacity int
	front    int
	rear     int
}

// Constructor 在这里初始化你的数据结构，设置队列的大小。
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		element:  make([]string, k),
		Capacity: k,
		front:    0,
		rear:     0,
	}
}

// EnQueue ** 将一个元素插入循环队列，操作成功，返回 true  */
func (q *MyCircularQueue) EnQueue(value string) bool {
	if (q.rear+1)%q.Capacity == q.front%q.Capacity { // 这种方案采用的是为了区分开满时和空时状态减一个容量的做法
		err := errors.New("队列已满！")
		log.Print("队列已满，当前时间戳为：", time.Now().Unix(), err.Error())
		return false
	}
	q.element[q.rear] = value
	q.rear = (q.rear + 1) % q.Capacity
	return true

}

// DeQueue ** 从循环队列中删除一个元素。成功返回 该元素 */
func (q *MyCircularQueue) DeQueue() string {
	if q.rear == q.front {
		err := errors.New("队列为空，没有数据出队列")
		log.Print("队列为空，当前时间戳为：", time.Now().Unix(), err.Error())
		return ""
	}
	res := q.element[q.front]
	q.element[q.front] = ""
	q.front = (q.front + 1) % q.Capacity
	return res
}

// Front ** 获取队列的第一个元素。 */
func (q *MyCircularQueue) Front() string {
	if q.rear == q.front {
		err := errors.New("队列为空，没有数据出队列")
		log.Print("队列为空，当前时间戳为：", time.Now().Unix(), err.Error())
		return ""
	}
	res := q.element[q.front]
	return res
}

// Rear ** 获取队列的最后一个元素。 */
func (q *MyCircularQueue) Rear() string {
	if q.rear == q.front {
		err := errors.New("队列为空，没有数据出队列")
		log.Print("队列为空，当前时间戳为：", time.Now().Unix(), err.Error())
		return ""
	}
	res := q.element[(q.rear-1+q.Capacity)%q.Capacity]
	return res
}

// IsEmpty ** 检查队列是否为空。 */
func (q *MyCircularQueue) IsEmpty() bool {
	if q.rear == q.front {
		return true
	}
	return false
}

// IsFull ** 检查队列是否为满。 */
func (q *MyCircularQueue) IsFull() bool {
	if (q.rear+1)%q.Capacity == q.front%q.Capacity { // 这种方案采用的是为了区分开满时和空时状态减一个容量的做法
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
