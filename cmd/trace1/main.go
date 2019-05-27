package main

import (
	"net/http"
)

// export GODEBUG=schedtrace=1000 ; ./trace1
// generate some traffic: hey -n 2000000 -c 100 http://localhost:8181/hello
// explain output of export GODEBUG=schedtrace=1000 ; ./trace1
// SCHED 241421ms: gomaxprocs=8 idleprocs=0 threads=44 spinningthreads=0 idlethreads=28 runqueue=35 [3 0 1 0 0 0 0 4]

// export GODEBUG=gctrace=1 ; ./trace1
// Explain:
// gc 438 @21.942s 1%: 0.061+1.8+0.045 ms clock, 0.48+1.7/2.1/0+0.36 ms cpu, 4->4->1 MB, 5 MB goal, 8 P

func main() {
	http.Handle("/hello", http.HandlerFunc(helloHandler))

	http.ListenAndServe("localhost:8181", http.DefaultServeMux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Gophers!"))
}