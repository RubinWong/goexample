package main

import (
	"fmt"
	"learning/etcd/discovery"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	s, err := discovery.NewDiscovery(
		discovery.SetName("services.frontend.sfu"),
		discovery.SetEtcdConf(clientv3.Config{
			Endpoints:   []string{"10.12.208.163:2371", "10.12.208.163:2381", "10.12.208.163:2391"},
			DialTimeout: time.Second * 3,
		}),
	)

	if err != nil {
		panic(err)
	}

	if _, err := s.Get(); err != nil {
		panic(err)
	}

	s.Node.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
