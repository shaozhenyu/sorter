package simpleselectsort

import (
	"testing"
)

func TestSimpleSelectSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	SimpleSelectSort(values)

	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("SimpleSelectSort() failed. Got ", values, "Expect 1 2 3 4 5")
	}
}

func TestSimpleSelectSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	SimpleSelectSort(values)

	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("SimpleSelectSort() failed. Got ", values, "Expect 1 2 3 5 5")
	}
}

func TestSimpleSelectSort3(t *testing.T) {
	values := []int{5}
	SimpleSelectSort(values)

	if values[0] != 5 {
		t.Error("SimpleSelectSort() failed. Got ", values, "Expect 5")
	}
}
