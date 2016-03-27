package mergesort

func merge(values []int, first, mid, last int, temp []int) {
	i := first
	j := mid + 1
	k := 0

	for i <= mid && j <= last {
		if values[i] <= values[j] {
			temp[k] = values[i]
			k++
			i++
		} else {
			temp[k] = values[j]
			k++
			j++
		}
	}

	for i <= mid {
		temp[k] = values[i]
		k++
		i++
	}

	for j <= last {
		temp[k] = values[j]
		k++
		j++
	}

	for i = 0; i < k; i++ {
		values[first+i] = temp[i]
	}
}

func mergeSort(values []int, first, last int, temp []int) {
	if first < last {
		mid := (first + last) / 2
		mergeSort(values, first, mid, temp)
		mergeSort(values, mid+1, last, temp)
		merge(values, first, mid, last, temp)
	}
}

func MergeSort(values []int) {
	temp := make([]int, len(values))
	mergeSort(values, 0, len(values)-1, temp)
}
