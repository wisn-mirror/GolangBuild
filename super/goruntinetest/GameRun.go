package main

import (
	"fmt"
	"sync"
	"time"
)
var wgRun sync.WaitGroup
func main() {
	runch:=make(chan int)
	wgRun.Add(1)
	go Runner(runch)
	runch<-1
	wgRun.Wait()

}
func Runner(baton chan int) {
	var newRunner int
	getBaton:=<-baton
	fmt.Printf("runner %d get Baton\n",getBaton)
	if getBaton!=5 {
		newRunner=getBaton+1
		fmt.Printf("creater runner %d wait\n",newRunner)
		go Runner(baton)
	}
	time.Sleep(time.Second)
	if getBaton==5{
		wgRun.Done()
		fmt.Printf(" runner is Over\n")

	}
	fmt.Printf(" runner %d put baton\n",getBaton)
	baton<-newRunner
}