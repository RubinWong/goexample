package main

import (
	"fmt"
	"time"
)

func main() {
	tm := time.Now().UnixNano()
	fmt.Println(tm)

	tm2 := tm + int64(time.Second*5)
	fmt.Println(tm2)

	ftm := float64(tm2)
	fmt.Println(ftm)
}