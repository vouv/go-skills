// +build OMIT

package main

import "fmt"

type Node struct {
	char         rune
	children     []*Node
	isEndingChar bool
}

func NewNode(char rune) *Node {
	return &Node{
		char:     char,
		children: make([]*Node, 26),
	}
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode('/'),
	}
}

func (t *Trie) Insert(text string) {
	p := t.root
	for _, c := range text {
		index := c - 'a'
		if p.children[index] == nil {
			newNode := NewNode(c)
			p.children[index] = newNode
			p.isEndingChar = false
		}
		p = p.children[index]
	}
	p.isEndingChar = true
}

func (t *Trie) Find(pattern string) bool {
	// 在Trie树中查找一个字符串
	p := t.root
	for _, c := range pattern {
		index := c - 'a'
		if p.children[index] == nil {
			return false // 不存在pattern
		}
		p = p.children[index]
	}
	if p.isEndingChar == false {
		return false // 不能完全匹配，只是前缀
	} else {
		return true // 找到pattern
	}
}

func main() {
	trie := NewTrie()
	for _, s := range []string{"how", "hi", "her", "hello", "so", "see"} {
		trie.Insert(s)
	}
	fmt.Println("find her:", trie.Find("her"))
}
