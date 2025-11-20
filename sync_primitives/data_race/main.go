package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	text := ""

	go func() {
		defer wg.Done()
		text = "hello world"
	}()

	go func() {
		defer wg.Done()
		fmt.Println(text)
	}()

	wg.Wait()

}

// go run -race main.go
