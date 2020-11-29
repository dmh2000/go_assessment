## Lessons Learned From A Browser Based Go Test Driver

The following is not news to experienced Go developers but it was to me.

I hadn't written a Go web server app before, so going in it seemed simple. Then
things started to go wrong. I learned that the Go runtime supports a form a recovery
that many other runtimes don't.

1. Defer-Recover is the best thing going for a server app
   - [read more about it here](https://blog.golang.org/defer-panic-and-recover)
   - something most conventional runtimes don't support
   - Go allows recovery from panics, and even syntax errors in tests
   - in a node.js server app,for example, the recovery strategy is typically to let the process die and have it auto restarted
   - in a Go program, you can use defer-recover to catch panics and possibly continue
   - for example, if one web handler is crashing because some resource is missing, you can do a deferred recovery, log something and potentially disable the handler, while letting the rest of the program continue
   - or, just do a comprehensive log and restart
2. Go native testing works as though tests are compiled at runtime
   - syntax errors don't show up until a test is run
   - with a syntax error, output will not be what you expect
   - using a deferred recovery you can continue and log information even after a syntax error in a test
3. all tests should probably have a deferred recovery procedure to catch panics
   - a panic in a test won't kill the whole test sequence, just that test. it will be marked Fail
   - with a deferred recovery you can log why the panic occurred instead of a cryptic fail message
4. a simple web server may not be so simple
   - see the rules for handler patterns in servemux doc
   - '/' will handle anything starting with '/'
   - longer patterns are checked first
   - chrome requests favicon.ico even if there is no link in the html
   - should have deferred recoveries to catch panics in handlers
