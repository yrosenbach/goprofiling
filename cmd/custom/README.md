# remote
'custom' A sample code for writing a custom profile for tracking a custom resource (in this case creation/close of MockFile instances).
For sake of simplicity both 'func main' and the MockFile struct are defined in the same file.

Compile:

    go build   

Run the program:

    ./custom

Open cpu profile:

    go tool pprof http://localhost:6060/debug/pprof/profile   
   
Open cpu profile in a web view:   
   
    go tool pprof -http :1234 http://localhost:6060/debug/pprof/profile 
   
Open memory profile:
   
    go tool pprof http://localhost:6060/debug/pprof/heap 
   
Open memory profile in a web view:   
      
    go tool pprof -http :1234 http://localhost:6060/debug/pprof/heap
    

More available profiles and views are accessible via (browser):

    http://localhost:6060/debug/pprof