package main

import (
	"fmt"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			//time.Sleep(time.Second)
			out <- i
			i++
		}
	}()
	return out
}

func workerg(id int, ch chan int) {
	for chtemp := range ch {
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
	gw2 := generWorker(2)
	for {
		select {
		case n := <-c1:
			gw1 <- n
		case n := <-c2:
			gw2 <- n
		}
	}

}
