package main

import "fmt"

type Queue struct {
	head *Node
	tail *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

// insert element at the end
func (q *Queue) EnQueue(val uint64) {

	node := &Node{
		Val: val,
	}

	if q.head == nil {
		q.head = node
		q.tail = node
	}

	q.tail.Next = node
	node.Previous = q.tail
	q.tail = node
}

func (q *Queue) DeQueue() uint64 {

	if q.head == nil {
		return 0
	}

	fmt.Printf("%+v ", q)
	val := q.head.Val
	if q.head.Next != nil {
		q.head = q.head.Next
		q.head.Previous = nil
	} else {
		q.head = nil
	}

	return val
}
