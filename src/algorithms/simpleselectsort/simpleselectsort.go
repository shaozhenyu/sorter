package simpleselectsort

func simpleSelectSort(values []int, lens int) {
	for i := 0; i < lens-1; i++ {
		min := i
		for j := i + 1; j < lens; j++ {
			if values[j] < values[min] {
				min = j
			}
		}

		if min > i {
			values[i], values[min] = values[min], values[i]
		}

	}

}

func SimpleSelectSort(values []int) {
	simpleSelectSort(values, len(values))
}
