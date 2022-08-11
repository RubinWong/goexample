package main

import (
	"fmt"
	"time"
)

func main() {
	test()
	fmt.Println("main continues")
	test1()
	fmt.Println("main continues")
	test2()
	fmt.Println("main continues")
	time.Sleep(time.Second * 2)
}

func Panic() {
	panic("A problem accured")
}

func test() {
	defer func() {
		// recover must be called within a deferred function. in the same goroutine
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()

	Panic()
	// code after panic will not be executed
	fmt.Println("Program continues")
}

func test1() {
	defer RecoverWithCallback(nil)
	Panic()
}

func test2() {
	GoWithRecover(Panic, nil)
}

func RecoverWithCallback(f func()) {
	if err := recover(); err != nil {
		fmt.Println("panic:", err)
		if f != nil {
			f()
		}
	}
}

func GoWithRecover(handler func(), f func(r interface{})) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic:", err)
				if f != nil {
					go func() {
						defer func() {
							if err := recover(); err != nil {
								fmt.Println("panic:", err)
							}
						}()
						f(err)
					}()
				}
			}
		}()
		handler()
	}()
}
