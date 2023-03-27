package main

import (
	"fmt"
	"runtime"
	"time"
)

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

func noEffect() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from "+"goroutine %d\n", i)
			}
		}(i)
	}
}

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

func raceError() {
	var a [1000]int
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				a[i]++
				runtime.Gosched() // handle out control // much even
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func main() {
	// parallelPrint()
	// noEffect()
	// tryGosched()
	raceError()
}
