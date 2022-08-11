package discovery

import (
	"context"
	"encoding/json"
	"errors"
	"learning/etcd/register"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var ErrLoadBalancingPolicy = errors.New("LoadBalancingPolicy is empty or not apply")

type NodeArray struct {
	Node []register.Options `json:"node"`
}

type Discovery struct {
	etcdCli *clientv3.Client
	// cc      resolver.ClientConn
	Node    sync.Map
	opts    *Options
}

func NewDiscovery(opt ...ClientOptions) (*Discovery, error) {
	s := &Discovery{
		opts: newOptions(opt...),
	}
	// if s.opts.LoadBalancingPolicy == VersionLB {
	// 	newVersionBuilder(s.opts)
	// } else {
	// 	return nil, ErrLoadBalancingPolicy
	// }
	etcdCli, err := clientv3.New(s.opts.EtcdConf)
	if err != nil {
		return nil, err
	}
	s.etcdCli = etcdCli
	return s, nil
}

// Build 当调用`grpc.Dial()`时执行
func (d *Discovery) Get() (*Discovery, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := d.etcdCli.Get(ctx, d.opts.SrvName, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	for _, v := range res.Kvs {
		if err = d.AddNode(v.Key, v.Value); err != nil {
			log.Println(err)
			continue
		}
	}
	if len(res.Kvs) == 0 {
		log.Printf("no %s service found , waiting for the service to join \n", d.opts.SrvName)
	}
	go func(dd *Discovery) {
		dd.watcher()
	}(d)
	return d, err
}

func (d *Discovery) AddNode(key, val []byte) error {
	var data = new(register.Options)
	err := json.Unmarshal(val, data)
	if err != nil {
		return err
	}

	d.Node.Store(string(key), data)
	return nil
}

func (d *Discovery) DelNode(key []byte) error {
	keyStr := string(key)
	d.Node.Delete(keyStr)
	return nil
}


func (d *Discovery) Scheme() string {
	return "discovery"
}

//watcher 监听前缀
func (d *Discovery) watcher() {
	rch := d.etcdCli.Watch(context.Background(), d.opts.SrvName, clientv3.WithPrefix())
	for res := range rch {
		for _, ev := range res.Events {
			switch ev.Type {
			case mvccpb.PUT: //新增或修改
				if err := d.AddNode(ev.Kv.Key, ev.Kv.Value); err != nil {
					log.Println(err)
				}
			case mvccpb.DELETE: //删除
				if err := d.DelNode(ev.Kv.Key); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func (s *Discovery) Close() {
	s.etcdCli.Close()
}
