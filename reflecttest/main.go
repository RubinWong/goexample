package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)
   
type Reindeer string 
  
func (r Reindeer) TakeOff() { 
    log.Printf("%q lifts off.", r) 
} 
  
func (r Reindeer) Land() { 
    log.Printf("%q gently lands.", r) 
} 
  
func (r Reindeer) ToggleNose() { 
    if r != "rudolph" { 
        panic("invalid reindeer operation") 
    } 
    log.Printf("%q nose changes state.", r) 
} 

func main() {
	reflectStruct()

	reflectKind()

	reflectTag()
}

func reflectStruct() {
	p:= Reindeer("rudolph") 
  
	pt := reflect.TypeOf(p)
	pv := reflect.ValueOf(p)
	numMethod := pt.NumMethod()
	for i := 0; i < numMethod; i++ {
		method := pt.Method(i)
		methodType := method.Type
		methodName := method.Name
		// methodName = utils.FirstToLower(methodName)
		methodValue := pv.Method(i)
		// NumIn include obj itself, so it is  eaqul normalNumIn+1
		fmt.Println(methodName, methodType.NumIn(), methodType.NumOut(), methodValue.Interface())

		// means has one input param
		if methodType.NumIn() == 2 {
			if methodType.NumOut() == 1 {
				// may be notification handlers
				outErr := methodType.Out(0)
				if outErr.Name() == "SignalError" {
					// p.conn.RegisterNotificationHandler(methodName, methodValue)
					fmt.Println(outErr.Name())
				}
			} else if methodType.NumOut() == 2 {
				// maybe request handlers
				outErr := methodType.Out(1)
				fmt.Println(outErr.Name(), outErr.Kind, outErr.String())
				if outErr.Name() == "SignalError" {
					// p.conn.RegisterRequestHandler(methodName, methodValue)
					fmt.Println(outErr.Name())
				}
			}
		}
	}
}

func reflectKind() {
	// var target []interface{} = 
	for i, v := range []interface{}{"Hello", 1, 2, 3, func() int {return 1} } {
		switch k := reflect.ValueOf(v); k.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				fmt.Println(i, "is int")
			case reflect.String:
				fmt.Println(i, "is string")
			default:
				fmt.Println(i, "is other ", k.Kind())
		}
	}
}

type S struct {
	A string `redis:"a"`
}

func reflectTag() {
	s := S{}
	s.A = "hello"
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.ValueOf(s))
	// fmt.Println(reflect.TypeOf(s).Elem())

	k := reflect.TypeOf(s)
	f := k.Field(0)
	fmt.Println(f.Tag)
	fmt.Println(f.Tag.Get("redis"))

	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Println(f.Type(), f.Kind(), f.Interface())
		SetField(f, "world")
	}
	fmt.Println(s)
}

func SetField(v reflect.Value, s string) error {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct but got %s", v.Kind())
	}
	t := v.Type()
	switch t.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64: {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return err
			}
			v.SetInt(i)
		}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64: {
			i, err := strconv.ParseUint(s, 10, 64)
			if err != nil {
				return err
			}
			v.SetUint(i)
		}
		case reflect.String:
			v.SetString(s)
		case reflect.Float32, reflect.Float64:
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return err
			}
			v.SetFloat(f)
		case reflect.Bool:
			b, err := strconv.ParseBool(s)
			if err != nil {
				return err
			}
			v.SetBool(b)
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
			return fmt.Errorf("unsupported type: %s", t.Kind())
		case reflect.Complex128, reflect.Complex64:
			return fmt.Errorf("unsupported type: %s", t.Kind())
		case reflect.Array:
			return fmt.Errorf("unsupported type: %s", t.Kind())
		case reflect.Struct:
			return fmt.Errorf("unsupported type: %s", t.Kind())
		case reflect.UnsafePointer:
			return fmt.Errorf("unsupported type: %s", t.Kind())
		case reflect.Invalid:
			return fmt.Errorf("unsupported type: %s", t.Kind())
		case reflect.Uintptr:
			return fmt.Errorf("unsupported type: %s", t.Kind())
		default:
			return fmt.Errorf("unsupported type: %s", t.Kind())
	}
	return nil
}