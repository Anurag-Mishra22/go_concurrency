package main

import "sync"

var buffer chan int

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			buffer = make(chan int)
		}()
	}

	wg.Wait()
}

// go run -race main.go
