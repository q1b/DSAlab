package main

import (
	"fmt"
	"math/rand"
	"time"

	sortfunc "github.com/sacarvy/DSAlab/src/sortfunc"
)

func shuffle(arr *[]int){
    rand.Shuffle(len(*arr),
        func(i, j int) {
            (*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i];
        },
    )
}

func main() {
    // Inisialised to perform suffle function
    rand.Seed(time.Now().UnixNano())
    // Sample array to apply our sorting functions
    sample := []int{3, 4, 5, 2, 1}
    // sortfunc.DirectInsertionSort(&sample)
    shuffle(&sample)
    fmt.Println(sample)
    sortfunc.BubbleSort(&sample);
    fmt.Println(sample)
    // sortfunc.InsertionSort(&sample)
    // shuffle(&sample)
    // fmt.Println(sortfunc.MergeSort(sample));

}

// hashtable "github.com/sacarvy/DSAlab/src/hashTable"
// hashTable := hashtable.Init()
// list := []string{
// 	"Eric",
// 	"KENNY",
// 	"Kyle",
// 	"TRans",
// 	"STAN",
// 	"UIN",
// 	"Okay",
// }
// for _,v := range list {
// 	hashTable.Insert(v)
// }
// hashTable.Delete("STAN")
// fmt.Println(hashTable.Search("STAN"))


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

