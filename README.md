mergeSort_golang
===============

Testing an in-place merge sort in GoLang in both the traditional, single-threaded way as well as the multi-threaded way. 


Build and Run
-------------
Clone/fork the repo and run the following commands to build and run the application:
```
go build
./golangMergeSort
```

This will run two merge sorts. The first is single-threaded, and the second is multi-threaded. The screen will print out the sorted arrays as well as the time (in nanoseconds) it took to complete the sort. E.g., 
```
4968
[1 2 3 4 5 6 7]
52439
[1 2 4 5 6 7 8]
```


Testing and Benchmarking
------------

To run all tests to ensure both merge sorts work correctly, run:
```
go test
```

To benchmark all merge sorts against a small (15), medium (2,000), and large (2,000,000) array, run:
```
go test -bench=.
```


Results
-------
From what I've seen, the parallel merge is around 10X slower than the single-threaded. This is likely due to the overhead of creating extra go routines and channels when splitting an array into tiny pieces. 

This could be improved in many ways, including having a pool of go routines to act on instead of creating two new go routines and channels on every split. This will be left for additional work. 