# hello
Micro level profiling example of profiling via a benchmark.

hello.go is a simple program and hello_test.go is a benchmark for the HelloToWorld function which isn't efficient in term of memory usage.

Compile:

    go build

Execute benchmark with cpu and memory profiles:

    go test -cpuprofile cpu.prof -memprofile mem.prof -run none -bench . -benchtime 3s -benchmem

Open cpu profile:

    go tool pprof ./cpu.prof 
   
Open cpu profile in a web view:   
   
    go tool pprof -http :1234 ./cpu.prof 
   
Open memory profile:
   
    go tool pprof ./mem.prof
   
Open memory profile in a web view:   
      
    go tool pprof -http :1234 ./mem.prof 
  