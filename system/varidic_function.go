package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	ints := []int{6, 7, 8, 9, 10}

	fmt.Println(sum(ints...))
}

func sum(ints ...int) int {
	total := 0
	// for...range
	for _, num := range ints {
		total += num
	}
	return total
}
