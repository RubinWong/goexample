package main

import (
	"math/rand"
)

func main() {
	for i := 0; i < 50; i++ {
		r := rand.Int()
		println(r, r % 100 < 100)
	}
}