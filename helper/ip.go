package main

import (
	"fmt"
	"net"
)

func GetIP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		// fmt.Errorf("Error: %v", err)
		fmt.Println("Error: ", err)
		return ""
	}
	for _, networkInterface := range interfaces {
		addresses, err := networkInterface.Addrs()
		if err != nil {
			fmt.Println("Error: ", err)
		}
		for _, address := range addresses {
			switch v := address.(type) {
			case *net.IPNet:
				if !v.IP.IsLoopback() {
					ip := v.IP.To4()
					if ip != nil {
						return ip.String()
					}
				}
			}
		}
	}
	return ""
}