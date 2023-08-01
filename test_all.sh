#!/bin/bash

rm results.txt
echo "Operators" 
echo "Operators" >>results.txt
go test -v ./app/operators*.go ./app/test_util.go ./app/app_types.go 2>&1>>results.txt

echo "Flow Control"
echo "Flow Control" >>results.txt
go test -v ./app/flowControl*.go ./app/test_util.go ./app/app_types.go 2>&1>>results.txt

echo "Strings" 
echo "Strings" >>results.txt
go test -v ./app/strings*.go  ./app/test_util.go ./app/app_types.go 2>&1>>results.txt

echo "Slices" 
echo "Slices" >>results.txt
go test -v ./app/slices*.go  ./app/test_util.go ./app/app_types.go 2>&1>>results.txt

echo "Binary" 
echo "Binary" >>results.txt
go test -v ./app/binary*.go ./app/test_util.go ./app/app_types.go 2>&1>>results.txt

echo "Functions" 
echo "Functions" >>results.txt
go test -v ./app/functions*.go ./app/test_util.go ./app/app_types.go 2>&1>>results.txt

echo "Methods"
echo "Methods" >>results.txt
go test -v ./app/method*.go  ./app/test_util.go ./app/app_types.go 2>&1>>results.txt

echo "Collections"
echo "Collections" >>results.txt
go test -v ./app/collections*.go  ./app/test_util.go ./app/app_types.go 2>&1>>results.txt

echo "Recursion"
echo "Recursion" >>results.txt
go test -v ./app/recursion*.go ./app/test_util.go ./app/app_types.go 2>&1>>results.txt

echo "Regexp"
echo "Regexp" >>results.txt
go test -v ./app/regex*.go  ./app/test_util.go ./app/app_types.go 2>&1>>results.txt

echo Async
echo "Async" >>results.txt
go test -v ./app/async*.go  ./app/test_util.go ./app/app_types.go 2>&1>>results.txt

exit 0