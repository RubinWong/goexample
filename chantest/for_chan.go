package main

import "fmt"

func sum(elems []int, c chan int) {
	res := 0
	for _, elem := range elems {
		res = res + elem;
	} 

	c <- res;
}

func main() {
	array  := [] int{7, 2, 8, -9, 4, 0}
	c := make(chan int, 2)
	go sum(array[:len(array)/2], c)
	go sum(array[len(array)/2:], c)

	x, y := <- c, <- c
	fmt.Println(x, y, x+y)
}
