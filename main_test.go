package main

import (
	"testing"
	"math/rand"
	"./mergeSort"
)


const smallBenchmarkTest = 15
const mediumBenchmarkTest = 2000
const largeBenchmarkTest = 2000000

func createRandomSliceOfInts(size int) []int {
	a := make([]int, 0)
	for i :=0 ; i < size; i++ {
		a = append(a, rand.Int())
	}
	return a
}

func areNumbersInAcendingOrder(numbers []int) bool {
	if len(numbers) < 2 {
		return true
	}
	for j := 1; j < len(numbers); j++ {
		if numbers[j-1] > numbers[j] {
			return false
		}
	}
	return true
}

func TestMergeSort(t *testing.T) {

	for i:= 0; i <= 7; i++ {
		a:= createRandomSliceOfInts(i)
		for ;!areNumbersInAcendingOrder(a); { // ensuring that the numbers given weren't already sorted
			a = createRandomSliceOfInts(i)
		}
		mergeSort.Merge(a)
		if len(a) != i{
			t.Fail()
		}
		if !areNumbersInAcendingOrder(a){
			t.Fail()
		}
	}
}

func BenchmarkMergeSort__SmallSort(b *testing.B) {
	a := createRandomSliceOfInts(smallBenchmarkTest)
	for i := 0; i < b.N; i++ {
		mergeSort.Merge(a)
	}
}

func BenchmarkMergeSort__MediumSort(b *testing.B) {
	a := createRandomSliceOfInts(mediumBenchmarkTest)
	for i := 0; i < b.N; i++ {
		mergeSort.Merge(a)
	}
}

func BenchmarkMergeSort__LargeSort(b *testing.B) {
	a := createRandomSliceOfInts(largeBenchmarkTest)
	for i := 0; i < b.N; i++ {
		mergeSort.Merge(a)
	}
}




func TestParallelMergeSort(t *testing.T) {

	for i:= 0; i <= 7; i++ {
		a:= createRandomSliceOfInts(i)
		for ;!areNumbersInAcendingOrder(a); { // ensuring that the numbers given weren't already sorted
			a = createRandomSliceOfInts(i)
		}

		mergeSort.ParallelMerge(a)
		if len(a) != i{
			t.Fail()
		}
		if !areNumbersInAcendingOrder(a){
			t.Fail()
		}
	}
}

func BenchmarkParallelMergeSort__SmallSort(b *testing.B) {
	a := createRandomSliceOfInts(smallBenchmarkTest)
	for i := 0; i < b.N; i++ {
		mergeSort.ParallelMerge(a)
	}
}

func BenchmarkParallelMergeSort__MediumSort(b *testing.B) {
	a := createRandomSliceOfInts(mediumBenchmarkTest)
	for i := 0; i < b.N; i++ {
		mergeSort.ParallelMerge(a)
	}
}

func BenchmarkParallelMergeSort__LargeSort(b *testing.B) {
	a := createRandomSliceOfInts(largeBenchmarkTest)
	for i := 0; i < b.N; i++ {
		mergeSort.ParallelMerge(a)
	}
}
