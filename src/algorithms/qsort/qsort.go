package qsort

func quickSort(values []int, left, right int) {
	//p := left
	i, j := left, right

	if left < right {
		temp := values[left]
		for i < j {
			for values[j] >= temp && i < j {
				j--
			}

			values[i] = values[j]

			for values[i] <= temp && i < j {
				i++
			}

			values[j] = values[i]
		}
		values[i] = temp
		quickSort(values, left, i-1)
		quickSort(values, j+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
