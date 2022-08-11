package main

import "fmt"

func FillMap(m map[string]interface{}) {
	m["roomId"] = uint32(1)
	m["PeerId"] = uint64(666)
}

func main() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["appData"] = uint32(1)
	m["streamId"] = uint64(666)

	FillMap(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
