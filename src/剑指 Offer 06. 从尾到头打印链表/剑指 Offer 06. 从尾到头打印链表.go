package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	point := head
	var finalList []int
	for {
		finalList = append([]int{point.Val}, finalList...)
		point = point.Next
		if point == nil {
			break
		}
	}
	fmt.Println("finalList", finalList)
	return finalList
}

func main() {
	//var c = ListNode{
	//	Val:  2,
	//	Next: nil,
	//}
	//var b = ListNode{
	//	Val:  3,
	//	Next: &c,
	//}
	//var a = ListNode{
	//	Val:  1,
	//	Next: &b,
	//}

	fmt.Printf("reverse = %d", reversePrint(nil))
}
