package main

import (
	"fmt"
	"time"
)

func goChannelDemo() {
	c := make(chan int)
	go func() {
		for {
			newInt := <-c
			fmt.Println(newInt)
		}
	}()

	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func chanDemoError() {
	c := make(chan int) // a channel contain int
	c <- 1              // push data to channel
	c <- 2

	n := <-c // receive data from channel
	fmt.Println(n)

}

func takeChannelAsInput(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %d\n", id, <-c)
	}
}

func useChannelAsInput() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go takeChannelAsInput(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		channels[i] <- i
	}
	time.Sleep(time.Millisecond)

}

func receiveOnly() chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	return c
}

func sendOnly() <-chan int {
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		c <- i
	}
	return c
}

func testReceiveSend() {
	receiveOnlyChannel, sendOnlyChannel := receiveOnly(), sendOnly()
	receiveOnlyChannel <- 1
	// fmt.Println(<- receiveOnlyChannel); error pull from receive only channel

	go func() {
		for n := range sendOnlyChannel {
			fmt.Println(n)
		}
	}()
	// sendOnlyChannel <- 1 error, push to send only channel
	time.Sleep(time.Millisecond)
}

func channelBuffer() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	// c <- 4 error, since
	go func() {

		for i := 0; i < 3; i++ {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(time.Millisecond)
	c <- 4 // this is also fine since buffer cleared
	c <- 5
	c <- 6
	c <- 7 // this is not fine since buffer full
}

func waitToReceive() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			c <- i
		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

}

func waitToSend() { // this function will create an error
	c := make(chan int)
	c <- 1
	go func() {
		fmt.Println(<-c)
	}()
}

func infiniteWait() {
	c := make(chan int)
	go func() {
		for {
			fmt.Println("Infinite Loop.")
		}
	}()
	fmt.Println(<-c)
}

func closeChannelAndDetection() {
	c1, c2 := make(chan int), make(chan int)
	go func() {

		for {
			val, ok := <-c1
			if !ok {
				fmt.Println("Finish")
				break
			} else {
				fmt.Println(val)
			}
		}
	}()

	c1 <- 1
	c1 <- 2
	close(c1)

	go func() {
		for val := range c2 {
			fmt.Println(val)
		}
		fmt.Println("Finished")
	}()
	c2 <- 1
	c2 <- 2
	close(c2)

	time.Sleep(time.Millisecond)
}

/*
func main() {
	// goChannelDemo()
	// chanDemoError()
	// useChannelAsInput()
	// testReceiveSend()
	// channelBuffer()
	// waitToReceive()
	// waitToSend()
	// infiniteWait()
	closeChannelAndDetection()

}
*/
