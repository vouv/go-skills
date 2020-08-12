// +build OMIT

package main

import (
	"fmt"
)

type MinStack struct {
	data []int
	mins []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: []int{},
		mins: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.mins) == 0 || this.mins[len(this.mins)-1] >= x {
		this.mins = append(this.mins, x)
	}
}

func (this *MinStack) Pop() {
	val := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	if val <= this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) Min() int {
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println("expect -3 and get", minStack.Min())
	minStack.Pop()
	fmt.Println("expect 0 and get", minStack.Top())
	fmt.Println("expect -2 and get", minStack.Min())
}
