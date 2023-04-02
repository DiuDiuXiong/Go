# select
`select` is a control structure in Go that is used for communication and synchronization between Goroutines. It is used to monitor multiple channels and blocks until one or more of the channels are ready to send or receive data.

The select statement lets a Goroutine wait on multiple communication operations. It blocks until one of its cases can proceed. If multiple cases can proceed, one is chosen randomly. The syntax of the select statement is as follows:
```go
select {
case channel1 <- message1:
    // do something
case message2 := <-channel2:
    // do something
default:
    // do something else
}
```
In this example, select is monitoring two channels: channel1 and channel2. If channel1 is ready to receive data, it will send message1 on that channel. If channel2 is ready to send data, it will receive message2 from that channel. If neither channel is ready, the default case will be executed.

The select statement is often used in conjunction with Goroutines to enable concurrent communication between them. It allows a Goroutine to block on multiple channels simultaneously, which can be useful in many different scenarios, such as load balancing, event handling, and distributed systems.

In select we can use Nil channel, `<-nil_val_chan` will be false in select statement.

Also, `time.After(duration*time.Second)` will send a message to tm after duration we defined, this enables us:
```go
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
```
After 10 seconds, tm will start to receive a message and terminate. Note that tm will be selected for sure since:

The select statement blocks until one of its cases can proceed, and if multiple cases are ready to proceed, one is chosen randomly. Therefore, if one channel is generating messages very quickly, it will be chosen more frequently by the select statement. However, if the other channels also have data ready, they will eventually be chosen as well.

Note for third case, it means that between two select, if time is more than 800ms, a timeout statement will be notified.
This is also a good point to do some debuging like buffer zone, awaited data amount etc...

- `time.After(x*time.Second)` send 1 message after x seconds
- `time.Tick(x*time.Second)` send 1 message every x seconds

Note if a channel is unbuffered, then a goroutine should be ready to receive from it before sending any data to it otherwise error.
If it is a buffered channel, in select statement, the if buffer channel full, then send to channel in select won't be true but won't cause error as well.
```go
package main

import "fmt"

func main() {
    ch := make(chan int, 1)
    ch <- 1

    select {
    case ch <- 2:
        fmt.Println("Sent")
    default:
        fmt.Println("Channel is full")
    }

    fmt.Println(<-ch)
}
```
Program below will print out channel is full all the time.