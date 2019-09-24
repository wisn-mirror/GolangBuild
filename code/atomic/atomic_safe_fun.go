package main

import (
	"fmt"
	"sync"
	"time"
)

type ins2 struct {
	value int
	lock  sync.Mutex
}

func (in *ins2) increment(step int) {
	func() {
		in.lock.Lock()
		defer in.lock.Unlock()
		in.value = in.value + step
	}()
}
func (in *ins2) get() int {
	in.lock.Lock()
	defer in.lock.Unlock()
	return in.value
}

func main() {
	var data ins2
	data.increment(2)
	go func() {
		data.increment(3)
	}()
	time.Sleep(time.Second)
	fmt.Println(data.get())
}
