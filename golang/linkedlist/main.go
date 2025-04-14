package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// you can extend it with other operations like inserting at a specific position, deleting at the end, or reversing the list
func main() {

	linkedList := &SinglyLinkedList{}
	linkedList.Append(2)
	linkedList.Append(4)
	linkedList.Append(5)
	linkedList.Append(6)

	linkedList.RemoveAtIndex(5)
	fmt.Println(Jsonify(linkedList))
	// linkedList.PrintAll()
}

func Jsonify(val interface{}) string {

	by, err := json.Marshal(val)
	if err != nil {
		log.Println("Error %s", err)
		return ""
	}

	return string(by)
}
