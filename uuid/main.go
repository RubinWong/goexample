package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuid := uuid.New()
	fmt.Println(uuid, ", ", uuid.String(), ", ", uuid.ID(), ", ", len((uuid.String())))

	b, err := uuid.MarshalBinary()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, c := range b {
		fmt.Print(uint32(c), " ")
	}
	fmt.Println()
}