package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 4000,
	})
	if err != nil {
		fmt.Println("dial udp failed, er: ", err)
		return
	}
	defer socket.Close()
	data := []byte("hello")
	_, err = socjet.Write(data)
	if err != nil {
		fmt.Println("write to udp failed, er: ", err)
		return
	}
	data = make([]byte, 1024)
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("read from udp failed, er: ", err)
		return
	}
	fmt.Println("read from udp, data: ", string(data[:n]), " addr: ", remoteAddr)
}