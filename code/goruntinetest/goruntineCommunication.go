package main

import (
	"fmt"
)

type community struct {
	send chan int
	done chan int
}
func workerc( community community, id int) {
	for {
		for result := range community.send {
			fmt.Printf("id :%d,result:%c \n", id, result)
			community.done<-1
		}
	}
}
func goruntinec() {
	var chanlist [10]community
	for i := 0; i < len(chanlist); i++ {
		chanlist[i] =community{
			send:make(chan int),
			done:make(chan int),
		}
		go workerc(chanlist[i], i)
	}
	for i,com:=range chanlist{
		com.send <- 'a' + i
	}
	for _,com:=range chanlist{
		<-com.done
	}
	for i,com:=range  chanlist{
		com.send <- 'A' + i
	}
	for _,com:=range chanlist{
		<-com.done
	}
}


func main() {
	goruntinec()
}






