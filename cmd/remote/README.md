# remote
'remote' is a sample program for exposing profiles in the default webserver.

Compile:

    go build   

Run the program:

    ./remote

Open cpu profile:

    go tool pprof http://localhost:6062/debug/pprof/profile   
   
Open cpu profile in a web view:   
   
    go tool pprof -http :1234 http://localhost:6062/debug/pprof/profile 
   
Open memory profile:
   
    go tool pprof http://localhost:6062/debug/pprof/heap 
   
Open memory profile in a web view:   
      
    go tool pprof -http :1234 http://localhost:6062/debug/pprof/heap
    

More available profiles and views are accessible via (browser):

    http://localhost:6062/debug/pprof