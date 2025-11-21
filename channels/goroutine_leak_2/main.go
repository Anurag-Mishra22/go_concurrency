package main

// First-response-wins strategy
func request() int {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func() {
			ch <- i // 4 goroutines will be blocked
		}()
	}

	return <-ch
}

// You start 5 goroutines, each trying to send i into an unbuffered channel.

// The main goroutine waits for 1 value from the channel.

// After one send succeeds, the function returns immediately.
