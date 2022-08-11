package main

import (
	"container/list"
	"fmt"
)

func main() {
	test()
	test1()
}

func test() {
	// list is a doubly linked list. Each element has a 'Value' and a 'Next' and 'Prev' pointer.
	// PushFront() adds a new element to the front of the list.
	l := list.New()

	e1 := l.PushFront(1)
	e4 := l.PushBack(4)

	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func test1() {
	l := list.New()
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}

	for i := 0; i < 10; i++ {
		l.PushFront(i)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value.(int), " ")
	}
	fmt.Println()

	e := l.Front()

	l.InsertBefore(4, e)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value.(int), " ")
	}

	fmt.Println()
}

// type Element
//     func (e *Element) Next() *Element                                   // 返回该元素的下一个元素，如果没有下一个元素则返回 nil
//     func (e *Element) Prev() *Element                                   // 返回该元素的前一个元素，如果没有前一个元素则返回nil

// type List                               
//     func New() *List                                                    // 返回一个初始化的list
//     func (l *List) Back() *Element                                      // 获取list l的最后一个元素
//     func (l *List) Front() *Element                                     // 获取list l的最后一个元素
//     func (l *List) Init() *List                                         // list l 初始化或者清除 list l
//     func (l *List) InsertAfter(v interface{}, mark *Element) *Element   // 在 list l 中元素 mark 之后插入一个值为 v 的元素，并返回该元素，如果 mark 不是list中元素，则 list 不改变
//     func (l *List) InsertBefore(v interface{}, mark *Element) *Element  // 在 list l 中元素 mark 之前插入一个值为 v 的元素，并返回该元素，如果 mark 不是list中元素，则 list 不改变
//     func (l *List) Len() int                                            // 获取 list l 的长度
//     func (l *List) MoveAfter(e, mark *Element)                          // 将元素 e 移动到元素 mark 之后，如果元素e 或者 mark 不属于 list l，或者 e==mark，则 list l 不改变
//     func (l *List) MoveBefore(e, mark *Element)                         // 将元素 e 移动到元素 mark 之前，如果元素e 或者 mark 不属于 list l，或者 e==mark，则 list l 不改变
//     func (l *List) MoveToBack(e *Element)                               // 将元素 e 移动到 list l 的末尾，如果 e 不属于list l，则list不改变             
//     func (l *List) MoveToFront(e *Element)                              // 将元素 e 移动到 list l 的首部，如果 e 不属于list l，则list不改变             
//     func (l *List) PushBack(v interface{}) *Element                     // 在 list l 的末尾插入值为 v 的元素，并返回该元素              
//     func (l *List) PushBackList(other *List)                            // 在 list l 的尾部插入另外一个 list，其中l 和 other 可以相等               
//     func (l *List) PushFront(v interface{}) *Element                    // 在 list l 的首部插入值为 v 的元素，并返回该元素              
//     func (l *List) PushFrontList(other *List)                           // 在 list l 的首部插入另外一个 list，其中 l 和 other 可以相等              
//     func (l *List) Remove(e *Element) interface{}                       // 如果元素 e 属于list l，将其从 list 中删除，并返回元素 e 的值
