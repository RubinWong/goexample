package main

import (
	"fmt"
	"time"
)

func foo1(x *int) func() {
  return func() {
    *x = *x + 1
    fmt.Printf("foo1 val = %d\n", *x)
  }
}

func foo2(x int) func() {
	return func() {
	  x = x + 1
	  fmt.Printf("foo2 val = %d\n", x)
	}
}

func foo3() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
	  fmt.Printf("foo3 val = %d\n", val)
	}
}

func show(v interface{}) {
	fmt.Printf("foo4 val = %v\n", v)
}

func foo4() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go show(val)
	}
}

func main() {
	num := 223
	x := foo1(&num)
	x()

	arr := []int{1,2,3,4,5,6}

	for _, v := range arr {
		fmt.Println("val of ", v)
	}

	for _, v := range arr {
		func()  {
			fmt.Println("func val of ", v)
		}()  // func立即执行，每个循环都能获取i的值；
	}

	for _, v := range arr {
		go func() {
			fmt.Println("go func val of ", v)
		}() // goroutine调度，无法每次循环都获取i的值
	}

	for _, v := range arr {
		go func() {
			fmt.Println("go func with sleep val of ", v)
		}()
		time.Sleep(time.Duration(1) * time.Millisecond)
		// goroutine调度，无法每次循环都获取i的值；主线程sleep释放cpu，模拟轮流调度，可以打印出i的值
	}

	ch := make(chan int, 3)

	go func() {
		for v := range ch {
			fmt.Println("go func with chan val of ", v)
		}
	}()

	for _, v := range arr {
		ch <- v
	}
	close(ch)

    c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }

	// time.Sleep(time.Duration(2) * time.Second)
}