package main

import (
	"fmt"
	"learning/etcd/register"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	s, err := register.NewRegister(
		register.SetName("services.frontend.sfu"),
		register.SetID(rand.Uint32()),
		register.SetAddress([]string{"10.12.112.165:1231"}),
		register.SetVersion("v1"),
		register.SetEtcdConf(clientv3.Config{
			Endpoints:   []string{"10.12.208.163:2371", "10.12.208.163:2381", "10.12.208.163:2391"},
			DialTimeout: time.Second * 3,
		}),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(s.Opts)
	fmt.Println(*s.Opts.Node)

	c := make(chan os.Signal, 1)
	go func() {
		if s.ListenKeepAliveChan() {
			c <- syscall.SIGQUIT
		}
	}()
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for a := range c {
		switch a {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("退出")
			_ = s.Close()
			return
		default:
			return
		}
	}
}
