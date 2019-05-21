# leaky
Instrumenting an application to generate cpu and memory pprof files.

leaky.go is a simple program which call _**fuc eatMemory**_ which allocates memory.
A code which generates cpu & memory profile is included in **_func main_**.
cpu and memory pprof files will be written as cpu.prof and mem.prof respectively.

Compile:

    go build
    
Or if you want to see an escape analysis:

    go build -gcflags="-m -m"    

Run the program:

    ./leaky

Open cpu profile:

    go tool pprof ./cpu.prof 
   
Open cpu profile in a web view:   
   
    go tool pprof -http :1234 ./cpu.prof 
   
Open memory profile:
   
    go tool pprof ./mem.prof
   
Open memory profile in a web view:   
      
    go tool pprof -http :1234 ./mem.prof 
  