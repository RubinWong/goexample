package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var cc, java, python, objectc bool

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func splite(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func myPow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Println(v, lim)
		return lim
	}
}

func sqrt(x float64) float64 {
	z := float64(x / 2)
	dis := float64(1)
	for i := 1; math.Abs(dis) > 0.00001; i++ {
		dis = (z*z - x) / (2 * z)
		z -= dis
		fmt.Printf("now dis %f loop count %d\n", dis, i)
	}
	return z
}

type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println("My favourite Number is ", rand.Intn(100))
	fmt.Println("pi is ", math.Pi)
	fmt.Println(add(1, 100))

	a, b := swap("easy", "hard")
	fmt.Println(swap("world", "Hello "))
	fmt.Println(a, b)

	c, d := splite(9)
	fmt.Println("c and d is ", c, d)

	var i int
	fmt.Println(cc, java, python, objectc, i)

	var m, n, s, t = true, 1, false, "No!"
	fmt.Println(m, n, s, t)

	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Printf("sum %d\n", sum)

	sum = 1
	i = 0
	for sum < 100 {
		sum += sum
		i++
		if sum > 50 {
			fmt.Printf("now sum is %d\n", sum)
		}
	}
	fmt.Printf("loop count %d, sum %d\n", i, sum)

	/*defer fmt.Println(
		myPow(2, 3, 10),
		myPow(3, 3, 20),
	)*/

	sqrt(9)
	//sqrt(8)
	//sqrt(100.1)
	//sqrt(1000000.1)
	//sqrt(1000000000.1)

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("mac os")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println(os)
	}

	switch today := time.Now().Weekday(); time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("the day after tomorrow")
	default:
		fmt.Println("far away")
	}

	/*fmt.Println("counting...")
	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("count end")*/

	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.x = 4
	v.y = 8
	fmt.Println(v)

	p := &v
	p.x = 16
	fmt.Println(p, *p)

	v1 := Vertex{x: 1}
	v2 := Vertex{y: 108}
	p2 := &Vertex{}
	fmt.Println(v1, v2, p2)

	var str [2]string
	str[0] = "Hello"
	str[1] = "world"
	fmt.Println(str[0], str[1], str)

	arr := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)

	var slice []int = arr[1:4]
	fmt.Println(slice)

	var silce2 []int = arr[3:6]
	silce2[0] = 100
	fmt.Println(slice, silce2, len(silce2), cap(silce2), len(slice), cap(slice))

	ss := []struct {
		numb  int
		value bool
	}{
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{6, false},
	}

	fmt.Println(ss)

	sli3 := make([]int, 6)
	printSlice("sli3", sli3)

	sli4 := make([]int, 0, 5)

	sli5 := sli4[:4]

	printSlice("sli4", sli4)
	printSlice("sli5", sli5)

	var pp []int
	for i := 0; i < 10; i++ {
		pp = append(pp, i*i)
	}
	pp = append(pp, 2, 5, 6, 7, 19)

	for i, v := range pp {
		fmt.Println(i, v)
	}

	for _, v := range pp {
		fmt.Println(v)
	}

	for i, _ := range pp {
		fmt.Println(pp[i])
	}

	for i := range pp {
		fmt.Println(pp[i])
	}
}

func printSlice(str string, x []int) {
	fmt.Printf("%s len %d cap %d\n", str, len(x), cap(x))
}
