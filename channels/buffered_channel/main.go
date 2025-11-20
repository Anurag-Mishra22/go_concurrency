package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 50
	ch <- 50
	fmt.Println("End of Program")
}
