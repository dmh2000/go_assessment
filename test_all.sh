#!/bin/bash

echo "" >$GITHUB_OUTPUT

echo "Operators"
go test -v ./app/operators*.go ./app/test_util.go ./app/app_types.go  2>&1>> $GITHUB_OUTPUT

echo "Flow Control" 
go test -v ./app/flowControl*.go ./app/test_util.go ./app/app_types.go 2>&1>> $GITHUB_OUTPUT

echo "Strings" 
go test -v ./app/strings*.go  ./app/test_util.go ./app/app_types.go 2>&1>> $GITHUB_OUTPUT

echo "Slices" 
go test -v ./app/slices*.go  ./app/test_util.go ./app/app_types.go 2>&1>> $GITHUB_OUTPUT

echo "Binary" 
go test -v ./app/binary*.go ./app/test_util.go ./app/app_types.go 2>&1>> $GITHUB_OUTPUT

echo "Functions" 
go test -v ./app/functions*.go ./app/test_util.go ./app/app_types.go 2>&1>> $GITHUB_OUTPUT

echo "Methods" 
go test -v ./app/method*.go  ./app/test_util.go ./app/app_types.go 2>&1>> $GITHUB_OUTPUT

echo "Collections" 
go test -v ./app/collections*.go  ./app/test_util.go ./app/app_types.go 2>&1>> $GITHUB_OUTPUT

echo "Recursion" 
go test -v ./app/recursion*.go ./app/test_util.go ./app/app_types.go 2>&1>> $GITHUB_OUTPUT

echo "Regexp" 
go test -v ./app/regex*.go  ./app/test_util.go ./app/app_types.go 2>&1>> $GITHUB_OUTPUT

echo "Async" 
go test -v ./app/async*.go  ./app/test_util.go ./app/app_types.go 2>&1>> $GITHUB_OUTPUT

exit 0