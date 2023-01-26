package main

import "fmt"

type CQueue struct {
	myList []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.myList = append(this.myList, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.myList) == 0 {
		return -1
	}
	firstNum := this.myList[0]
	this.myList = this.myList[1:]
	return firstNum
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {
	obj := Constructor()
	obj.AppendTail(11)
	param_2 := obj.DeleteHead()
	fmt.Printf("%d", param_2)
}
