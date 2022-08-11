package main

import "fmt"

func main() {
	defer fmt.Println("main")

	defer func() {
		defer func() {
			panic("again and again")
		}()
		panic("again")
	}()
	panic("once")
}
