package main

import (
	"fmt"
	"time"
)
const count=45
var fibs [count]uint64
func fibbonacciMermory(i int)(res uint64 )  {
	if i>=0&&fibs[i]!=0{
		res= fibs[i]
		return
	}
	if i<=1{
		res= 1
	}else{
		res= fibbonacciMermory(i-1)+ fibbonacciMermory(i-2)
	}
	fibs[i]=res
	return res
}

func  exFoq(ch chan int, i int)  {
	fmt.Println("result:",i, fibbonacciMermory(i))
	ch<-i
}


func main() {
	ch := make(chan int, count)
	starttime := time.Now().UnixNano()
	for i := 0; i < count; i++ {
		go exFoq(ch, i)
		//exFoq(ch, i)
	}
	for i := 0; i < count; i++ {
		fmt.Print(<-ch," ")
	}
	endtime := time.Now().UnixNano()
	timesum := (endtime - starttime) / 1e6
	fmt.Println("用时毫秒：", timesum)
}
