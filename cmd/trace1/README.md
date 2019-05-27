# trace1
main.go contains a sample webapp for demonstrating the execution tracer using GODEBUG environment variable. 

**Compile**:

    go build
       

**Run the program with scheduler tracer:**

    GODEBUG=schedtrace=1000 ./trace1

**OR:**

**Run the program with GC tracer:**

    GODEBUG=gctrace=1 ./trace1
  
**Generate some load on the webapp. For example**:

    hey -n 2000000 -c 100 http://localhost:8181/hello
    
**Scheduler trace output line example**:

    SCHED 241421ms: gomaxprocs=8 idleprocs=0 threads=44 spinningthreads=0 idlethreads=28 runqueue=35 [3 0 1 0 0 0 0 4]
    
Where:

    241421ms - is the number of miliseconds since the program started and this sample has taken.
    gomaxprocs=8 - is the number of OS threads allocated for executing your goroutines (default is the number logical processors).
    idleprocs=0 - the number of idle OS threads allocated for your executing your goroutines.
    threads=44 - total number of OS threads
    spinningthreads=0 - number spinning OS threads (threads which steals runnable goroutines from other "proc" threads in order to keep hardware core busy).
    idlethreads=28 - number of idle OS threads.
    runqueue=35 - number of goroutines in the global queue.
    [3 0 1 0 0 0 0 4] - number of goroutines in the run queues of each of the logical processors.
     
**GC trace output line example:**

    gc 438 @21.942s 1%: 0.061+1.8+0.045 ms clock, 0.48+1.7/2.1/0+0.36 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
    
Where (see also: https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html):
    // General
    gc 438       : The 1404 GC run since the program started
    @21.942s     : Six seconds since the program started
    1%           : Eleven percent of the available CPU so far has been spent in GC
    
    // Wall-Clock
    0.061ms      : STW        : Mark Start       - Write Barrier on
    1.8ms        : Concurrent : Marking
    0.045ms      : STW        : Mark Termination - Write Barrier off and clean up
    
    // CPU Time
    0.48ms       : STW        : Mark Start
    1.7ms        : Concurrent : Mark - Assist Time (GC performed in line with allocation)
    2.1ms        : Concurrent : Mark - Background GC time
    0ms          : Concurrent : Mark - Idle GC time
    0.36ms       : STW        : Mark Term
    
    // Memory
    4MB          : Heap memory in-use before the Marking started
    4MB          : Heap memory in-use after the Marking finished
    1MB          : Heap memory marked as live after the Marking finished
    5MB          : Collection goal for heap memory in-use after Marking finished
     
    // Threads
    8P           : Number of logical processors or threads used to run Goroutines
    L             