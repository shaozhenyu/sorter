#!bin/bash

filelist=`ls src/algorithms`
for file in $filelist
do
	go build algorithms/$file
	go test algorithms/$file
done

for file in $filelist
do
	go install algorithms/$file
done

go build sorter
go install sorter
