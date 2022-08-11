package register

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Register struct {
	etcdCli       *clientv3.Client
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	Opts          *Options
	name          string
}

func NewRegister(opt ...RegisterOptions) (*Register, error) {
	s := &Register{
		Opts: newOptions(opt...),
	}
	var ctx, cancel = context.WithTimeout(context.Background(), time.Duration(s.Opts.RegisterTTL)*time.Second)
	defer cancel()
	data, err := json.Marshal(s.Opts)
	if err != nil {
		return nil, err
	}
	etcdCli, err := clientv3.New(s.Opts.EtcdConf)
	if err != nil {
		return nil, err
	}
	s.etcdCli = etcdCli
	//申请租约
	resp, err := etcdCli.Grant(ctx, s.Opts.RegisterTTL)
	if err != nil {
		return s, err
	}
	s.name = fmt.Sprintf("%s/%v", s.Opts.Node.Path, s.Opts.Node.ID)
	//注册节点
	_, err = etcdCli.Put(ctx, s.name, string(data), clientv3.WithLease(resp.ID))
	if err != nil {
		return s, err
	}
	//续约租约
	s.keepAliveChan, err = etcdCli.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (s *Register) ListenKeepAliveChan() (isClose bool) {
	for range s.keepAliveChan {
	}
	return true
}

// Close 注销服务
func (s *Register) Close() error {
	if _, err := s.etcdCli.Delete(context.Background(), s.name); err != nil {
		return err
	}
	return s.etcdCli.Close()
}
