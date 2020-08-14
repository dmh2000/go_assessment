#!/bin/bash

echo "Flow Control"
echo "Flow Control" >results.txt
go test -v ./app/flowControl*.go ./app/test_util.go ./app/app_types.go &>>results.txt

echo "Logical Operators" 
echo "Logical Operators" >>results.txt
go test -v ./app/logicalOperators*.go ./app/test_util.go ./app/app_types.go &>>results.txt

echo "Strings" 
echo "Strings" >>results.txt
go test -v ./app/strings*.go  ./app/test_util.go ./app/app_types.go &>>results.txt

echo "Slices" 
echo "Slices" >>results.txt
go test -v ./app/slices*.go  ./app/test_util.go ./app/app_types.go &>>results.txt

echo "Binary" 
echo "Binary" >>results.txt
go test -v ./app/binary*.go ./app/test_util.go ./app/app_types.go &>>results.txt

echo "Functions" 
echo "Functions" >>results.txt
go test -v ./app/functions*.go ./app/test_util.go ./app/app_types.go &>>results.txt

echo "Recursion"
echo "Recursion" >>results.txt
go test -v ./app/recursion*.go ./app/test_util.go ./app/app_types.go &>>results.txt

echo "Methods"
echo "Methods" >>results.txt
go test -v ./app/method*.go  ./app/test_util.go ./app/app_types.go &>>results.txt

echo "Regexp"
echo "Regexp" >>results.txt
go test -v ./app/regex*.go  ./app/test_util.go ./app/app_types.go &>>results.txt

echo Async
echo "Async" >>results.txt
go test -v ./app/async*.go  ./app/test_util.go ./app/app_types.go &>>results.txt

exit 0