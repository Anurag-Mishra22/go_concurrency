package main

import "sync"

var mutex sync.Mutex

func operation() {}

func doSomething() {
	mutex.Lock()
	operation()
	mutex.Unlock()

	// some long operation

	mutex.Lock()
	operation()
	mutex.Unlock()
}
