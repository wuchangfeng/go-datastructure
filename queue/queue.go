package queue

import (
	"errors"
)

// Queue declaration
type Queue struct {
	// s :=[] int {1,2,3 } 这种格式就是切片，下面的意思是存储任意类型的元素
	data []interface{} // use slice to store queue data
	len  int           // queue size
}

// New()函数返回一个 结构体声明
func New() *Queue {
	return &Queue{}
}

// 需要一个结构体指针作为参数
func (q *Queue) Length() int {
	return q.len
}

// IsEmpty checks if queue is empty
func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

// 在队列末端插入一个元素
func (q *Queue) Push(element interface{}) {
	q.data = append(q.data, element)
	q.len++
}

// 移除一个元素并且返回
// Pop removes an element from the start of queue and return it
func (q *Queue) Pop() (interface{}, error) {
	// 为空的话，返回错误提示
	if q.IsEmpty() {
		return nil, errors.New("Can not pop from an empty queue.")
	}

	// 删除队列前面的第一个元素
	item := q.data[0]
	// 保留第二个到最后一个的元素
	q.data = q.data[1:]
	q.len--
	return item, nil
}

// 只返回第一个元素就好了
func (q *Queue) Peek() interface{} {
	return q.data[0]
}