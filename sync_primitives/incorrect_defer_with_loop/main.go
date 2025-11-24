package main

import "sync"

var mutex sync.Mutex
var values []int

func doSomething() {
	for number := 0; number < 10; number++ {
		mutex.Lock()
		defer mutex.Unlock()
		values = append(values, number)
	}
}
