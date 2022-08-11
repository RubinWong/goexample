package main

import "fmt"

func Counter(f func()) func() int {
	i := 0
	return func() int {
		f()
		i++
		return i
	}
}

func Hello() {
	fmt.Println("Hello")
}

func main() {
	Cnt := Counter(Hello)
	Cnt()
	Cnt()
	Cnt()
	Cnt()
	Cnt()
	fmt.Println(Cnt())
}
