package main

import (
	"fmt"
	"runtime"
	"time"
)


func worker(chantemp chan int, id int) {
	for {
		//错误写法，当channel 关闭的时候 不断读取0
		//fmt.Printf("id :%d,result:%d \n", id, <-chantemp)
			result, ok := <-chantemp
			if !ok {
				break
			}
			fmt.Printf("id :%d,result:%c \n", id, result)

		/*for result := range chantemp {
			fmt.Printf("id :%d,result:%c \n", id, result)
		}*/
	}
}


func closeChan()  {
	closechan := make(chan int, 4)
	go worker(closechan, 9)
	for i := 0; i < 4; i++ {
		closechan <- 'A' + i
	}
	close(closechan)
	time.Sleep(time.Second * 2)
}
func main() {
	closeChan()
}


func demo2() {
	var chanlist [10]chan int
	for i := 0; i < len(chanlist); i++ {
		chanlist[i] = make(chan int)
		go worker(chanlist[i], i)
	}
	for i := 0; i < len(chanlist); i++ {
		chanlist[i] <- 'a' + i
	}
	for i := 0; i < len(chanlist); i++ {
		chanlist[i] <- 'A' + i
	}
	time.Sleep(time.Second)
}


func bufferchanel()  {
	chanbuf := make(chan int, 4)
	go worker(chanbuf, 9)
	for i := 0; i < 4; i++ {
		chanbuf <- 'a' + i
	}
	time.Sleep(time.Second)
}




func demo1() {
	ch := make(chan int)
	go func() {
		for {
			result := <-ch
			time.Sleep(time.Second)
			fmt.Println("result", result)
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	time.Sleep(time.Second * 2)
}

func test2() {
	slice := make([]int, 10)
	for i := 0; i < 10; i++ {
		go func(temp int) {
			for{
				//Found 1 data race(s)  需要通过channel解决
				slice[temp]++
				//手动交出控制权
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(slice)
}

func test1() {
	for i := 0; i < 100; i++ {
		go func(temp int) {
			for{
				fmt.Println("temp：", temp)
			}
		}(i)
	}
	time.Sleep(time.Second)
}

