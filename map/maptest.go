package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	//"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

type Vertext struct {
	height, width float64
}

func (v Vertext) Abs() float64 {
	return math.Sqrt(v.height*v.height + v.width*v.width)
}

func Abs(v Vertext) float64 {
	return math.Sqrt(v.height*v.height + v.width*v.width)
}

func (v *Vertext) Scale(x float64) {
	v.height = v.height * x
	v.width = v.width * x
}

var mm map[int]Vertext

func main() {
	m := make(map[string]Vertext)
	m["Hello"] = Vertext{54.1, 43.2}

	fmt.Println(m)

	mm = make(map[int]Vertext)
	mm[101] = Vertext{11.1, 101.1}
	fmt.Println(mm)

	wc.Test(WordCount)
	//pic.Show(Pic)

	var fn = func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(fn(3, 4))
	fmt.Println(compute(fn))
	fmt.Println(compute(math.Pow))

	pos, neg := addr(), addr()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(i))
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	vv := Vertext{12.4, 567.8}
	fmt.Println(vv.Abs(), Abs(vv))

	vv.Scale(2.1)
	fmt.Println(vv)
}

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	for _, substr := range strings.Fields(s) {
		res[substr]++
	}

	return res
}

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	//var res = make([]uint8)
	for i := 0; i < dy; i++ {
		res[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			res[i][j] = uint8(rand.Intn(256))
		}
	}
	return res
}

func compute(fn func(x, y float64) float64) float64 {
	return fn(6, 9)
}

func addr() func(x int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	res := 1
	s1 := 0
	i := 0
	return func() int {
		if i == 0 {
			res = 0
		} else if i == 1 {
			res = 1
		} else {
			res = res + s1
			s1 = res - s1
		}
		i++
		return res
	}
}
