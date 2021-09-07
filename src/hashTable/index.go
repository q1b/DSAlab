package hashtable

import (
	"fmt"
)

const ArraySize = 7
// HashTable structure 

type HashTable struct{
	array [ArraySize]*bucket	
}
// bucket is a linked list
type bucket struct{
	head *bucketNode
}
// bucketNode strucutre
type bucketNode struct{
	key string
	next *bucketNode
}
// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key);
	h.array[index].insert(key)
}
// Search
func (h *HashTable) Search(key string) bool {
	index := hash(key);
	return h.array[index].search(key)
}
// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key);
	h.array[index].delete(key)
}

// insert 
func (b *bucket) insert(k string){
	if !b.search(k) {
		newNode := &bucketNode{key:k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Randy is on exsit")
	}
}
// search
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
// delete
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode = previousNode.next
		}
	}
}
// hash
func hash(key string) int {
	sum := 0 
	for _,v := range key{
		sum+=int(v)
	}
	return sum % ArraySize
}
// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}
