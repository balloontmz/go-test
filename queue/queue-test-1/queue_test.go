package queue

import "testing"

var q ItemQueue

func initQueue() *ItemQueue {
	if q.items == nil {
		q = ItemQueue{}
		q.New()
	}

	return &q
}

func TestItemQueue_Enqueue(t *testing.T) {
	q := initQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if size := q.Size(); size != 3 {
		t.Errorf("错误的统计数，正确的统计数应该是3，但是取得了 %d 数", size)
	}
}

func TestItemQueue_Dequeue(t *testing.T) {
	q.Dequeue()

	if size := q.Size(); size != 2 {
		t.Errorf("错误的统计数，正确的统计数应该是2，但是取得了 %d 数", size)
	}

	q.Dequeue()
	q.Dequeue()
	if !q.IsEmpty() {
		t.Errorf("该队列此处应该为空")
	}
}
