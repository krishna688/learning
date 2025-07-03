package main

import (
	"log"
)

type SinglyLinkedList struct {
	Head *Node
}

func (sll *SinglyLinkedList) Append(val uint64) {

	node := &Node{
		Val: val,
	}

	if sll.Head == nil {
		sll.Head = node
		return
	}

	currNode := sll.Head

	for currNode.Next != nil {
		currNode = currNode.Next
	}

	currNode.Next = node
}

func (sll *SinglyLinkedList) PrintAll() {

	if sll.Head.Next == nil {
		log.Println(sll.Head.Val)
		return
	}

	currNode := sll.Head

	for currNode != nil {
		log.Println(currNode.Val)
		currNode = currNode.Next
	}
}

func (sll *SinglyLinkedList) RemoveAtEnd() {

	curr := sll.Head

	if curr.Next == nil {
		curr = nil
		return
	}

	for curr.Next.Next != nil {
		curr = curr.Next
	}

	curr.Next = nil
}

func (sll *SinglyLinkedList) RemoveAtStart() {

	if sll.Head.Next != nil {
		sll.Head = sll.Head.Next
	} else {
		sll.Head = nil
	}
}

func (sll *SinglyLinkedList) Count() (count uint64) {

	curr := sll.Head
	for curr != nil {
		count++
		curr = curr.Next
	}

	return
}

func (sll *SinglyLinkedList) Read(elementId uint64) uint64 {

	curr := sll.Head
	var currIndex uint64
	for currIndex < elementId {
		currIndex++
		curr = curr.Next
	}

	return curr.Val
}

func (sll *SinglyLinkedList) InsertAtStart(value uint64) {

	node := &Node{
		Val: value,
	}
	node.Next = sll.Head

	sll.Head = node
}

func (sll *SinglyLinkedList) InsertAtIndex(value, index uint64) {

	node := &Node{
		Val: value,
	}

	if index == 0 {
		node.Next = sll.Head
		sll.Head = node
		return
	}

	curr := sll.Head
	for i := 0; i < int(index)-1; i++ {
		if curr == nil {
			panic("invalid linked list")
		}

		curr = curr.Next
	}

	node.Next = curr.Next
	curr.Next = node
}

func (sll *SinglyLinkedList) RemoveAtIndex(index int) {

	cur := sll.Head
	if index == 0 {
		sll.Head = cur.Next
		return
	}

	cur_index := 1
	for cur != nil && cur_index < index {
		cur = cur.Next
		cur_index++
	}

	if cur != nil && cur.Next != nil {
		cur.Next = cur.Next.Next
	}

}
