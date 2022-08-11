package main

import (
	"fmt"
	"time"
)

func main() {
	/* panic只会执行本协程内的defer
	 * main协程的defer不会被执行
	 */
	defer fmt.Println("defer main")

	go func() {
		defer fmt.Println("defer goroutine")
		panic("just")
	}()

	time.Sleep(time.Second)
}
