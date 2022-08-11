package main

import (
	"fmt"
)

type Message struct {
	id   int
	addr string
}

func NewMessage(id int, addr string) Message {
	return Message{
		id:   id,
		addr: addr,
	}
}

type Notify struct {
	id    int
	addr  string
	mType int
}
var defNotify = Notify{id: 0, addr: "123", mType: 1}

type Option func(*Notify)

func NewNotify(opts ...Option) Notify {
	m := defNotify
	for _, o := range opts {
		o(&m)
	}
	return m
}

func WithID(id int) func(*Notify) {
	return func(n *Notify) {
		n.id = id
	}
}

func WithAddr(addr string) func(*Notify) {
	return func(n *Notify) {
		n.addr = addr
	}
}

func WithMessageType(mType int) func(*Notify) {
	return func(n *Notify) {
		n.mType = mType
	}
}

func main() {
	m := NewMessage(1, "123")
	fmt.Println(m)

	n := NewNotify(WithID(123), WithAddr("my addr"), WithMessageType(3))
	fmt.Println(n)
}
