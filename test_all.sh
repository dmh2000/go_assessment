#!/bin/bash


echo "Operators"
go test -v ./app/operators*.go ./app/test_util.go ./app/app_types.go  

echo "Flow Control" 
go test -v ./app/flowControl*.go ./app/test_util.go ./app/app_types.go

echo "Strings" 
go test -v ./app/strings*.go  ./app/test_util.go ./app/app_types.go

echo "Slices" 
go test -v ./app/slices*.go  ./app/test_util.go ./app/app_types.go

echo "Binary" 
go test -v ./app/binary*.go ./app/test_util.go ./app/app_types.go

echo "Functions" 
go test -v ./app/functions*.go ./app/test_util.go ./app/app_types.go

echo "Methods" 
go test -v ./app/method*.go  ./app/test_util.go ./app/app_types.go

echo "Collections" 
go test -v ./app/collections*.go  ./app/test_util.go ./app/app_types.go

echo "Recursion" 
go test -v ./app/recursion*.go ./app/test_util.go ./app/app_types.go

echo "Regexp" 
go test -v ./app/regex*.go  ./app/test_util.go ./app/app_types.go

echo "Async" 
go test -v ./app/async*.go  ./app/test_util.go ./app/app_types.go

