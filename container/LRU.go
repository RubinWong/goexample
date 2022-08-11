package main

import (
	"container/list"
	"fmt"
	"sync"
)

type LRU struct {
	l   *list.List
	m   map[interface{}]*list.Element
	cap int
	mu  *sync.Mutex
}

func NewLRU(c int) *LRU {
	l := list.New()
	m := make(map[interface{}]*list.Element)
	return &LRU{l, m, c, new(sync.Mutex)}
}

func (l *LRU) Add(key interface{}, val interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if e, ok := l.m[key]; ok {
		l.l.MoveToFront(e)
		e.Value = val
	} else {
		if l.l.Len() >= l.cap {
			del := l.l.Back()
			l.l.Remove(del)
			delete(l.m, del.Value)
		}
		e := l.l.PushFront(val)
		l.m[key] = e
	}
}

func (l *LRU) Del(key interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if e, ok := l.m[key]; ok {
		l.l.Remove(e)
		delete(l.m, key)
	}
}

func (l *LRU) Get(key interface{}) interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()
	if e, ok := l.m[key]; ok {
		l.l.MoveToFront(e)
		return e.Value
	}
	return nil
}

func main() {
	l := NewLRU(3)
	l.Add(1, 1)
	l.Add(2, 2)
	l.Add(3, 3)
	l.Add(4, 4)

	fmt.Println(l.Get(1))
	fmt.Println(l.Get(4))
	l.Add(5, 5)
	fmt.Println(l.Get(2))
}
