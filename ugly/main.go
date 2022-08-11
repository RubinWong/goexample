package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	{ /* 三种不同时间
		*
		 */
		fmt.Println(time.Now().UnixNano())
		fmt.Println(time.Now().Nanosecond())
		fmt.Println(time.Now().Unix())
	}

	{
		/* defer 函数参数优先计算值；
		 * defer 函数后调先触发
		 */
		a := 1
		b := 2
		defer calc("1", a, calc("10", a, b))
		a = 0
		defer calc("2", a, calc("20", a, b))
		b = 1
	}

	{
		/* make 拥有默认值 */
		s := make([]int, 5)
		s = append(s, 1, 2, 3)
		fmt.Println("slice: ", s)
	}
	{
		/* var 无默认值，append是否有内存分配?? */
		var s []int
		s = append(s, 1, 2, 3)
		fmt.Println("slice: ", s)
	}

	{
		// var peo People = Student{}
		var peo People = &Student{}
		think := "bitch"
		fmt.Println(peo.Speak(think))
	}

	{
		if live() == nil {
			fmt.Println("AAAAAAA")
		} else {
			fmt.Println("BBBBBBB")
		}
	}

	{
		s := new(Show)
		s.Param = make(Param)
		s.Param["RMB"] = 10000
	}

	{
		js := `{
        "name":"11"
    	}`
		var p People2
		err := json.Unmarshal([]byte(js), &p)
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		fmt.Println("people: ", p)
	}

	{
		p := &People3{}
		fmt.Println(p.String())
	}
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func live() People {
	var stu *Student
	return stu
}

type Param map[string]interface{}

type Show struct {
	Param
}

type People2 struct {
	name string `json:"name"`
}

type People3 struct {
	Name string
}

func (p *People3) String() string {
	return fmt.Sprintf("print: %v", p)
}
