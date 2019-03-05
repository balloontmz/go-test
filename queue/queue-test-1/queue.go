package queue

// Iten 队列元素的类型
type Item interface {
}

// ItemQueue 队列的基本元素
type ItemQueue struct {
	items []Item
}

// ItemQueuer 队列接口
type ItemQueuer interface {
	New() ItemQueue
	Enqueue(t Item)
	Dequeue() *Item
	IsEmpty() bool
	Size() int
}

// New 创建一个新队列
func (s *ItemQueue) New() *ItemQueue {
	s.items = []Item{}
	return s
}

// Enqueue 为队列添加一个元素
func (s *ItemQueue) Enqueue(t Item) {
	s.items = append(s.items, t)
	return
}

// Dequeue 在队列中删除一个元素
func (s *ItemQueue) Dequeue() *Item {
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	return &item
}

// IsEmpty 判断队列是否为空
func (s *ItemQueue) IsEmpty() bool {
	return len(s.items) == 0
}

// Size 判断队列的大小
func (s *ItemQueue) Size() int {
	return len(s.items)
}
