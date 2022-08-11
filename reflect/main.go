package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id int
	Name string
}

func (s Student) Hello() {
	fmt.Println("我是一个学生.")
}

func main() {
	reflectView()
	reflectModify()
	reflectFunc(Add)
}

func reflectView() {
	s := Student{Id: 1, Name:"羊驼"}

	t := reflect.TypeOf(s)
	fmt.Println("class type: ", t.Name())

	v := reflect.ValueOf(s)
	for i:= 0; i < t.NumField(); i++ {
		key := t.Field(i)

		value := v.Field(i).Interface()

		fmt.Printf("第%d个字段时：%s:%v = %v\n", i + 1, key.Name, key.Type, value)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("第%d个方法是: %s:%v\n", i + 1, m.Name, m.Type)
	}

	for i := 0; i < v.NumField(); i++ {
		fmt.Println("field of v: ", v.Field(i).Interface())
	}

}

func reflectModify() {
	s := &Student{Id: 1, Name:"羊驼"}

	t := reflect.TypeOf(s)
	fmt.Println("class type: ", t.Name())

	v := reflect.ValueOf(s)
	// 通过反射修改字端
	if v.Kind() != reflect.Ptr {
		fmt.Println("not ptr type, can't modify")
		return;
	}

	v = v.Elem()
	name := v.FieldByName("Name")

	if name.Kind() == reflect.String {
		name.SetString("小学生")
	}

	fmt.Printf("%#v \n", *s)
}

func Add(a int, b int) (res int, err error) {
	res = a + b
	err = nil
	return
}

type ginwebRequest interface {
	Parse() error
	Validate() error
}

func reflectFunc(f interface{}) {
	t := reflect.TypeOf(f)
	fmt.Println(t.Kind())
	if t.Kind() != reflect.Func {
		fmt.Println("not a func")
		return
	}

	fmt.Println("in parameters: ", t.NumIn())
	fmt.Println("out parameters: ", t.NumOut())

	numIn := t.NumIn()
	requestFields := []int{}
	for i := 0; i < numIn; i++ {
		if t.In(i).Implements(reflect.TypeOf((*ginwebRequest)(nil)).Elem()) {
			requestFields = append(requestFields, i)
		}
	}

	fmt.Println("len of request size: ", len(requestFields))
}