package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

func main() {
	tcpServer, _ := net.ResolveTCPAddr("tcp4", ":4000")

	listener, _ := net.ListenTCP("tcp", tcpServer)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	go func() {
		response, _ := ioutil.ReadAll(conn)
		fmt.Println(string(response))
	}()

	time.Sleep(1 * time.Second)
	now := time.Now().String()
	conn.Write([]byte(now))
}
