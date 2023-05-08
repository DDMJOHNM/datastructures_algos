package main

import (
	"container/heap"
	"fmt"
)

type StringHeap []string

func (sheap StringHeap) Len() int { return len(sheap) }

func (sheap StringHeap) Less(i, j int) bool {
	return sheap[i] > sheap[j]
}

func (sheap StringHeap) Swap(i, j int) {
	sheap[i], sheap[j] = sheap[j], sheap[i]
}

func (sheap *StringHeap) Push(heapstrf interface{}) {
	*sheap = append(*sheap, heapstrf.(string))
}

func (sheap *StringHeap) Pop() interface{} {
	var n int
	var x1 string
	var previous StringHeap = *sheap
	n = len(previous)
	x1 = previous[n-1]
	*sheap = previous[0 : n-1]
	return x1
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	data := []string{"{", "}", "(", ")", "[", "]"}
	var strHeap *StringHeap = &StringHeap{}
	heap.Init(strHeap)

	for _, j := range data {

		heap.Push(strHeap, j)

		if strHeap.Len() > 0 {

			switch j {

			case "{":
				if contains(data, "}") {
					heap.Pop(strHeap)
				}
			case "[":
				if contains(data, "]") {
					heap.Pop(strHeap)
				}

			case "(":
				if contains(data, ")") {
					heap.Pop(strHeap)
				}

			case "}":
				if contains(data, "{") {
					heap.Pop(strHeap)
				}
			case "]":
				if contains(data, "[") {
					heap.Pop(strHeap)
				}

			case ")":
				if contains(data, "(") {
					heap.Pop(strHeap)
				}
			}
		}

	}

	if strHeap.Len() == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
