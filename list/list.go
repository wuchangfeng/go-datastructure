/* List is a simple yet often used data structure.

l := list.New()   // create a new empty list
l.PushBack(1, 2, 3) // append one or multiple value to list
l.Length()           // return the length of list
l.Index(index)    // get the value at index
l.Find(value)     // get the index of the first appearence of value
l.Lpop()          // remove and return the first value in a list
l.Rpop()          // remove and return the last value in a list
*/

package list

import (
	"fmt"
)

// List 中节点的声明，Value 是任意类型，并且带有前后指针，指针类型为 node
type Node struct {
	Value      interface{} // List can store any type of value
	next, prev *Node       // double linked list
}

// List 数据结构的声明
type List struct {
	root   Node // sentinal node
	length int  // how many nodea are in the list
}

// 返回一个空链表
func New() *List {
	l := &List{}
	l.length = 0

	// 空链表的头结点的前驱和后继节点都是头结点本省
	l.root.next = &l.root
	l.root.prev = &l.root

	return l
}

// 判空函数，为什么 &l.root 这种形式的呢？
func (l *List) IsEmpty() bool {
	return l.root.next == &l.root
}

// 计算出 l 的长度
func (l *List) Length() int {
	return l.length
}

// 在 list 的前面添加元素，也就是传统的头插入方法
func (l *List) PushFront(elements ...interface{}) {
	for _, element := range elements {
		n := &Node{Value: element}
		n.next = l.root.next
		// 注意这里的 & 为什么会存在
		n.prev = &l.root
		l.root.next.prev = n
		l.root.next = n
		l.length++
	}
}

// 尾部插入
func (l *List) PushBack(elements ...interface{}) {
	for _, element := range elements {
		n := &Node{Value: element}
		n.next = &l.root     // since n is the last element, its next should be the head
		n.prev = l.root.prev // n's prev should be the tail
		l.root.prev.next = n // tail's next should be n
		l.root.prev = n      // head's prev should be n
		l.length++
	}
}

// Find the element in list, return the index if found, otherwise return -1
func (l *List) Find(element interface{}) int {
	index := 0
	p := l.root.next
	for p != &l.root && p.Value != element {
		p = p.next
		index++
	}

	if p != &l.root {
		return index
	}

	return -1
}

func (l *List) indexFrontwise(index int) *Node {
	pos := 0
	p := l.root.next

	for p != &l.root && pos < index {
		p = p.next
		pos++
	}

	if p == &l.root {
		return nil
	}
	return p
}

func (l *List) indexBackwise(index int) *Node {
	pos := 1
	p := l.root.prev

	for p != &l.root && pos < index {
		p = p.prev
		pos++
	}

	if p == &l.root {
		return nil
	}
	return p
}

func (l *List) index(index int) *Node {
	var n *Node
	if index >= 0 {
		n = l.indexFrontwise(index)
	} else {
		n = l.indexBackwise(-index)
	}
	return n
}

// Index returns the element at index, if element is not valid, return error
// Support negative index, like -1, -2 etc, it will count backwise though.
func (l *List) Index(index int) (interface{}, error) {
	n := l.index(index)
	if n == nil {
		return nil, fmt.Errorf("Index %d is not valid.", index)
	}
	return n.Value, nil
}

// 移除链表中的元素
func (l *List) remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	// 将 n 这个节点孤立出来
	n.next = nil
	n.prev = nil
	// 长度减去
	l.length--
}

// 移除并返回 list 头元素
func (l *List) Lpop() interface{} {
	// 长度为空的话，直接返回
	if l.length == 0 {
		return nil
	}
	// 获取头结点后面的那个元素
	n := l.root.next
	// 移除
	l.remove(n)
	// 返回 value
	return n.Value
}

// 移除并返回最后一个元素呢
func (l *List) Rpop() interface{} {
	if l.length == 0 {
		return nil
	}
	// 这个是一个环形的链表
	n := l.root.prev
	l.remove(n)
	return n.Value
}

// Given a index of list, return the normal index between 0 and len-1
func (l *List) normalIndex(index int) int {
	if index > l.length-1 {
		index = l.length - 1
	}

	if index < -l.length {
		index = 0
	}

	index = (l.length + index) % l.length
	return index
}

// Range returns a slice containing elements in range
// end can be negative, like -1, -2 etc
// if *start* is large than *end*, return empty slice `[]'
// if the end is large than the actual end, it will be treated like the last element
func (l *List) Range(start, end int) []interface{} {
	// When start or end exceeds list length
	start = l.normalIndex(start)
	end = l.normalIndex(end)
	res := []interface{}{}
	if start > end {
		return res
	}

	sNode := l.index(start)
	eNode := l.index(end)
	for n := sNode; n != eNode; {
		res = append(res, n.Value)
		n = n.next
	}
	res = append(res, eNode.Value)
	return res
}
