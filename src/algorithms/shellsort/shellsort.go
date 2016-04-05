package shellsort

func shellSort(values []int, len int) {
	gap := len
	for gap > 1 {
		gap = gap / 2
		flag := 1
		for flag != 0 {
			flag = 0
			for i := 0; i < len-gap; i++ {
				j := i + gap

				if values[i] > values[j] {
					values[i], values[j] = values[j], values[i]
					flag = 1
				}
			}
		}
	}
}

func ShellSort(values []int) {
	shellSort(values, len(values))
}
