package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	go func() {
		ch <- 1
	}()

	value := 0
	value += <-ch
	value += <-ch
	fmt.Println(value)
}

// Don’t share memory to communicate — communicate to share memory.
// Instead of multiple goroutines updating the same shared variable,
// send values through channels and let one goroutine own the variable.
