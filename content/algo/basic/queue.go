// +build OMIT

package main

import "fmt"

type Queue interface {
	// 入队
	Enqueue(e ...int)
	// 出队
	Dequeue() int

	// 获取队首元素，不出队
	Peek() int

	// 队列是否为空你
	Empty() bool

	// 返回队列大小
	Size() int

	// 清空队列
	Clear()
}

// 使用数组切片实现的队列
type ArrayQueue []int

// 入队操作
func (s *ArrayQueue) Enqueue(e ...int) {
	*s = append(*s, e...)
}

// 出队
func (s *ArrayQueue) Dequeue() int {
	if s.Empty() {
		return -1
	}
	first := (*s)[0]

	// 出队操作
	*s = (*s)[1:]
	return first
}

// 获取队首元素，不出队
func (s *ArrayQueue) Peek() int {
	if s.Empty() {
		return -1
	}
	return (*s)[0]
}

// 队列是否为空
func (s *ArrayQueue) Empty() bool {
	return s.Size() == 0
}

// 返回队列大小
func (s *ArrayQueue) Size() int {
	return len(*s)
}

// 清空队列
func (s *ArrayQueue) Clear() {
	*s = ArrayQueue{}
}

func NewArrayQueue() Queue {
	return &ArrayQueue{}
}

type node struct {
	Value int
	Next  *node
}

// 使用链表实现的队列
// 这里使用带有头尾节点的链表
type LinkedQueue struct {
	Head *node
	Tail *node
	size int
}

func (s *LinkedQueue) Enqueue(e ...int) {
	for _, v := range e {
		s.Tail.Next = &node{
			Value: v,
			Next:  nil,
		}
		s.Tail = s.Tail.Next
		s.size++
	}
}

func (s *LinkedQueue) Dequeue() int {
	if s.Empty() {
		return -1
	}
	first := s.Head.Next.Value

	// 出队
	s.Head = s.Head.Next
	s.size--
	return first
}

func (s *LinkedQueue) Peek() int {
	if s.Empty() {
		return -1
	}
	return s.Head.Next.Value
}

func (s *LinkedQueue) Empty() bool {
	return s.Head == s.Tail
}

func (s *LinkedQueue) Size() int {
	return s.size
}

func (s *LinkedQueue) Clear() {
	s.Head.Next = nil
	s.Tail = s.Head
	s.size = 0
}

func NewLinkedQueue() Queue {
	head := &node{}
	return &LinkedQueue{
		Head: head,
		Tail: head,
		size: 0,
	}
}

func main() {
	q := NewArrayQueue()
	//q := NewLinkedQueue()

	fmt.Println("Empty:", q.Empty())
	fmt.Println("Size:", q.Size())

	q.Enqueue(1, 2, 3, 4, 5, 6)
	fmt.Println("Enqueue 1,2,3,4,5,6")
	fmt.Println("Size:", q.Size())
	fmt.Println("Peek:", q.Peek())

	fmt.Println("Dequeue:", q.Dequeue(), q.Dequeue(), q.Dequeue())
	fmt.Println("Size:", q.Size())
	fmt.Println("Empty:", q.Empty())

	fmt.Println("Dequeue:", q.Dequeue(), q.Dequeue(), q.Dequeue())
	fmt.Println("Size:", q.Size())

	fmt.Println("Empty:", q.Empty())

	fmt.Println("Enqueue:1,2,3")
	q.Enqueue(1, 2, 3)

	fmt.Println("Dequeue:", q.Dequeue(), q.Dequeue())

	fmt.Println("Clear")
	q.Clear()
	fmt.Println("Empty:", q.Empty(), "Size:", q.Size())

}
