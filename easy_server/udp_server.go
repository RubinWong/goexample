package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 4000,
	})
	if err != nil {
		fmt.Println("listen failed, er: ", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from udp failed, er: ", err)
			continue
		}
		fmt.Println("read from udp, data: ", string(data[:n]), " addr: ", addr)
		_, err = listen.WriteToUDP(data[:], addr)
		if err != nil {
			fmt.Println("write to udp failed, er: ", err)
			continue
		}
	}
}