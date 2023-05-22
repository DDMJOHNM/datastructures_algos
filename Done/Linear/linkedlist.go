package main

import (
	"fmt"
)


type Node struct{
	property int
	nextNode *Node	

}

type LinkedList struct{
	headNode *Node
}


func (linkedList *LinkedList) AddToHead(property int){
	var node = Node{}
	node.property = property

	if node.nextNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

func (linkedList *LinkedList) IterateList(){
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode{
		fmt.Println(node.property)
	}
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node{
	
	var node *Node 
	var nodeWith *Node
	
	for node = linkedList.headNode; node != nil; node = node.nextNode{
		if node.property == property {
			nodeWith = node
			break;
		}
	}
	
	return nodeWith
} 

func (linkedList *LinkedList) AddAfter (nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
	  node.nextNode = nodeWith.nextNode
	  nodeWith.nextNode = node		
	}

}


func main(){

	//add to head method
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(2)
	linkedList.AddToHead(3)
	linkedList.AddToHead(4)
	//fmt.Println(linkedList.headNode.property)

	//iterate list method 
	//linkedList.IterateList();
	//fmt.Println(linkedList.NodeWithValue(3));
	linkedList.AddAfter(4,5)
	linkedList.IterateList()	
	
}