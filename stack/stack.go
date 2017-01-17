package stack

import (
	"errors"
)

// 堆栈是一个先进后出的数据结构
type Stack struct {
	data []interface{}
	len  int
}

// 初始化一个空堆栈，返回一个指针
func New() *Stack {
	return &Stack{}
}

// 返回堆栈的数组长度
func (s *Stack) Length() int {
	return s.len
}

// 判断是否为空
func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

// Peek return the top-most element in stack, or nil if stack is empty
func (s *Stack) Peek() interface{} {
	return s.data[s.len-1]
}

// Push adds an element to stack
func (s *Stack) Push(element interface{}) {
	s.len++
	s.data = append(s.data, element)
}

// 返回类型的第一个参数是 Interface{} 类型的，即任意类型的
func (s *Stack) Pop() (interface{}, error) {

	if s.IsEmpty() {
		return nil, errors.New("Pop an empty stack.")
	}
	element := s.Peek()
	s.data = s.data[:s.len-1]
	s.len--
	return element, nil
}
