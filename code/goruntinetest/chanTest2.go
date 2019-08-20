package main

import "fmt"

func printCount(c chan int )  {
	for{
		cresult:=<-c
		if cresult==-1{
			break
		}
		fmt.Println("result",cresult)
	}
	c<--1

}
func main() {
	c:=make(chan int)
	go printCount(c)
	for i:=0;i<10;i++{
		c<-i
	}
	c<--1
	fmt.Println("end1")

	<-c
	fmt.Println("end")

}