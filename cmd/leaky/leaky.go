package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

var s []string // s is used in order to ensure that GC won't clean the slice created by leakyFunction

func eatMemory() []string {
	s := make([]string, 3)
	for i:= 0; i < 100000000 ; i++{
		s = append(s, "just text ...")
	}

	return s
}

func main() {
	// set CPU profile - START
	f, err := os.Create("./cpu.prof")
	if err != nil {
		fmt.Println("could not create CPU profile: ", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("could not start CPU profile: ", err)
		}
	defer pprof.StopCPUProfile()
	//set CPU profile - END

	// main work - START
	s = eatMemory()
	// main work - END

	//save a memory profile
	f, err = os.Create("./mem.prof")
	if err != nil {
		fmt.Println("could not create memory profile: ", err)
	}

	defer f.Close()
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil { // WriteHeapProfile is syntactic sugar for running: pprof.Lookup("heap").WriteTo(some_file, 0)
		fmt.Println("could not write memory profile: ", err)
	}
}
