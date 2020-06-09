#!/bin/bash

rm test.txt
echo "@Flow Control" >>test.txt
go test -v ./app/flowControl.go ./app/flowControl_test.go ./app/test_util.go &>>test.txt
echo "@Slices" >>test.txt
go test -v ./app/slices.go ./app/slices_test.go ./app/test_util.go &>>test.txt
echo "@Logical Operators" >>test.txt
go test -v ./app/logicalOperators.go ./app/logicalOperators_test.go ./app/test_util.go &>>test.txt
echo "@Numbers" >>test.txt
go test -v ./app/numbers.go ./app/numbers_test.go ./app/test_util.go &>>test.txt
echo "@Strings" >>test.txt
echo "@Functions" >>test.txt
go test -v ./app/functions.go ./app/functions_test.go ./app/test_util.go &>>test.txt
echo "@Recursion" >>test.txt
go test -v ./app/recursion.go ./app/recursion_test.go ./app/test_util.go &>>test.txt
echo "@Methods" >>test.txt
echo "@Regex" >>test.txt
echo "@Count" >>test.txt
echo "@Async" >>test.txt

exit 0