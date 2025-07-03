package main

import (
	"fmt"
)

type Node struct {
	Val   uint64
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func main() {
	fmt.Println("working on Tree")

	tree := &Tree{}

	tree.BSTInsertion(55)
	tree.BSTInsertion(57)
	tree.BSTInsertion(58)
	tree.BSTInsertion(50)

	tree.BSTInsertion(65)
	tree.BSTInsertion(67)
	tree.BSTInsertion(68)
	tree.BSTInsertion(60)

	tree.BSTInsertion(45)
	tree.BSTInsertion(47)
	tree.BSTInsertion(48)
	tree.BSTInsertion(40)

	tree.BSTInsertion(25)
	tree.BSTInsertion(27)
	tree.BSTInsertion(78)
	tree.BSTInsertion(70)
	tree.BSTInsertion(35)

	tree.BSTInsertion(59)
	tree.BSTInsertion(63)

	// log.Println(tree.Search(60))
	// tree.Delete(60)
	// log.Println(tree.Search(60))

	// log.Printf("%+v", tree.Root)

	// tree.Traversal()
	// tree.Delete(67)
	tree.Delete(65)

}

func NewNode(val uint64) *Node {
	return &Node{
		Val: val,
	}
}

func (tree *Tree) BSTInsertion(val uint64) {

	if tree.Root == nil {
		tree.Root = NewNode(val)
		return
	}

	insert(val, tree.Root)
}

func insert(val uint64, node *Node) {

	if val < node.Val {

		if node.Left == nil {
			node.Left = NewNode(val)
		} else {
			insert(val, node.Left)
		}
	}

	if val > node.Val {
		if node.Right == nil {
			node.Right = NewNode(val)
		} else {
			insert(val, node.Right)
		}
	}
}

func (tree *Tree) Search(val uint64) bool {
	return search(val, tree.Root)
}

func search(val uint64, node *Node) bool {

	if node == nil {
		return false
	}

	if val == node.Val {
		return true
	}

	if val < node.Val {
		return search(val, node.Left)
	} else {
		return search(val, node.Right)
	}
}

// func (tree *Tree) Delete(val uint64) {
// 	tree.Root = deleteNode(tree.Root, val)
// }

// func deleteNode(node *Node, val uint64) *Node {
// 	if node == nil {
// 		return nil
// 	}

// 	if val < node.Val {
// 		node.Left = deleteNode(node.Left, val)
// 	} else if val > node.Val {
// 		node.Right = deleteNode(node.Right, val)
// 	} else {
// 		// Node to be deleted found

// 		// Case 1: No child
// 		if node.Left == nil && node.Right == nil {
// 			return nil
// 		}

// 		// Case 2: One child
// 		if node.Left == nil {
// 			return node.Right
// 		}
// 		if node.Right == nil {
// 			return node.Left
// 		}

// 		// Case 3: Two children
// 		// Find in-order successor (smallest in right subtree)
// 		successor := findMin(node.Right)
// 		node.Val = successor.Val
// 		node.Right = deleteNode(node.Right, successor.Val)
// 	}
// 	return node
// }

// func findMin(node *Node) *Node {
// 	current := node
// 	for current.Left != nil {
// 		current = current.Left
// 	}
// 	return current
// }

func (tree *Tree) Traversal() {
	InOrderTraversal(tree.Root)
	fmt.Println()
}

func InOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	InOrderTraversal(node.Left)
	fmt.Print(" ", node.Val)
	InOrderTraversal(node.Right)
}

func (tree *Tree) Delete(val uint64) {
	tree.Traversal()
	deleteNode(val, tree.Root)
	tree.Traversal()

}

func deleteNode(val uint64, node *Node) *Node {

	if node == nil {
		return nil
	}

	if val < node.Val {
		node.Left = deleteNode(val, node.Left)
	} else if val > node.Val {
		node.Right = deleteNode(val, node.Right)
	} else {
		// case 1:  if the node is the leaf node, it'll not have both right and left node
		if node.Left == nil && node.Right == nil {
			return nil
		}

		// case 2: if the node has one child
		if node.Left != nil && node.Right == nil {
			return node.Left
		} else if node.Right != nil && node.Left == nil {
			return node.Right
		}

		// case 3: if the deleted node has two children then replace the deleted node with the successor node
		// successor node is least of the values higher than deleted node
		successor := findMin(node.Right)
		node.Val = successor.Val
		node.Right = deleteNode(successor.Val, node.Right)
	}

	return node
}

func findMin(node *Node) *Node {
	curr := node
	for curr.Left != nil {
		curr = curr.Left
	}

	return curr
}
