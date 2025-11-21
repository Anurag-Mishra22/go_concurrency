package main

import (
	"fmt"
	"sync"
)

var actions int
var mu sync.Mutex
var buffer chan struct{}

func consumer() {
	for i := 0; i < 1000; i++ {

		fmt.Printf("[consumer %d] Trying to acquire mutex...\n", i)
		mu.Lock()
		fmt.Printf("[consumer %d] Acquired mutex\n", i)

		actions++
		fmt.Printf("[consumer %d] Incremented actions = %d\n", i, actions)

		fmt.Printf("[consumer %d] Waiting to RECEIVE from channel...\n", i)
		<-buffer
		fmt.Printf("[consumer %d] Received from channel\n", i)

		mu.Unlock()
		fmt.Printf("[consumer %d] Released mutex\n\n", i)
	}
}

func producer() {
	for i := 0; i < 1000; i++ {

		fmt.Printf("[producer %d] Trying to SEND to channel...\n", i)
		buffer <- struct{}{}
		fmt.Printf("[producer %d] Sent to channel\n", i)

		fmt.Printf("[producer %d] Trying to acquire mutex...\n", i)
		mu.Lock()
		fmt.Printf("[producer %d] Acquired mutex\n", i)

		actions++
		fmt.Printf("[producer %d] Incremented actions = %d\n", i, actions)

		mu.Unlock()
		fmt.Printf("[producer %d] Released mutex\n\n", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	buffer = make(chan struct{}, 1)

	go func() {
		defer wg.Done()
		consumer()
	}()

	go func() {
		defer wg.Done()
		producer()
	}()

	wg.Wait()
}
