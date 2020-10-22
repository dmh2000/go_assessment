@rem start the test web server
@rem on Windows, 'go run' can conflict with some antivirus, so
@rem this script will build the .exe first, then run it.
@rem that seems to work around some of the antivirus tools
@rem build step
go build .
@rem to specify a port, change this to 'goassessment.exe -port 8081'
goassessment.exe 
