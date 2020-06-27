#!/bin/bash

echo "Flow Control"
echo "Flow Control" >test.txt
go test -v ./app/flowControl*.go ./app/test_util.go &>>test.txt

echo "Slices" 
echo "Slices" >>test.txt
go test -v ./app/slices*.go  ./app/test_util.go &>>test.txt

echo "Binary" 
echo "Binary" >>test.txt
go test -v ./app/binary*.go ./app/test_util.go &>>test.txt

echo "Logical Operators" 
echo "Logical Operators" >>test.txt
go test -v ./app/logicalOperators*.go ./app/test_util.go &>>test.txt

echo "Strings" 
echo "Strings" >>test.txt
go test -v ./app/strings*.go  ./app/test_util.go &>>test.txt

echo "Functions" 
echo "Functions" >>test.txt
go test -v ./app/functions*.go ./app/test_util.go &>>test.txt

echo "Recursion"
echo "Recursion" >>test.txt
go test -v ./app/recursion*.go ./app/test_util.go &>>test.txt

echo "Methods"
echo "Methods" >>test.txt
go test -v ./app/method*.go  ./app/test_util.go &>>test.txt

echo "Regex"
echo "Regex" >>test.txt
go test -v ./app/regex*.go  ./app/test_util.go &>>test.txt

echo Async
echo "Async" >>test.txt
go test -v ./app/async*.go  ./app/test_util.go &>>test.txt

cat test.txt
exit 0