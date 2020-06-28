# go-assessment

[A clone of Rebecca Murphey's js-assessment for Go](https://github.com/rmurphey/js-assessment)

I found her js-assessment to be a valuable tool for refreshing my javascript capability, when I
had not written any js for a few months. I think a similar tool for Go will be useful
for the same reason. The tool will be useful for go beginners to go beyond a tutorial and
implement some simple (and a few not-so-simple) Go code without being spoon fed.

This tool is not meant to be a leetcode challenge. Most of the tests and implementations are
straightforward. A few are harder, require a bit of algorithm knowledge or the ability to look
up a solution. The point is to learn or refresh you skill with Go.

The way this tool works is that it relies on Test Driven Development using the [native Go test
framework](https://golang.org/pkg/testing/). The ./app directory
contains a set of test files (xxxx_test.go) and a corresponding skeleton file (xxxx.go) with
the required function definitions but missing the implementations. The \*\_test files
are the specifications for the functions. The goal is to provide the implementations in the
skeleton files to pass the tests. If the tests are not passing, dig into the \*\_test.go files
to figure out what the requirement is. You are done when all tests pass. You should not have to
modify anything in the test files, only the skeleton files (unless the tests have a bug in which
case feel free to report it).

To run the tests you have some choices.

1. Visual Studio Code
   - when set up for Go programming, you should be able to run individual tests from inside the IDE. VS Code gives you clickable links to run individual tests. This would apply to any IDE with similar functionality.
2. From the Shell
   - you can run individual test files using 'go test...'
     - example for the strings test : 'go test -v ./app/strings\*.go ./app/test_util.go '
   - run a shell script that runs all the tests
     - for Linux or Mac, the script is 'test_all.sh'
     - for Windows, the script is 'test_all.cmd'
       - if your tests and other go programs build slowly on Windows, its probably due to Windows Defender or other antivirus
3. From a web browser
   - the tool includes a web server implementation that will run the tests and display the results for you.
     - in the root of the repo, run 'go run \*.go' and connect to the ip:8080 with a browser. You will get a display of all the test results. You can change the port number in './main.go'.
     - after editing files, you can update the test results by refreshing the we page. this will rerun the tests and update the results.
   - the browser tool needs to be connected to the internet because it uses bootstrap from a CDN. If you want to run locally you can download the required bootstrap files and modify './static/index.html' to point to their local copies.
