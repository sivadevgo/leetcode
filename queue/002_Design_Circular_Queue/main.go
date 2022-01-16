package main

import (
	"fmt"
)

func main() {
	q := Constructor(6)
	fmt.Println(q.EnQueue(6))
	fmt.Println(q.Rear())
	fmt.Println(q.Rear())
	fmt.Println(q.DeQueue())
	fmt.Println(q.EnQueue(5))
	fmt.Println(q.Rear())
	fmt.Println(q.DeQueue())
	fmt.Println(q.Front())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())

}

type MyCircularQueue struct {
	Elements []int
	head     int
	tail     int
	size     int
}

func Constructor(k int) MyCircularQueue {
	s := make([]int, k)
	return MyCircularQueue{
		Elements: s,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Elements[this.tail] = value
	this.tail = (this.tail + 1) % len(this.Elements)
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if !this.IsEmpty() {
		this.head = (this.head + 1) % len(this.Elements)
		this.size--
		return true
	}
	return false
}

func (this *MyCircularQueue) Front() int {
	if this.size <= 0 {
		return -1
	}
	return this.Elements[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.size <= 0 {
		return -1
	}
	return this.Elements[(this.tail+len(this.Elements)-1)%len(this.Elements)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == len(this.Elements)

}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
