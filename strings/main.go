package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.FieldsFunc("Hellow world \rnewbee\tto go", func(r rune) bool {
		return unicode.IsSpace(r)
	}))
}
