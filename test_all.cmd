@echo "Flow Control" >results.txt
go test -v ./app/flowControl.go ./app/flowControl_test.go ./app/test_util.go ./app/app_types.go  1>> results.txt 2>&1
@echo "Logical Operators" ./app/app_types.go  1>> results.txt 2>&1
go test -v ./app/logicalOperators.go ./app/logicalOperators_test.go ./app/test_util.go ./app/app_types.go  1>> results.txt 2>>&1
@echo "Strings" ./app/app_types.go  1>> results.txt 2>&1
go test -v ./app/strings.go ./app/strings_test.go  ./app/test_util.go ./app/app_types.go  1>> results.txt 2>>&1
@echo "Slices" ./app/app_types.go  1>> results.txt 2>&1
go test -v ./app/slices.go ./app/slices_test.go  ./app/test_util.go ./app/app_types.go  1>> results.txt 2>>&1
@echo "Binary" ./app/app_types.go  1>> results.txt 2>&1
go test -v ./app/binary.go ./app/binary_test.go ./app/test_util.go ./app/app_types.go  1>> results.txt 2>>&1
@echo "Functions" ./app/app_types.go  1>> results.txt 2>&1
go test -v ./app/functions.go ./app/functions_test.go ./app/test_util.go ./app/app_types.go  1>> results.txt 2>>&1
@echo "Recursion" ./app/app_types.go  1>> results.txt 2>&1
go test -v ./app/recursion.go ./app/recursion_test.go./app/test_util.go ./app/app_types.go  1>> results.txt 2>>&1
@echo "Methods" ./app/app_types.go  1>> results.txt 2>&1
go test -v ./app/methods_test.go ./app/methods_test.go  ./app/test_util.go ./app/app_types.go  1>> results.txt 2>>&1
@echo "Regexp" ./app/app_types.go  1>> results.txt 2>&1
go test -v ./app/regex_test.go  ./app/regex_test.go ./app/test_util.go ./app/app_types.go  1>> results.txt 2>>&1
@echo "Async" ./app/app_types.go  1>> results.txt 2>&1
go test -v ./app/async.go ./app/async_test.go ./app/test_util.go ./app/app_types.go  1>> results.txt 2>>&1


