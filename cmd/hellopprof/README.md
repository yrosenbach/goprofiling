# hellopprof
Getting started pprof guide which introduces some useful commands.

We will use cpu.prof created in the leaky sample app (cmd/leaky).
This guide assumes that cpu.prof and mem.prof exist in the current directory. 


**pprof interactive shell**
```
go tool pprof ./cpu.prof
```

After entering the interactive shell you will see output similar to:   
```
Type: cpu
Time: May 21, 2019 at 1:22pm (IDT)
Duration: 16.59s, Total samples = 21.70s (130.84%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) 
``` 
     
Where:
- Type: cpu - is the is the type of the profile we opened (cpu in this case).
- Time: May 21, 2019 at 1:22pm (IDT) - is the time where this profile was recorded.
- Duration: 16.59s, Total samples = 21.70s (130.84%)
  - 16.59s: is the duration (seconds) which this profile was recrded.
  - Total samples = 21.70s (130.84%): is the total CPU time and CPU percentage which this profile is counted for.
  

**top**:
<br />
Type 'top' and you will see output similar to (topN where N is an integer will display the N heaviest stacktrace nodes):
```
(pprof) top
Showing nodes accounting for 18950ms, 87.33% of 21700ms total
Dropped 36 nodes (cum <= 108.50ms)
Showing top 10 nodes out of 57
      flat  flat%   sum%        cum   cum%
    4170ms 19.22% 19.22%     4170ms 19.22%  runtime.memmove
    3550ms 16.36% 35.58%     5610ms 25.85%  runtime.scanobject
    3190ms 14.70% 50.28%     3190ms 14.70%  runtime.usleep
    2400ms 11.06% 61.34%     3520ms 16.22%  runtime.findObject
    1310ms  6.04% 67.37%     3270ms 15.07%  runtime.wbBufFlush1
    1280ms  5.90% 73.27%    11700ms 53.92%  main.eatMemory
     970ms  4.47% 77.74%     1130ms  5.21%  runtime.spanOf
     760ms  3.50% 81.24%     4450ms 20.51%  runtime.bulkBarrierPreWriteSrcOnly
     690ms  3.18% 84.42%      690ms  3.18%  runtime.nanotime
     630ms  2.90% 87.33%      660ms  3.04%  runtime.heapBitsSetType
(pprof) 
```  

**flat vs cum:**
<br />
By default the nodes listed by the output of the top command are listed by the 'flat' value of the node which is the direct CPU usage contributed by this line. 
<br />
You can type 'cum' and then 'top' in order to get the top nodes sorted by their cumulative value which the CPU usage due to execution paths where this node was on the stacktrace:
```$xslt
(pprof) cum
(pprof) top
Showing nodes accounting for 5.62s, 25.90% of 21.70s total
Dropped 36 nodes (cum <= 0.11s)
Showing top 10 nodes out of 57
      flat  flat%   sum%        cum   cum%
         0     0%     0%     13.90s 64.06%  runtime.systemstack
     1.28s  5.90%  5.90%     11.70s 53.92%  main.eatMemory
         0     0%  5.90%     11.70s 53.92%  main.main
         0     0%  5.90%     11.70s 53.92%  runtime.main
         0     0%  5.90%      9.96s 45.90%  runtime.growslice
         0     0%  5.90%      9.22s 42.49%  runtime.gcBgMarkWorker.func2
     0.03s  0.14%  6.04%      9.22s 42.49%  runtime.gcDrain
         0     0%  6.04%      6.23s 28.71%  runtime.gcBgMarkWorker
     3.55s 16.36% 22.40%      5.61s 25.85%  runtime.scanobject
     0.76s  3.50% 25.90%      4.45s 20.51%  runtime.bulkBarrierPreWriteSrcOnly
(pprof) 
``` 

**list** 
<br />
The 'list' command allows you to inspect a detailed usage within a stacktrace node which is a function. You will be able to see the function's code lines and the usage of each line.
<br />
The usage is: **list [regex]** where regex is a regular expression which match a function name.
<br /> Using a loose regex may result with more than one result.
<br /> Now let's inspect the CPU usage of **_main.eatMemory_**:
```$xslt
(pprof) list eatMemory
Total: 21.70s
ROUTINE ======================== main.eatMemory in /Users/yrosenbach/go/src/github.com/yrosenbach/goprofiling/cmd/leaky/leaky.go
     1.28s     11.70s (flat, cum) 53.92% of Total
         .          .     10:var s []string // s is used in order to ensure that GC won't clean the slice created by eatMemory
         .          .     11:
         .          .     12:func eatMemory() []string {
         .          .     13:	s := make([]string, 3)
         .          .     14:	for i:= 0; i < 100000000 ; i++{
     1.28s     11.70s     15:		s = append(s, "just text ...")
         .          .     16:	}
         .          .     17:
         .          .     18:	return s
         .          .     19:}
         .          .     20:
(pprof) 
```
Explanation:
1. **_list eatMemory_** - has matched a single node which is the function main.eatMemory.
2. The code lines of main.eatMemory are expanded and we can see both the flat and cum values per line (when applicable).
   This is a CPU profile so we can see both the total (flat, cum) and by line values.
   We can see that line 15: **_s = append(s, "just text ...")_** is where the CPU has been (mostly) used in main.eatMemory.