package main

import (
	"fmt"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Second/2)
			out <- i
			i++
		}
	}()
	return out
}

func workerg(id int, ch chan int) {
	for chtemp := range ch {
		time.Sleep(time.Second*2)
		fmt.Printf(" id:%d, result%d \n", id, chtemp)
	}
}

func generWorker(id int) chan int {
	out := make(chan int)
	go workerg(id, out)
	return out
}

func main() {
	var c1, c2 = generator(), generator()
	gw1 := generWorker(1)
	temp:=0
	var values  []int
	timechan:=time.After(time.Second*4)
 	for {
		var worker chan int
		if len(values) > 0 {
			worker = gw1
			temp = values[0]
		}
		select {
		case temp = <-c1:
			values = append(values, temp)
		case temp = <-c2:
			values = append(values, temp)
		case worker <- temp:
			values = values[1:]
		case <-timechan:
			fmt.Println(values)
			timechan=time.After(time.Second*4)
		}
	}

}
