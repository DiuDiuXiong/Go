# go channel
Channel is a way different go routines communicate with each other.

- To create a channel for typeX, use `ch := make(chan <typeX>)`
- To push a data to channel, use `<-` operator: ```var x typeX; c <- x;```
- To pull a data from channel, use `<-` operator: ```var y typeX; y <- c;```

With those in mind, the following code block main goroutine push data to channel `c` and another go routine pull from it.

```go
func goChannelDemo() {
    c := make(chan int)
        go func(){
            for {
                newInt := <- c
                fmt.Println(newInt)
            }
        }()
    c <- 1
    c <- 2
    time.Sleep(time.Millisecond)
}
```
This function have few key points:
- goroutine defined start from line 3 subscribe to channel `c` before main goroutine start to push to it, since push to a channel
without subscriber is a deadlock error in golang.
- we let time sleep for a while so it has time to pull data from channel and IO print.
```go
func chanDemoError() {
c := make(chan int) // a channel contain int
c <- 1              // push data to channel
c <- 2

n := <-c // receive data from channel
fmt.Println(n)

}
```
This will cause error since push to unsubscribed channel.

Channel can also be used as function parameter:
```go
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
```
This will output `Worker x received x` for x range from 0 to 9 orderless.

Channel can also be used as output. There are three possibility:
- `chan <type>`: Return a general type channel
- `chan<- <type> `: Send only channel, this type of channel can only be pushed data into it but cannot pull data from it once returned.
- `<-chan <type>`: Receive only channel, this type of channel can only be pulled data from but cannot push data into it.
For example:
```go
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
```

However, subscribers many not ready when receivers are ready to receive data, so channel buffer enable data to be buffered before being received.
Using `make(chan <type>, <num>)` will create space enough for `num` amount of `<type>`. And as long as buffer still have vacancies, send to it is fine.

Some special cases:
- Try to send/receive when there are only one alive routine. (This will cause infinite waiting for input or pushing something that noone ever gonna receive)

E.g.
```go
func waitToSend() { // this function will create an error
	c := make(chan int)
	c <- 1
	go func() {
		fmt.Println(<-c)
	}()
}
```
Following function will cause the error that try to use channel when only one routine being active.
- Wait To Read
```go
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
```
In this function, it will print 0-9 with 1 second interval between each. Since `fmt.Println(<-c)` will wait until there are something new
to receive. The function below will cause infinite loop. Since main co-routine will wait for input infinitely, while there
are more than one goroutine so no error will not pop out.
```go
func infiniteWait() {
	c := make(chan int)
	go func() {
		for {
			fmt.Println("Infinite Loop.")
		}
	}()
	fmt.Println(<-c)
}
```

However, following two code will generate error:
```go
func infiniteWait() {
	c := make(chan int)
	go func() {
		for i:=0; i<100000; i++{
			fmt.Println("Infinite Loop.")
		}
	}()
	fmt.Println(<-c)
	c <- 1
}
```
Keep one of `fmt.Println(<-c)` or `c<-1` will cause error since when there are only one goroutine left and channel operation not cleared, error will occur.
- close chanel `close(<channel>)`
  - Sender side decide when to close
  - Receiver side have two ways to determine whether channel is closed already:
    - `val, ok := <-<channel>` or `for val := range <channel>`
For example:
```go
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
```
- Go channel is based on CSP model.