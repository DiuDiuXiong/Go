package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGenerator(serviceName string) <-chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			c <- fmt.Sprintf("Service %s iteration %d", serviceName, i)
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			i++
		}
	}()
	return c
}

// If get string, boolean should be true, false otherwise
// Since there is a default section, thus if no data from c, it will not be blocked
func nonBlockingWait(c <-chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

func useNonBlocking() {
	m1 := msgGenerator("Service 1")
	for {
		time.Sleep(1000 * time.Millisecond)
		if m, ok := nonBlockingWait(m1); ok {
			fmt.Println(m)
		} else {
			fmt.Println("No message from service 1")
		}
	}
}

// Replace the return empty string with a time.After
func timeoutWait(c <-chan string, t time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(t * time.Millisecond):
		return "", false
	}
}

func useTimeoutWait() {
	m1 := msgGenerator("Service 1")
	for {
		if m, ok := timeoutWait(m1, 1000); ok {
			fmt.Println(m)
		} else {
			fmt.Println("No message for service 1")
		}
	}
}

// elegant exiting
func msgGenWithDone(name string, done <-chan struct{}) <-chan string { // struct{} contains less data
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(1500 * time.Millisecond):
				c <- fmt.Sprintf("Services %s: message %d", name, i)
			case <-done:
				fmt.Println("Cleaning up")
				return
			}

			i++
		}
	}()
	return c
}

func useMsgGenWithDone() {
	done := make(chan struct{})
	m1 := msgGenWithDone("Service 1", done)
	for i := 0; i < 5; i++ {
		if m, ok := timeoutWait(m1, 1000); ok {
			fmt.Println(m)
		} else {
			fmt.Println("Timeout")
		}
	}
	done <- struct{}{}           // first is a definition for struct: struct{}, second {} is for initialization
	time.Sleep(time.Millisecond) // let go routine have enough time for cleaning
}

// More elegant existing
func msgGenWithDone1(name string, done chan struct{}) <-chan string { // struct{} contains less data
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(1500 * time.Millisecond):
				c <- fmt.Sprintf("Services %s: message %d", name, i)
			case <-done:
				fmt.Println("Cleaning up")
				done <- struct{}{}
				return
			}

			i++
		}
	}()
	return c
}

func useMsgGenWithDone1() {
	done := make(chan struct{})
	m1 := msgGenWithDone1("Service 1", done)
	for i := 0; i < 5; i++ {
		if m, ok := timeoutWait(m1, 1000); ok {
			fmt.Println(m)
		} else {
			fmt.Println("Timeout")
		}
	}
	done <- struct{}{}
	<-done // wait for finish signal from working go routines
}

func main() {
	//useNonBlocking()
	//useTimeoutWait()
	//useMsgGenWithDone()
	useMsgGenWithDone1()
}
