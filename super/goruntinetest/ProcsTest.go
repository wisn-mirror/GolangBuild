package main

import (
	"fmt"
	"runtime"
)

type Vector []float64

func doSome(start, end int, control chan int) {
	for ;start<end;start++{
		//todo add
	}
	control <- 1
}

func doAll(colleciton Vector) {
	cpu := runtime.NumCPU()
	fmt.Println("cup num", cpu)
	ch := make(chan int, cpu)
	fmt.Println("start")
	for i := 0; i < cpu; i++ {
		go doSome(i*cap(colleciton), (i+1)*cap(colleciton), ch)
	}
	for j := 0; j < cpu; j++ {
		<-ch
	}

}
func main() {

}
