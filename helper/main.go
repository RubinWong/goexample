package main

import (
	"fmt"
	"net"
	"os"
	"time"
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

func checkParentAlive() {
	parentPid := os.Getppid()
	go func() {
		for {
			if parentPid == 1 || os.Getppid() != parentPid {
				fmt.Println("parent no alive, exit")
				os.Exit(0)
			}
			_, err := os.FindProcess(parentPid)
			if err != nil {
				fmt.Println("parent no alive, exit")
				os.Exit(0)
			}

			time.Sleep(5 * time.Second)
		}
	}()
}

func main() {
	fmt.Println(GetIP())

	go checkParentAlive()

	time.Sleep(time.Second * 1)
}
