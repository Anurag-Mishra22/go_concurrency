package main

import "fmt"

func isClosed(ch chan int) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func main() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(isClosed(ch))
	close(ch)
	fmt.Println(isClosed(ch))
}
