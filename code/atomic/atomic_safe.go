package main

import (
	"fmt"
	"sync"
	"time"
)

type ins struct {
	value int
	lock sync.Mutex
}

func (in *ins )increment(step int )  {
	in.lock.Lock()
	defer in.lock.Unlock()
	in.value=in.value+step
}
func (in *ins ) get() int  {
	in.lock.Lock()
	defer  in.lock.Unlock()
	return in.value
}

func main() {
	var data ins
	data.increment(2)
	go func() {
		data.increment(3)
	}()
	time.Sleep(time.Second)
	fmt.Println(data.get())
}
