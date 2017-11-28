package mergeSort


func merge(arr []int, l, m, r int){

	n1 := m - l + 1
	n2 :=  r - m

	/* create temp slices */
	L := make([]int, n1)
	R := make([]int, n2)

	/* Copy data to temp arrays L[] and R[] */
	for i := 0; i < n1; i++{
		L[i] = arr[l + i]
	}
	for j := 0; j < n2; j++{
		R[j] = arr[m + 1+ j]
	}


	/* Merge the temp arrays back into arr[l..r] */
	i := 0 // Start first subarray
	j := 0 // Start second subarray
	k := l // Start of merged subarray
	for ;i < n1 && j < n2; {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		}else{
			arr[k] = R[j]
			j++
		}
		k++
	}

	/* Copy the remaining elements of L[] */
	for ; i < n1 ; {
		arr[k] = L[i]
		i++
		k++
	}

	/* Copy the remaining elements of R[] */
	for ; j < n2 ; {
		arr[k] = R[j]
		j++
		k++
	}
}


/* l is for left index and r is right index of the sub-array of arr to be sorted */
func mergeSort(arr []int, l, r int){
	if l < r {
		// Same as (l+r)/2, but avoids overflow for
		// large l and h
		m := l+(r-l)/2

		// Sort first and second halves
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)

		merge(arr, l, m, r)
	}
}

func Merge(arr []int){
	mergeSort(arr, 0, len(arr)-1)
}





/* l is for left index and r is right index of the sub-array of arr to be sorted */
func parallelMergeSort(arr []int, l, r int, signalChannel chan bool){
	if l < r {
		// Same as (l+r)/2, but avoids overflow for
		// large l and h
		m := l+(r-l)/2

		// Making communications channels and sorting first and second halves
		firstChannel := make(chan bool)
		secondChannel := make(chan bool)
		defer close(firstChannel)
		defer close(secondChannel)
		go parallelMergeSort(arr, l, m, firstChannel)
		go parallelMergeSort(arr, m+1, r, secondChannel)

		// Ensuring all went well and that both finish before a merge
		childrenMergedSuccessfully := (<-firstChannel) && (<- secondChannel)
		if ! childrenMergedSuccessfully {
			signalChannel <- false
		}else{
			merge(arr, l, m, r)
		}
	}

	// signaling complete
	signalChannel <- true
}

func ParallelMerge(arr []int){
	toReturnChannel := make(chan bool)
	defer close(toReturnChannel)
	go parallelMergeSort(arr, 0, len(arr)-1, toReturnChannel)
	if ! <- toReturnChannel {
		panic("Something failed!")
	}
}
