package main

import (
	"container/heap"
	"fmt"
)

//Heaps - data structure based on a heap property
//used in selection, graph and kway merge alogrithms
//operations such as finding, insertion keys changes and deleting all perfromed on heaps
//heap order value is stored at each node is greater than equal to its children
//order descending max heap otherwise minimum heap
//not sorted but partial sorted
//heap structure, heap sorting alogrithm
//a php example https://www.sitepoint.com/data-structures-3/
//https://www.w3resource.com/php-exercises/searching-and-sorting-algorithm/searching-and-sorting-algorithm-exercise-2.php

type IntegerHeap []int

func (iheap IntegerHeap) Len() int { return len(iheap) }

func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] > iheap[j]
}

func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

//interheap method pushes the item
func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

//interheap method pops the item from the heap

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minumum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("d% \n", heap.Pop(intHeap))
	}

}
