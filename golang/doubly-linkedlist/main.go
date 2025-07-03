package main

import (
	"fmt"
)

type Node struct {
	Val      uint64
	Previous *Node
	Next     *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (dll *LinkedList) InsertAtEnd(val uint64) {
	node := &Node{
		Val: val,
	}

	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		return
	}

	node.Previous = dll.Tail
	dll.Tail.Next = node
	dll.Tail = node
}

func (dll *LinkedList) InsertAtBegining(val uint64) {
	node := &Node{
		Val: val,
	}

	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		return
	}

	node.Next = dll.Head
	dll.Head.Previous = node
	dll.Head = node
}

func (dll *LinkedList) PrintForward() {

	curr := dll.Head

	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
