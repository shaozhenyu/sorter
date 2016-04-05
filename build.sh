#!bin/bash
go build algorithms/qsort
go test algorithms/qsort

go build algorithms/bubblesort
go test algorithms/bubblesort

go build algorithms/mergesort
go test algorithms/mergesort

go build algorithms/heapsort
go test algorithms/heapsort

go install algorithms/qsort
go install algorithms/bubblesort
go install algorithms/mergesort
go install algorithms/heapsort

go build sorter
go install sorter
