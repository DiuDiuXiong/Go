package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 01 generator
/**
- generate a channel that return as a function output
- within the function, start a go routine that sends message or receive message from the channel
*/
func msgGen() <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("Messag %d", i)
			fmt.Println("Next")
		}
	}()
	return c
}

func msgGenUse() {
	m := msgGen()
	for {
		fmt.Println(<-m)
	}
}

// 02 Combination
// Start multiple Goroutine, each of them receive from one channel, and output all of them to one identical channel

func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1 //receive data from c1 and send it to c
		}
	}()
	go func() {
		for {
			c <- <-c2 //receive data from c2 and send it to c
		}
	}()
	return c
}

func useFanIn() {
	c1, c2 := msgGen(), msgGen()
	f := fanIn(c1, c2)
	for {
		fmt.Println(<-f)
	}
}

// 03 Select
// `select` as advised before

func selectCombo(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

/*
Comparison between select/fanIn
1. If there are multiple goroutine required (unsure how much, then we should use fanIn method)
2. Otherwise, we can just predefine each branch in select
3. For undefined amount of inputs, see example below
*/

// 坑
// This version have an issue that only the last channel ch will send data to c, this is why
// In each for _,ch {...} clause, it will first create each go routine, and starts to execute for {c<-<-ch} clause only
// when fmt.Println(<-c) actually be called. But at that time, ch already change to the last value of chs, which is last channel
func fanIns坑(chs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ch := range chs {
		go func() {
			for {
				c <- <-ch
			}
		}()
	}
	return c
}

func useFanIns坑() {
	c1, c2, c3 := msgGen(), msgGen(), msgGen()
	c := fanIns坑(c1, c2, c3)
	for {
		fmt.Println(<-c)
	}
}

// This leads to the idea that we should copy iteration variable if goroutine created within iteration execute depend on iteration variable
func fanIns不坑(chs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ch := range chs {
		chCopy := ch
		go func() {
			for {
				c <- <-chCopy
			}
		}()
	}
	return c
}

// an alternative way to copy is by function param copy property
/*
...
for _, ch:= range chs {
	go func (in <-chan string) {
		for {
			c <- <-in
		}
	}(ch)
}

*/

func useFanIns不坑() {
	c1, c2, c3 := msgGen(), msgGen(), msgGen()
	c := fanIns不坑(c1, c2, c3)
	for {
		fmt.Println(<-c)
	}
}

/*
func main() {
	//msgGenUse()
	//useFanIn()
	//useFanIns坑() (only one branch of 0-inf)
	useFanIns不坑() // multiple branch of 0-inf
}
*/
