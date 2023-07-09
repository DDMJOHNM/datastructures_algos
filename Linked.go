package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(val int) {

	n := Node{}
	n.value = val

	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}

	ptr := l.head

	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}

}

func (ll *LinkedList) GetAt(pos int) *Node {
	ptr := ll.head
	if pos < 0 {
		return ptr
	}

	if pos > (ll.len - 1) {
		return nil
	}

	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}

	return ptr

}

func (ll *LinkedList) Traverse() {
	currentNode1 := ll.head
	for currentNode1 != nil {
		fmt.Println(currentNode1.value)
		currentNode1 = currentNode1.next
	}
}

func (ll *LinkedList) AddAfter(nodeProperty int, property int) {

	var node = &Node{}
	node.value = property
	node.next = nil
	var nodeWith *Node
	nodeWith = ll.GetAt(nodeProperty)
	if nodeWith != nil {
		node.next = nodeWith.next
		nodeWith.next = node
	}

}

func mergedSortedLists(l1 *LinkedList, l2 *LinkedList) LinkedList {

	mergedList := LinkedList{}
	currentNode1 := l1.head
	currentNode2 := l2.head

	mergedList.Insert(currentNode1.value)
	mergedList.Insert(currentNode2.value)

	for currentNode1 != nil && currentNode2 != nil && currentNode1.next != nil && currentNode2.next != nil {
		currentNode1 = currentNode1.next
		currentNode2 = currentNode2.next
		mergedList.Insert(currentNode1.value)
		mergedList.Insert(currentNode2.value)
	}

	return mergedList

}

func main() {

	//populate original linked lists
	l1 := &LinkedList{}
	l1.Insert(1)
	l1.Insert(2)
	l1.Insert(4)

	l2 := &LinkedList{}
	l2.Insert(1)
	l2.Insert(3)
	l2.Insert(4)

	result := mergedSortedLists(l1, l2)
	result.Traverse()

}

//Input: list1 = [1,2,4], list2 = [1,3,4]
//Output: [1,1,2,3,4,4]
//insert between
