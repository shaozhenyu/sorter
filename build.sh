#!bin/bash
go build algorithms/qsort
go test algorithms/qsort

go build algorithms/bubblesort
go test algorithms/bubblesort

go install algorithms/qsort
go install algorithms/bubblesort

go build sorter
go install sorter
