// A sample code for writing a custom profile for tracking a custom resource (in this case creation/close of MockFile instances)
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // will install http handlers on the default router of /net/http (http.DefaultServeMux)
	"runtime/pprof"
	"time"
)

var profilerAddr = "localhost:6062" // in order to access remotely to a production end point you will probably
// need to create an ssh tunnel. For example:
// ssh -L 6062:localhost:6062 <your-host-name>
// which will allow you to execute the pprof tool. For example:
// go tool pprof -http :1234 http://localhost:6062/debug/pprof/heap
// OR:
// go tool pprof http://localhost:6062/debug/pprof/heap

// MockFile is a mock file implementation (in a real program this code should be in a separated package)
var openMockFileProfile = pprof.NewProfile("mockfile.Open")

type MockFile struct {
	path string
}

func Open(path string) *MockFile {
	mockFile := &MockFile{path}

	openMockFileProfile.Add(mockFile, 2) // add the current execution stack to the profile
	return mockFile
}

// Close closes f , the MockFile instance
func (f *MockFile) Close() error {
	openMockFileProfile.Remove(f)
	return nil
}

func main() {
	go func() {
		// http.ListenAndServe may return an error so we would like to print it just in case ...
		fmt.Println(http.ListenAndServe(profilerAddr, nil)) // The handler is nil, which means to use DefaultServeMux
	}()

	for i := 0; i < 1000; i++ {
		go func(i int) {
			path := fmt.Sprintf("/filename-%d", i)
			b := Open(path)
			defer b.Close()

			// Simulate some work with sleep in order to give us some time to inspect the 'mockfile.Open' profile
			time.Sleep(2 * time.Minute)
		}(i)
	}

	// usually the main goroutine is long lived ...
	// we will simulate this by reading from a nil channel which blocks forever
	var c chan struct{}

	c <- struct{}{} // blocks forever
}
