package main

import (
	"fmt"
)

// Output: 55, array passed by value, not by reference
func sumArray(arr [10]int) int {
	sum := 0
	fmt.Printf("%p\n", &arr)
	for _, v := range arr {
		sum += v
	}
	arr[5] = 7
	return sum
}

func sumArrayPointer(arr *[10]int) int {
	sum := 0
	fmt.Printf("%p\n", arr)
	for _, v := range arr {
		sum += v
	}
	return sum
}

// Output: 55, slice passed by reference, not by value
func sumSlice(sli []int) int {
	sum := 0
	fmt.Printf("%p\n", sli)
	for _, v := range sli {
		sum += v
	}
	sli[5] = 7
	return sum
}

func sumSlicePointer(sli *[]int) int {
	sum := 0
	fmt.Printf("%p\n", *sli)
	for _, v := range *sli {
		sum += v
	}
	return sum
}

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("%p\n", &arr)
	sumArray(arr)
	fmt.Println(arr)
	sumArrayPointer(&arr)
	fmt.Println()

	fmt.Printf("%p\n", sli)
	sumSlice(sli)
	fmt.Println(sli)
	sumSlicePointer(&sli)
}
