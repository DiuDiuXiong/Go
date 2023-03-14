# Error handling
## defer
`defer` make sure it is called when function end (no matter end by what, e.g. panic, return...)
1. `defer` is stack like, so if there are more than one defer, last occur defer will be called first
2. `defer`'s argument is based on when parameter of defer is created, so will not be affected if parameter of defer changed later after defer line. (If is pointer will still be affected) check [forDefer1](01_defer/defer.go)

## Error handling
1. Check [`writeFileError`](01_defer/defer.go) To see how to handle error
   - `if pathError, ok := err.(*os.PathError); !ok {...}` to do error type assertion
   - if not type we expected, panic, otherwise handle
   - create a self defined error by: `err = errors.New(<err str>)`
2. It is dirty code to handle error separately across the code. See [web_go_naive.go](02_file_listing_server/web_go_naive.go) for example.
Here, use a separation method:
   - Define the actual function of handling request in [a separate file](./02_file_listing_server/file_listing/handler.go)
   - Let this function return error, if everything is great, return nil
   - Then [a wrapper function](./02_file_listing_server/web.go) that takes same input, which will call the separate function, if it returns error, then handle error, and only response error code
   - In this way error is hidden from user
     - file not exist ==> 404
     - permission fail (`sudo chmod 500 forbidden.txt`)to make it only accessible by root ==> 403
     - Other ==> 500

## `panic` & `recover`
1. Stop current function execution (raise error)
2. Execute defer of each layer
3. If didn't see `recover` then exit.

`recover` will only occur in `defer`, it get `panic` value, if unable to handle, re-`panic`. This means when you panic, recover catch that, and do some final handling.

Check [recover.go](03_recover/recover.go) for example. So the rule of thumb is:
- If the error is expected, return error (and handle correspondingly)
- If not use panic
- Use `defer` after open file, some tasks that need clean up no matter what, since panic is reverse order, so it is fine. 
- Since brute throw panic is bad, can use `defer func(...){}()` call function, and within function call, have something like
`if r := recover(); r != nil {...}` And within if, that means functions terminated due to `panic` reason. So can do handling correspondingly.
Use `r.(error)` to get the actual error type. (And can do type assertion from here).
