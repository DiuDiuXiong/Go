package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	i := 0
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

/*
For the function below, we see that speed of c1 and c2 being generated is purly random, whoever have data first will be called.
*/
func nonBlocking() {
	var c1, c2 = generator(), generator()
	for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1:", n)
		case n := <-c2:
			fmt.Println("Received from c2:", n)
		}
	}
}

func useTimerToReturn() {
	var c1, c2 = generator(), generator()
	tm := time.After(10 * time.Second)
	for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1:", n)
		case n := <-c2:
			fmt.Println("Received from c2:", n)
		case <-time.After(800 * time.Millisecond):
			fmt.Println("Cause more than 800 ms")
		case <-tm:
			fmt.Println("Bye")
			return
		}
	}
}

func main() {
	//nonBlocking()
	useTimerToReturn()

}
