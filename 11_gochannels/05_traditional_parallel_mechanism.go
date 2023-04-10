package main

import (
	"fmt"
	"sync"
	"time"
)

// -------------------------- Wait Group -------------------------------------

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

// --------------------------- Mutex ------------------------------------------
// in golang, atomic.<...> can perform some atomic operation

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() { // it is recommended to use system provided operations
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

func main() {
	// useWaitGroup()
	useMutex()
}
