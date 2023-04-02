# acknowledge done

In code example below:
```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
```
We wait for 1ms for all job to be done, this is ineffective. To modify this, we change the following.

```go
package main

import (
	"fmt"
)

func doWorker(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
	done <- true

}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _, worker := range workers {
		close(worker.in)
	}

	// wait for all of them, every worker will have 1 done send before done
	for _, worker := range workers {
		<-worker.done
	}
}

func main() {
	chanDemo()
}
```
We make the `doWorker` function send message out to another channel, indicate task done. And a for loop in
main goroutine to wait for done signal send out from each worker. Also, we packed two channel as a struct. To enable
this to work, we have to close all channels, this will allow for range function to terminate and line proceed to sending message
done to channel.

Or we can use a wait group:
```go
type workerWaitGroup struct {
	in chan int
	wg *sync.WaitGroup
}

func doWorkWaitGroup(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
	wg.Done() // acknowledge one work done
}

func createWorkerWaitGroup(id int, wg *sync.WaitGroup) workerWaitGroup {
	w := workerWaitGroup{
		in: make(chan int),
		wg: wg,
	}
	go doWorkWaitGroup(id, w.in, w.wg)
	return w
}

func chanDemoWaitGroup() {
	var workers [10]workerWaitGroup
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		workers[i] = createWorkerWaitGroup(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _, worker := range workers {
		close(worker.in)
	}
	wg.Wait()
}

func main() {
	chanDemoWaitGroup()
}
```
Here we pass the wait group, it will wait for 10 `.Done()` to be called before proceeding. We can `.Add(x)` and `.Done()` multiple times.
- `var wg sync.WaitGroup` to define
- `wg.Add(x)` make the counting += 1
- `wg.Done()` make the counting -= 1
- `wg.Wait()` wait the counting == 0

One modification for code above will be instead let worker have `WaitGroup`, let it have a function that whenever called call `wg.Done()`