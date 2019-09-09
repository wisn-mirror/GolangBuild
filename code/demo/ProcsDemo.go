package main

import (
	"fmt"
	"runtime"
	"time"
)


func doSomeD(start, end int, control chan int) {
	var temp int
	for ;start<end;start++{
		//todo add
		temp=temp+start
	}
	control <- temp
}

func doAllD() {
	cpu := runtime.NumCPU()
	fmt.Println("cup num", cpu)
	chD := make(chan int, cpu)
	temp:=1935553358
	size:=temp/cpu
	fmt.Println("start size ",size)
	starttime:=time.Now().UnixNano()
	for i := 0; i <= cpu; i++ {
		end:=(i+1)*size
		if end>temp{
			end=temp
		}
		go doSomeD(size*i, end, chD)
		fmt.Println("slice:start",size*i," end," ,end)
	}
	var ResultD int
	for j := 0; j <= cpu; j++ {
		resultTemp:=<-chD
		ResultD=ResultD+resultTemp
	}
	endtime:=time.Now().UnixNano()
	timesum:=(endtime-starttime)/1e6
	fmt.Println("Result:",ResultD,"用时毫秒：",timesum)

}
func main() {
	//fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	//fmt.Printf("时间戳（纳秒）：%v;\n",time.Now().UnixNano())
	//fmt.Printf("时间戳（毫秒）：%v;\n",time.Now().UnixNano() / 1e6)
	//fmt.Printf("时间戳（纳秒转换为秒）：%v;\n",time.Now().UnixNano() / 1e9)

	doAllD()
}
