package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{2, 1, 3, 5, 4}
	heap.Init(h)
	heap.Push(h, 100)
	heap.Push(h, 50)
	fmt.Println(heap.Pop(h))

	fmt.Println(h)

	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}

// type HeapInterface interface {
//     sort.Interface
//     Push(x interface{}) // add x as element Len()
//     Pop() interface{}   // remove and return element Len() - 1.
// }
