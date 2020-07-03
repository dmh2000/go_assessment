1. testing works as though tests are compiled at runtime
    - syntax errors don't show up until a test is run
    - with syntax error , output will not be what you expect 
2. tests should probably have a deferred recovery procedure to catch panics
   - a panic in a test won't kill the whole test sequence, just that test
   -  with a deferred recovery you can output why the panic occurred
3. a simple web server may not be so simple
    - see the rules for handler patterns
    - '/' will handle anything starting with '/'
    - longer patterns are checked first
    - chrome requests favicon.ico even if there is no link in the html
    - should have deferred recoveries to catch panics in handlers