package main

import (
	"fmt"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))
	// fmt.Println(dateOut)

	// lsCmd := exec.Command("ls", " -a")
	lsCmd := exec.Command("zsh", "-c", "ls -lrt")
	lsOut, err := lsCmd.Output()
	if err != nil {
		fmt.Println(err)
		switch e := err.(type) {
		case *exec.ExitError:
			fmt.Println("ls command exit rc=", e.ExitCode())
		case *exec.Error:
			fmt.Println("ls failed")
		default:
			fmt.Println("ls failed unknown error")
		}
		panic(err)
	}

	// fmt.Println(lsOut)
	fmt.Println(string(lsOut))

	test()
}
