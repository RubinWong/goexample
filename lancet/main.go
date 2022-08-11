package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/duke-git/lancet/netutil"
	"github.com/duke-git/lancet/retry"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s := "Hello, World!"
	fmt.Println(strutil.CamelCase(s))

	fmt.Println(netutil.GetIps())
	fmt.Println(netutil.GetInternalIp())

	ctx, cancel := context.WithCancel(context.TODO())
	var num int
	increase := func() error {
		num++
		fmt.Println(num)
		if num == 5 {
			cancel()
			return errors.New("don't work with five")
		}
		return errors.New("don't work")
	}

	err := retry.Retry(increase, retry.RetryDuration(1*time.Second), retry.Context(ctx))
	if err != nil {
		fmt.Println(err)
	}

	HTTPGet()
}

func httpGet(ctx context.Context, cancel context.CancelFunc) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		select {
		// this timeout is not effective
		case <-time.After(time.Second * 2):
			fmt.Println("timeout")
			cancel()
			return
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Println("still doing")
		}
	}
}

func HTTPGet() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer time.Sleep(time.Second * 6)
	go httpGet(ctx, cancel)

	// for i := 0; i < 5; i++ {
	// 	select {
	// 	case <-time.After(time.Second * 3):
	// 		fmt.Println("timeout")
	// 		cancel()
	// 	default:
	// 		time.Sleep(time.Second)
	// 	}
	// }
}
