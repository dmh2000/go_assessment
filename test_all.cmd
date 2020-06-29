@echo "Flow Control" >test.txt
go test -v ./app/flowControl.go ./app/flowControl_test.go ./app/test_util.go 1>> test.txt 2>&1
@echo "Slices" 1>> test.txt 2>&1
go test -v ./app/slices.go ./app/slices_test.go  ./app/test_util.go 1>> test.txt 2>>&1
@echo "Binary" 1>> test.txt 2>&1
go test -v ./app/binary.go ./app/binary_test.go ./app/test_util.go 1>> test.txt 2>>&1
@echo "Logical Operators" 1>> test.txt 2>&1
go test -v ./app/logicalOperators.go ./app/logicalOperators_test.go ./app/test_util.go 1>> test.txt 2>>&1
@echo "Strings" 1>> test.txt 2>&1
go test -v ./app/strings.go ./app/strings_test.go  ./app/test_util.go 1>> test.txt 2>>&1
@echo "Functions" 1>> test.txt 2>&1
go test -v ./app/functions.go ./app/functions_test.go ./app/test_util.go 1>> test.txt 2>>&1
@echo "Recursion" 1>> test.txt 2>&1
go test -v ./app/recursion.go ./app/recursion_test.go./app/test_util.go 1>> test.txt 2>>&1
@echo "Methods" 1>> test.txt 2>&1
go test -v ./app/methods_test.go ./app/methods_test.go  ./app/test_util.go 1>> test.txt 2>>&1
@echo "Regex" 1>> test.txt 2>&1
go test -v ./app/regex_test.go  ./app/regex_test.go ./app/test_util.go 1>> test.txt 2>>&1
@echo "Async" 1>> test.txt 2>&1
go test -v ./app/async.go ./app/async_test.go ./app/test_util.go 1>> test.txt 2>>&1


