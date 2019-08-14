package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ch = make(chan int)
var chSendOnly chan<- int = ch
var chRecvOnly <-chan int = ch
var wgg sync.WaitGroup
func main() {
	//timer2 :=time.NewTicker(time.Second)
	court:=make(chan int )
	wgg.Add(2)
	go player("wisn",court)
	go player("test",court)
	court<-1
	wgg.Wait()
}
func player(name string, court chan int) {
	defer wgg.Done()
	for{
		ball,ok:=<-court
		if !ok{
			fmt.Printf("player %s won\n",name)
			return ;
		}
		rand.Seed(time.Now().UnixNano())
		n:=rand.Intn(10)
		if n%3==0{
			fmt.Printf("player %s miss %d\n",name,n)
			close(court)
			return
		}
		fmt.Printf("player %s hit %d\n",name,ball)
		ball++
		court<-ball

	}
}
