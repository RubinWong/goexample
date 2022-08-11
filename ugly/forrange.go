package main

import "fmt"

func main() {
	/* exporting a pointer for the loop variable v (exportloopref)
	 */
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}

	/* for-range遍历map时，不保证遍历顺序
	 * 每次执行打印的结果可能都不一样
	 */
	hash := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}

	for k, v := range hash {
		fmt.Println(k, v)
	}

	/* defer特性：
	 * 1. 所有defer组成一个链表，后调用的defer放在队头，先被执行；
	 * 2. defer调用函数的参数会被优先计算
	 */
	num := 1
	defer fmt.Println("first defer", num)
	defer func(num int) { fmt.Println("second defer", num) }(num)
	num = 4
	defer func() { fmt.Println("third defer", num) }()
}
