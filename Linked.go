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
	currentNode := ll.head
	for currentNode != nil {
		fmt.Println(currentNode.value)
		currentNode = currentNode.next
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

	//built in function container.list

	mergedList := LinkedList{}
	//traverse list one
	//if l1.next is greater than l2.value
	//addafter

	return mergedList

}

func main() {

	//populate original lisnked lists
	l1 := LinkedList{}
	l1.Insert(1)
	l1.Insert(2)
	l1.Insert(4)

	l2 := LinkedList{}
	l2.Insert(1)
	l2.Insert(3)
	l2.Insert(4)

	//l2.AddAfter(2, 5)
	// /mergedSortedList := LinkedList{}
	//iterate
	//insert between
	//l := l1.GetAt(1)
	//fmt.Println(l.next.value)
	//l2.Traverse()

}

//Input: list1 = [1,2,4], list2 = [1,3,4]
//Output: [1,1,2,3,4,4]
//insert between
