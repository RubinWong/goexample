package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup
	tick := time.Now()

	urls := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.golang.org",
	}

	for _, url := range urls {
		wg.Add(1)
		// defer wg.done()
		go func(url string) {
			tick := time.Now()
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			fmt.Println(resp.Status, url, time.Since(tick))
		}(url)
	}
	wg.Wait()
	fmt.Println("All goroutines finished in:", time.Since(tick))

	for _, url := range urls {
		wg.Add(1)
		go worker(url, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished in:", time.Since(tick))

	sem := make(chan struct{}, 3)
	defer close(sem)
	for i := 0; i < 10; i++ {
		select {
		case sem <- struct{}{}:
			fmt.Println("Semaphore acquired")
			go func(i int) {
				// defer func() { <-sem }()
				fmt.Println("Goroutine", i)
			}(i)
		default:
			fmt.Println("Semaphore blocked")
		}
	}

	counter()
}

func worker(url string, wg *sync.WaitGroup) {
	tick := time.Now()
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status, url, time.Since(tick))
}

func counter() {
	var wt sync.WaitGroup

	var c int64
	for i := 0; i < 50; i++ {
		wt.Add(1)
		go func() {
			defer wt.Done()
			for j := 0; j < 5000; j++ {
				atomic.AddInt64(&c, 1)
			}
		}()
	}

	wt.Wait()

	fmt.Println("Counter:", c)
}
