package main

import (
	"fmt"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	curr := head
	var pre Node
	newHead := &pre
	var tmpMap map[*Node]*Node
	for curr != nil {

		var tmpNode Node
		// 如果不在map里面
		if mapNode, ok := tmpMap[curr]; !ok {
			var newNode Node
			newNode.Val = curr.Val
			tmpNode = newNode
			tmpMap[curr] = &newNode
		} else {
			tmpNode = *mapNode
		}

		newHead.Next = &tmpNode
		newHead = newHead.Next

		curr = curr.Next
	}
	return pre.Next
}

func main() {
	node1 := Node{
		Val:    1,
		Next:   nil,
		Random: nil,
	}

	node10 := Node{
		Val:    10,
		Next:   &node1,
		Random: nil,
	}

	node11 := Node{
		Val:    11,
		Next:   &node10,
		Random: nil,
	}

	node13 := Node{
		Val:    13,
		Next:   &node11,
		Random: nil,
	}

	node7 := Node{
		Val:    7,
		Next:   &node13,
		Random: nil,
	}

	node7.Random = nil
	node13.Random = &node7
	node11.Random = &node1
	node10.Random = &node11
	node1.Random = &node7

	aaa := copyRandomList(&node7)
	fmt.Printf("%s", aaa)
}
