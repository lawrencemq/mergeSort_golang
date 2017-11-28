package main


import (
	"fmt"
	"time"
	"./mergeSort"
)


func main(){
	a := []int{2,1,3,4,7,6,5}
	beforeMergeTime := time.Now()
	mergeSort.Merge(a)
	afterMergeTime := time.Since(beforeMergeTime)
	fmt.Printf("%d\n", afterMergeTime)
	fmt.Printf("%v\n", a)

	b := []int{2,1,4,6,7,5,8}
	beforePMergeTime := time.Now()
	mergeSort.ParallelMerge(b)
	afterPMergeTime := time.Since(beforePMergeTime)
	fmt.Printf("%d\n", afterPMergeTime)
	fmt.Printf("%v\n", b)
}



