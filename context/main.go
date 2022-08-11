package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// go doSomething()
	doSomethingWithDeadline()

	waitSignal()
}

func doSomething() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c := make(chan int, 3)

	go doAnother(ctx, c)

	for i := 0; i < 10; i++ {
		c <- i * 20
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println("doSomethingWithDeadline finish")
	time.Sleep(time.Millisecond * 100)
}

func doSomethingWithDeadline() {
	ctx0 := context.Background()
	deadline := time.Now().Add(time.Millisecond * 1500)

	ctx, cancel := context.WithDeadline(ctx0, deadline)
	defer cancel()

	c := make(chan int, 3)
	go doAnother(ctx, c)

	for i := 0; i < 100; i++ {
		if err := ctx.Err(); err != nil {
			fmt.Println("doSomethingWithDeadline context deadline: ", err.Error())
			break
		}
		c <- i * 10
		time.Sleep(time.Millisecond * 50)
	}

	fmt.Println("doSomethingWithDeadline finish")
}

func doAnother(ctx context.Context, c chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Println("context err: ", err.Error())
			}
			fmt.Println("context done")
			return
		case num := <-c:
			fmt.Println("receive a num: ", num)
		}
	}
}

func waitSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for a := range c {
		switch a {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("exit")
			return
		default:
			return
		}
	}
}
