// +build OMIT

package main

import "fmt"

type MaxQueue struct {
	data  []int
	maxes []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		data:  []int{},
		maxes: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxes) == 0 {
		return -1
	}
	return this.maxes[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.data = append(this.data, value)
	for len(this.maxes) > 0 && value >= this.maxes[len(this.maxes)-1] {
		this.maxes = this.maxes[:len(this.maxes)-1]
	}
	this.maxes = append(this.maxes, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.data) == 0 {
		return -1
	}
	val := this.data[0]
	this.data = this.data[1:]
	if val == this.maxes[0] {
		this.maxes = this.maxes[1:]
	}
	return val
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

func main() {
	q := Constructor()

	q.Push_back(3)
	q.Push_back(2)
	q.Push_back(1)
	fmt.Println("pushed 3-2-1")
	fmt.Println("max", q.Max_value())
	fmt.Println("pop", q.Pop_front())
	fmt.Println("max", q.Max_value())
	fmt.Println("pop", q.Pop_front())
	fmt.Println("max", q.Max_value())
}
