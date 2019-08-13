package main

import (
	"fmt"
	"time"
)

func main() {
	ch:=make(chan  int)
	go func() {
		for i:=5;i>=0;i--{
			ch<-i
			time.Sleep(time.Second)
		}
	}()
	for result:=range ch{
		fmt.Println(" result",result)
		if result==0{
			break
		}
	}
}