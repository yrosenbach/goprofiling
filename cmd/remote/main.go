package main

// 'remote' is a sample program on how to expose remotely profiles to the pprof tool by a long lived program.

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // will install http handlers on the default router of /net/http (http.DefaultServeMux)
	"sync"
	"time"

	// for exposing runtime profiling data to be used by the pprof tool.
)

var (
	s []string
)

const (
	maxLen = 1000000
	profilerAddr   = "localhost:6060"
)

func main() {

	go func() {
		// http.ListenAndServe may return an error so we would like to print it just in case ...
		fmt.Println(http.ListenAndServe(profilerAddr, nil)) // The handler is nil, which means to use DefaultServeMux
	}()

	workers := 2
	mutex := &sync.Mutex{}

	// do some work...
	for i := 0; i < workers; i++ {
		go func(i int) {
			workerName := fmt.Sprintf("worker-%d", i)
			for {
				if len(s) < maxLen {
					mutex.Lock()
					s = append(s, workerName)
					mutex.Unlock()
				} else {
					time.Sleep(50 * time.Millisecond)
				}
			}
		}(i)
	}

	// usually the main goroutine is long lived ...
	// we will simulate this by reading from a nil channel which blocks forever
	var c chan struct{}

	c <- struct{}{} // blocks forever
}