package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr("tcp4", ":4000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	listener, err := net.ListenTCP("tcp", tcpServer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

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

	for {
		// var buf [4096]byte
		var buf []byte
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("read success, size", n)

		if n == 0 {
			continue
		}

		fmt.Println(string(buf))

		now := time.Now().String()

		res := new(bytes.Buffer)

		binary.Write(res, binary.BigEndian, []byte(now))
		binary.Write(res, binary.BigEndian, []byte("  "))
		binary.Write(res, binary.BigEndian, buf[:n])
		n, err = conn.Write(res.Bytes())
		if err != nil {
			fmt.Println("write fail", err.Error())
			return
		}
		fmt.Println("write success, size", n)
	}
}
