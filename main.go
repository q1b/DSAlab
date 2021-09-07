package main

import (
	"fmt"

	hashtable "github.com/sacarvy/DSAlab/src/hashTable"
)

func main() {
	hashTable := hashtable.Init()
	list := []string{
		"Eric",
		"KENNY",
		"Kyle",
		"TRans",
		"STAN",
		"UIN",
		"Okay",
	}
	for _,v := range list {
		hashTable.Insert(v)
	}
	// HashTable.Delete("STAN")
	fmt.Println(hashTable.Search("STAN"))
}

// github.com/sacarvy/DSAlab/src/linkedList
// myList := lL.LinkedList{}
// node1 := &lL.Node{Data:48}
// node2 := &lL.Node{Data:90}
// node3 := &lL.Node{Data:900}
// myList.Prepend(node1)
// myList.Prepend(node2)
// myList.Prepend(node3)
// myList.DeleteWithValue(90)
// myList.PrintListData()

