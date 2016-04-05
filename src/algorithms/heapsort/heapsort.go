package heapsort

func createHeap(values []int, i, n int) {
	k := 2*i + 1

	for k < n {
		if k < n-1 && values[k] < values[k+1] {
			k++
		}

		if values[k] > values[i] {
			values[k], values[i] = values[i], values[k]
			i = k
			k = 2*i + 1
		} else {
			break
		}
	}
}

func heapSort(values []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		createHeap(values, i, n)
	}

	for i := n; i > 0; i-- {
		values[i], values[0] = values[0], values[i]
		createHeap(values, 0, i)
	}
}

func HeapSort(values []int) {
	heapSort(values, len(values)-1)
}
