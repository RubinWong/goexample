package main

type elem struct {
	Priority int
	Value    interface{}
}

// type heap.Interface interface {
// 	sort.Interface
// 	Push(x any) // add x as element Len()
// 	Pop() any   // remove and return element Len() - 1.
// }

type PriorityQueue []elem

func (p *PriorityQueue) Len() int {
	return len(*p)
}

func (p *PriorityQueue) Less(i, j int) bool {
	return (*p)[i].Priority > (*p)[j].Priority
}

func (p *PriorityQueue) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x.Value
}

func (p *PriorityQueue) Push(x interface{}) {
	*p = append(*p, elem{Value: x})
}

func main() {}