@echo "Flow Control" >test.txt
go test -v ./app/flowControl*.go ./app/test_util.go 1> test.txt 2>&1
@echo "Slices" >>test.txt
go test -v ./app/slices*.go  ./app/test_util.go 1> test.txt 2>&1
@echo "Binary" >>test.txt
go test -v ./app/binary*.go ./app/test_util.go 1> test.txt 2>&1
@echo "Logical Operators" >>test.txt
go test -v ./app/logicalOperators*.go ./app/test_util.go 1> test.txt 2>&1
@echo "Strings" >>test.txt
go test -v ./app/strings*.go  ./app/test_util.go 1> test.txt 2>&1
@echo "Functions" >>test.txt
go test -v ./app/functions*.go ./app/test_util.go 1> test.txt 2>&1
@echo "Recursion" >>test.txt
go test -v ./app/recursion*.go ./app/test_util.go 1> test.txt 2>&1
@echo "Methods" >>test.txt
go test -v ./app/method*.go  ./app/test_util.go 1> test.txt 2>&1
@echo "Regex" >>test.txt
go test -v ./app/regex*.go  ./app/test_util.go 1> test.txt 2>&1
@echo "Async" >>test.txt

type test.txt
exit 0