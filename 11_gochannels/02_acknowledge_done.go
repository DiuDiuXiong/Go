package main

import (
	"fmt"
	"sync"
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

// wait group

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

/*
func main() {
	chanDemoWaitGroup()
}
*/
