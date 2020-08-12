// +build OMIT

package main

import "fmt"

type LRUCache struct {
	head  *Node
	tail  *Node
	mp    map[int]*Node
	cap   int
	count int
}

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func (c *LRUCache) Add(val int) {
	if node, ok := c.mp[val]; ok {
		if node.Prev != nil {
			node.Prev = node.Next
		}
		if node.Next != nil {
			node.Next = node.Prev
		}
		node.Prev = nil
		node.Next = c.head
		c.head = node
		return
	}
	node := &Node{
		Val:  val,
		Prev: nil,
		Next: c.head,
	}
	c.mp[val] = node
	if c.head != nil {
		c.head.Prev = node
	}
	c.head = node
	c.count++

	if c.count > c.cap {
		// delete last
		c.tail = c.tail.Prev
		if c.tail != nil {
			c.tail.Next = nil
		}
		c.count--
	}
}

func (c *LRUCache) GetLeast() int {
	first := c.head
	if first != nil {
		return first.Val
	}
	panic("get from empty cache")
}

func NewLRUCache(cap int) *LRUCache {
	var head = &Node{} // 哨兵
	return &LRUCache{
		head: head,
		tail: head,
		mp:   make(map[int]*Node),
		cap:  cap,
	}
}

func main() {
	lru := NewLRUCache(4)
	lru.Add(1)
	lru.Add(2)
	lru.Add(3)
	lru.Add(4)
	fmt.Println(lru.GetLeast())
	lru.Add(5)
	fmt.Println(lru.GetLeast())
}
