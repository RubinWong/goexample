package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		return
	}

	server := os.Args[1]
	addr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Addr: ", addr)

	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println("Dail success:", conn)
	defer conn.Close()

	tm := time.NewTimer(time.Second * 10)

	for {
		select {
		case <-tm.C:
			fmt.Println("time to close")
			break
		default:
			_, err = conn.Write([]byte("hello world"))
			if err != nil {
				fmt.Println(err)
				os.Exit(3)
			}
			fmt.Println("write success")

			var buf [1024]byte
			n, err := conn.Read(buf[:])
			if err != nil {
				fmt.Println("read fail", err.Error())
				break
			}
			fmt.Println("read success, len", n, string(buf[:n]))
			time.Sleep(time.Second * 1)
		}
	}
}
