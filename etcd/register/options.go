package register

import (
	"fmt"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Node struct {
	Name    string   `json:"name"`
	ID      uint32   `json:"id"`
	Path    string   `json:"path"`
	Address []string `json:"address"`
	SvrType string   `json:"server_type"`
	Idc     string   `json:"idc"`
	Version string   `json:"version"`
}

type Options struct {
	EtcdConf    clientv3.Config      `json:"-"`
	RegisterTTL int64                `json:"-"`
	Node        *Node                `json:"node"`
	Metadata    map[string]string    `json:"metadata"`
}

type RegisterOptions func(*Options)

func newOptions(opt ...RegisterOptions) *Options {
	opts := &Options{
		EtcdConf: clientv3.Config{
			Endpoints:   []string{"127.0.0.1:2379"},
			DialTimeout: 5 * time.Second,
		},
		Node:        &Node{Version: "latest"},
		RegisterTTL: 10,
	}
	for _, o := range opt {
		o(opts)
	}
	return opts
}

func SetName(name string) RegisterOptions {
	return func(options *Options) {
		path := strings.Split(name, ".")
		options.Node.Name = path[len(path)-1]
		// options.Node.ID = fmt.Sprintf("%s", uuid.Must(uuid.NewUUID()).String())
		options.Node.Path = fmt.Sprintf("/%s", strings.Join(path, "/"))
		fmt.Println(path, options.Node.Name, options.Node.ID, options.Node.Path)
	}
}

func SetID(id uint32) RegisterOptions {
	return func(options *Options) {
		options.Node.ID = id
	}
}

func SetServerType(svrType string) RegisterOptions {
	return func(options *Options) {
		options.Node.SvrType = svrType
	}
}

func SetIDC(idc string) RegisterOptions {
	return func(options *Options) {
		options.Node.Idc = idc
	}
}

func SetRegisterTTL(registerTTL int64) RegisterOptions {
	return func(options *Options) {
		options.RegisterTTL = registerTTL
	}
}

func SetVersion(version string) RegisterOptions {
	return func(options *Options) {
		options.Node.Version = version
	}
}
func SetEtcdConf(conf clientv3.Config) RegisterOptions {
	return func(options *Options) {
		options.EtcdConf = conf
	}
}

func SetAddress(address []string) RegisterOptions {
	return func(options *Options) {
		options.Node.Address = address
	}
}

func SetMetadata(metadata map[string]string) RegisterOptions {
	return func(options *Options) {
		options.Metadata = metadata
	}
}
