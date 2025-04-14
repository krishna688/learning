package main

import (
	"fmt"
	"log"
	"testing"
)

func DefaultLinkedList() *LinkedList {
	dll := &LinkedList{}

	dll.InsertAtEnd(1)
	dll.InsertAtEnd(3)
	dll.InsertAtEnd(5)
	dll.InsertAtEnd(6)

	return dll
}

func TestInsertAtEnd(t *testing.T) {

	dll := DefaultLinkedList()

	dll.PrintForward()

	log.Println("print")
	dll.InsertAtEnd(8)

	dll.PrintForward()
}

func TestInsertAtBegining(t *testing.T) {

	dll := DefaultLinkedList()
	dll.PrintForward()
	log.Println("print")
	dll.InsertAtBegining(8)

	dll.PrintForward()
}

func TestQueue(t *testing.T) {
	q := NewQueue()

	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)

	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())

}
