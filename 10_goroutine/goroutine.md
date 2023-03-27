# Goroutine

Goroutine is a way to achieve parallel programming (but not with thread of OS).

1. Use `go func(?){...}(?)` to execute functions parallel. Each func statement after `go ...` will be registered as a goroutine. E.g.:
```go
func parallelPrint() {
    for i := 0; i < 1000; i++ {
        go func(i int) {
            for {
                fmt.Printf("Hello from "+"goroutine %d\n", i)
            }
        }(i)
    }
    time.Sleep(time.Millisecond)
}
```
This will cause 1000 different goroutine, each responsible for printing `goroutine i` run parallel.

2. Note that example above have the following phenomenons:
    - Each goroutine started by `go ...` are parallel with each other
    - Some of them get executed some of them not
    - let `main` function sleep for 1ms is the key to get output. This is because:
`main` will keep going alone side with goroutine, if program doesn't sleep, after main finish, will terminate directly
which will also shut down all goroutines, this is why without sleep, following function will have no effect.
```go
func noEffect() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from "+"goroutine %d\n", i)
			}
		}(i)
	}
}
```

3. Note that each goroutine (code block within `go` statement):
   - Is a lightweight version of coroutine
   - Unlike thread, goroutine defines where to opt out itself (have few opt out possibilities). 
Which means only at certain line of code the go scheduler can stop the code and pass the CPU resources to other goroutines.
Whereas threads can break at any line of code, since OS will determine switching itself.
(Note that for newer version of go, the issue that if a goroutine block have no switch point causing CPU being blocked is modified).
So that even when a task have no switch point, it can also be halted by go scheduler
  - For golang, compiler explainer & go embedded scheduler together determine context switching between goroutines
  - Multiple goroutine can run in same/different thread, this is determined by go scheduler.

4. Possible goroutine switch point includes:
   - I/O, select
   - channel
   - waiting lock
   - call functions
   - `runtime.Gosched()`

5. `rumtime.Gosched()`

Rather than other points of code where go scheduler might switch, `runtime.Gosched()` tells goroutine that at this line of code,
I am happy to be stopped and switch context. However:
- This does not guarantee a context switch, scheduler may skip this switch point and let goroutine keep going
- Even you have no switch point define above, go scheduler will still try to switch if one routine take too long to avoid dead block issue.

An example:
```go
func tryGosched() {
	var a, b [1000]int
	var aContinue, bContinue = true, true
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for aContinue {
				a[i]++
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	aContinue = false

	for i := 0; i < 1000; i++ {
		go func(i int) {
			for bContinue {
				b[i]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	bContinue = false
	fmt.Println(a, b)

}
```
Code below cause a wth many fields 0 while others with huge number, while every field of b is non-zero. This is because 
forced switch by goroutine is very redundant, that will only cause switch after a long time, and will switch between first
few registered task, this cause many go routine about modify a list not executed at all in 1ms. While for b, since we manually
define switching point, all routines have some chance to execute tasks with sufficient time (1ms).

6. Note that we pass a function call (code block) to `go` to generate a goroutine. And at the time we pass it, it causes a closure.
So following code will generate out of range error:
```go
func raceError() {
	var a [1000]int
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				a[i]++
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
```
This is because, i is not passed as a parameter (hence not copy). This forms a closure. And `main` will keep going alongside with goroutines.
Then at the time for ends, it reaches the termination condition that i == 1000, and start sleeping, goroutines still keep going.
And a[1000] cause out of range error. Even without this error, race condition that writing while reading is dangerous. In this code
when `main` finish sleeping and execute `fmt.Println(a)`, goroutines are still writing to it. This can also cause race condition.

Use: `go run -race <go file>.go` to check race points during code.

7. Other language support for coroutine
- C++: Boost.Coroutine
- Java: NOPE
- Python 3.5-: yield
-Python 3.5+: async def (define a coroutine)

As long as we call function as go func()...  can make it coroutine, unlike python have to define it is a async function during definition

8. `main` is also a goroutine
9. Different goroutines can communicate with each other, this is done through go channels.