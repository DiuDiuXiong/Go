# Traditional parallel mechanism

We can also use traditional way despite `select` for parallel programming, for example:
- `sync.WaitGroup` as defined before
- `Mutex`: A mutex acts as a lock that allows only one thread or process to access the shared resource at a time
- `Cond`: conditional variable

## sync.WaitGroup
- `<sync.WairGroup>.Add(x)` add x counter
- `<sync.WaitGroup>.Done()` Add(-1)
- `<sync.WaitGroup>.Wait()` Wait until count == 0
```go
func useWaitGroup() {
	fmt.Println("Use WaitGroup Start ------------------------------------")
	wg := &sync.WaitGroup{}
	wg.Add(200)
	doWork := func(repeat int, id int, wg *sync.WaitGroup) {
		for i := 0; i < repeat; i++ {
			fmt.Println(id)
			wg.Done()
		}
	}
	for i := 0; i < 10; i++ {
		go doWork(20, i, wg)
	}
	wg.Wait()
	fmt.Println("Use WaitGroup Stop  ------------------------------------")
}
```

## sync.Mutex
```go
type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() { // if we uncomment these two lines, it will always add to 201
	//a.lock.Lock()
	//defer a.lock.Unlock()
	a.value++

}

// note by adding lock/unlock to a function, we create an atomic function

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func useMutex() {
	var a atomicInt
	a.increment()
	for i := 0; i < 200; i++ {
		go func() {
			a.increment()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(a.get())
}
```
1. If we uncomment these two lines, it will always add to 201, otherwise `useMutex` will add to something < 201, since it will parallel write.
2. Note use of `defer a.lock.Unlock()` right after `Lock()`
```go
var lo sync.Mutex
lo.Lock()
defer lo.Unlock()
```
3. By adding a lock&unlock to function, we create an atomic functions

## Cond
This is based on Mutex, which is based on lock and unlock.
```go
package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var cv *sync.Cond = sync.NewCond(&mu)

func main() {
    wg.Add(2)

    // Goroutine 1 waits for the condition to become true
    go func() {
        defer wg.Done()
        mu.Lock()
        defer mu.Unlock()

        fmt.Println("Goroutine 1 waiting...")
        cv.Wait()
        fmt.Println("Goroutine 1 done!")
    }()

    // Goroutine 2 signals the condition after a short delay
    go func() {
        defer wg.Done()
        mu.Lock()
        defer mu.Unlock()

        fmt.Println("Goroutine 2 setting condition...")
        cv.Signal()
    }()

    wg.Wait()
}
```

- `var cv *sync.Cond = sync.NewCond(&<implements Lock() and Unlock()>)`
- `cv.Wait()`
- `cv.Signal()`

In this example, we create a `sync.Mutex` variable mu and a `sync.Cond` variable cv just like in the previous example.

We then create two goroutines. The first goroutine waits for the condition to become true by calling `cv.Wait()` while holding the lock on mu. The second goroutine signals the condition by calling `cv.Signal()` while holding the lock on mu.

When the program runs, Goroutine 1 will print "Goroutine 1 waiting..." and then block on `cv.Wait()`. Goroutine 2 will print "Goroutine 2 setting condition..." and then signal the condition by calling `cv.Signal()`. Goroutine 1 will then unblock and print "Goroutine 1 done!" before the program exits.

Note that in this example, we don't actually use a boolean variable to represent the condition. Instead, we simply use the fact that `cv.Wait()` will block until cv.Signal() is called to signal the condition. However, in a more realistic example, you would typically use a boolean variable to represent the condition that you are waiting for.