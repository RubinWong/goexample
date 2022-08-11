package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func test() {
	bin, err := exec.LookPath("ls")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(bin)

	args := []string{"ls", "-l", "-r", "-t"}
	env := os.Environ()
	err = syscall.Exec(bin, args, env)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	test()
}
