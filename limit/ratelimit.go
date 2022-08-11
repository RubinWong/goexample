package main

import (
	"fmt"
	"time"
)

type RateLimit struct {
	// The number of requests allowed per time unit.
	// Default value is `1`.
	Limit int64 `json:"limit"`
	tick  *time.Ticker
}

func NewRateLimit(lt int64, du time.Duration) *RateLimit {
	return &RateLimit{
		Limit: lt,
		tick:  time.NewTicker(du),
	}
}

func (r *RateLimit) Get() {
	<-r.tick.C
}

func main() {
	rl := NewRateLimit(1, time.Second)
	for i := 0; i < 10; i++ {
		rl.Get()
		fmt.Println("get", i)
	}
}