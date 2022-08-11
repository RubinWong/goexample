package main

import (
	"fmt"
	"time"
)

type MyError struct {
	when time.Time
	why  string
}

func run() error {
	return &MyError{
		time.Now(),
		"it works!",
	}
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.why)
}

func main() {
	/*for i := 0; i < 5; i++ {
		msg := run()
		fmt.Println(msg)
	}*/

	if err := run(); err != nil {
		fmt.Println(err)
	}

	//go say("Hello")
	//say("world")

	intarr := []int{8, 7, 6, 5, -9, 11, 12}
	c := make(chan int)
	go sum(intarr[:len(intarr)/2], c)
	go sum(intarr[len(intarr)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y)
}

func say(s string) {
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
