package main

import "fmt"

var i int

func GenSeq() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func intSeq() func() int {
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := GenSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := GenSeq()
	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())

	i = 10
	intgen := intSeq()
	i = 20
	fmt.Println(intgen(), i)
	i = 30
	// closure can access the variable i in the outer function
	// close holds a reference to the outer variable i
	fmt.Println(intgen(), i)
	fmt.Println(intgen(), i)
	fmt.Println(intgen(), i)

	// Instance() = "Axe2"
	name := Instance()
	fmt.Println(name, "heal")
	name = "Axe3"
	fmt.Println(name, "heal")
	name2 := Instance()
	fmt.Println(name2, "heal")

	fmt.Println(fibanacci(10))
	fmt.Println(fibanacci2(10))
	fmt.Println(fibanacci4(10))
}

func sigleton() func() string {
	// closure holds a reference to variable inst which can't be modifed outside the function
	inst := "Axe"
	return func() string {
		return inst
	}
}

var Instance = sigleton()

func fibanacci(n int) int {
	if n < 2 {
		return n
	}
	return fibanacci(n-1) + fibanacci(n-2)
}

func fibanacci2(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func fibanacci3() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func fibanacci4(n int) int {
	f := fibanacci3()
	for i := 1; i < n; i++ {
		f()
	}
	return f()
}
