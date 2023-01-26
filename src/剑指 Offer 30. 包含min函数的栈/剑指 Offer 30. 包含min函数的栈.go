package main

import (
	"fmt"
)

type MinStack struct {
	myList []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.myList = append(this.myList, x)
}

func (this *MinStack) Pop() {
	this.myList = this.myList[:len(this.myList)-1]
}

func (this *MinStack) Top() int {
	return this.myList[len(this.myList)-1]
}

func (this *MinStack) Min() int {
	minNum := this.myList[0]
	for _, x := range this.myList {
		if x < minNum {
			minNum = x
		}
	}
	return minNum
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	param_4 := obj.Min()

	fmt.Printf("Min %d", param_4)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
