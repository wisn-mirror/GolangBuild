package main

import (
	"fmt"
	"runtime"
	"sync"
)

var(
	countinc int
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incrCounter()
	go incrCounter()
	wg.Wait()
	fmt.Println(countinc)
}
func incrCounter() {
	defer wg.Done()
	for i:=1;i<=2;i++{
		temp:=countinc
		runtime.Gosched()
		fmt.Println("temp:",temp)
		temp++
		countinc=temp
	}
}