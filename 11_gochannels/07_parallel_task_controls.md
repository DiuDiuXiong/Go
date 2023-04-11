# Parallel task control

Includes:
- Non-blocking waiting
- Time out mechanism
- Task halting, exiting
- Responsive exit

## Non-blocking wait:
Use `select` and add a default branch that will just return empty. This way the `select` statement won't wait forever

```go
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
			fmt.Println("No message from services2")
		}
	}
}
```

## Timeout mechanism
Instead of returning false + empty string right away, wait a time by `<-time.After(duration)` and put that in a select.
```go
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
```

## Existing mechanism
For code above, we have to let main exit to shut down other go routines, which is not efficient, instead, those
go routines can have one signal channel as input, and whenever receive information from that channel, indicates shut
down go routine. In this way we have control of when to end thr routine. Code below generates a message every 1.5sec until receive return signal
```go
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
```

## Elegant existing
However, in code above we use time.Sleep to give time for cleaning, but we are not sure if after 1ms cleaning will be finish.
Instead, we can let main thread wait for done then exist.
```go
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
```