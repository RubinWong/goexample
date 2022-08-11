package sync

import (
	"fmt"
	"runtime/debug"
)

// WorkerPool provides a pool for goroutines
type WorkerPool interface {
	// Schedule try to acquire pooled worker goroutine to execute the specified task,
	// this method would block if no worker goroutine is available
	Schedule(task func())

	// Schedule try to acquire pooled worker goroutine to execute the specified task first,
	// but would not block if no worker goroutine is available. A temp goroutine will be created for task execution.
	ScheduleAlways(task func())

	ScheduleAuto(task func())
}

type workerPool struct {
	work chan func()
	sem  chan struct{}
}

// NewWorkerPool create a worker pool
func NewWorkerPool(size int) WorkerPool {
	return &workerPool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

func (p *workerPool) Schedule(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.spawnWorker(task)
	}
}

func (p *workerPool) ScheduleAlways(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.spawnWorker(task)
	default:
		// new temp goroutine for task execution
		// if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		// 	log.DefaultLogger.Debugf("[syncpool] workerpool new goroutine")
		// }
		fmt.Println("[syncpool] workerpool new goroutine")
		GoWithRecover(func() {
			task()
		}, nil)
	}
}

func (p *workerPool) ScheduleAuto(task func()) {
	select {
	case p.work <- task:
		return
	default:
	}
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.spawnWorker(task)
	default:
		// new temp goroutine for task execution
		// if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		// 	log.DefaultLogger.Debugf("[syncpool] workerpool new goroutine")
		// }
		fmt.Println("[syncpool] workerpool new goroutine")
		GoWithRecover(func() {
			task()
		}, nil)
	}
}

func (p *workerPool) spawnWorker(task func()) {
	defer func() {
		if r := recover(); r != nil {
			// log.DefaultLogger.Alertf("syncpool", "[syncpool] panic %v\n%s", p, string(debug.Stack()))
			fmt.Println("[syncpool] panic", r, string(debug.Stack()))
		}
		<-p.sem
	}()
	for {
		task()
		task = <-p.work
	}
}

// GoWithRecover wraps a `go func()` with recover()
func GoWithRecover(handler func(), recoverHandler func(r interface{})) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// recoverLogger(os.Stderr, r)
				fmt.Println(r)
				if recoverHandler != nil {
					go func() {
						defer func() {
							if p := recover(); p != nil {
								// recoverLogger(os.Stderr, p)
								fmt.Println(p)
							}
						}()
						recoverHandler(r)
					}()
				}
			}
		}()
		handler()
	}()
}
