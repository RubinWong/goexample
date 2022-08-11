package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int, 100)

	go func ()  {
		for {
			ch <- rand.Int()
			// for {}
			// time.Sleep(time.Second * 10000)
		}
	}()

	for {
		select {
		case res := <- ch:
			fmt.Println(res)
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1")
		}
		time.Sleep(time.Second * 10000)
		// x := <- ch
		// fmt.Println(x)
	}
}