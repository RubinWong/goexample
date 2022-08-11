package main

import (
	"fmt"
	"math"
	"time"
)

type ABSer interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	x, y MyFloat
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return math.Sqrt(float64(-f))
	} else {
		return math.Sqrt(float64(f))
	}
}

func (v Vertex) Abs() float64 {
	X := float64(v.x)
	Y := float64(v.y)
	return math.Sqrt(X*X + Y*Y)
}

func main() {
	var a ABSer
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a)

	a = &v
	fmt.Println(a)

	var i I

	ff := F(math.Pi)
	ff.M()
	i = ff
	describe(i)

	v.M()
	i = v
	i.M()
	describe(i)

	var ii I

	var vv *Vertex
	ii = vv // a interface with nil value
	//ii.M()
	describe(ii)

	var iii I // a nil interface
	//iii.M()
	describe(iii)

	var in interface{}
	describe1(in)

	num := 664
	describe1(num)

	stri := "mamamiya"
	describe1(stri)

	var gents uint32 = uint32(time.Now().Unix())
	fmt.Println("%u", gents)
}

type I interface {
	M()
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func (v Vertex) M() {
	fmt.Println(v)
}

func describe(i I) {
	fmt.Println(i)
}

func describe1(i interface{}) {
	fmt.Printf("(%v %T)\n", i, i)
}
