package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func  main()  {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
	}

	server := os.Args[1]
	addr, err := net.ResolveTCPAddr("tcp4", server)
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	response, _ := ioutil.ReadAll(conn)
	fmt.Println(string(response))
	os.Exit(0)
}