# custom
'custom' A sample code for writing a custom profile for tracking a custom resource (in this case creation/close of MockFile instances).
For sake of simplicity both 'func main' and the MockFile struct are defined in the same file.

Compile:

    go build   

Run the program:

    ./custom
    
See in your browser all the available profiles including 'mockfile.Open':

    http://localhost:6062/debug/pprof

Open mockfile.Open profile:

    go tool pprof http://localhost:6062/debug/pprof/mockfile.Open   
   
Open mockfile.Open profile in a web view:   
   
    go tool pprof -http :1234 http://localhost:6062/debug/pprof/mockfile.Open 