package main

import (
	"fmt"
	"reflect"
)

func main() {
	var m map[string]interface{}

	m = make(map[string]interface{})
	// var v uint32 = 1
	var vv uint64 = 666
	m["appData"] = uint32(1)
	m["streamId"] = vv

	if appData, ok := m["appData"].(uint32); ok {
		fmt.Println(appData)
	}
	if streamId, ok := m["streamId"].(uint64); ok {
		fmt.Println(streamId)
	}

	var a uint32 = 1
	fmt.Println(reflect.TypeOf(a + 1))
}
