package main

import "fmt"

func main() {
	/*
	 * recover必须发生在panic之后，因此这里的recover不会触发
	 * 如果想要recover捕捉到panic，必须将recover放在defer中
	 */
	defer fmt.Println("main")
	if err := recover(); err != nil {
		fmt.Println(err)
		return
	}

	panic("just")
}
