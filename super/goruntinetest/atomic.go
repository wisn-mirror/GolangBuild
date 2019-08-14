package main

import (
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock sync.Mutex
}
func (atomicInt *atomicInt)increment(){
	atomicInt.value++
}
func (atomicInt *atomicInt) getValue() int {
	atomicInt.lock.Lock()
	defer  atomicInt.lock.Unlock()
	return atomicInt.value
}

func main() {
	var testvalue atomicInt
	testvalue.increment()
	go func() {
		testvalue.increment()
		testvalue.increment()
		testvalue.increment()
		testvalue.increment()
	}()
	go func() {
		testvalue.increment()
		testvalue.increment()
		testvalue.increment()
		testvalue.increment()
	}()
	time.Sleep(time.Microsecond*2)
	println(testvalue.getValue())

}