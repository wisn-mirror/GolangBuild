package temp

import "fmt"

func main() {
	//sendOnCloseChanel()
	readOnCloseChan()

}

func readOnCloseChan() {
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
	//多试一次
	for i := 0; i < cap(ch)+1; i++ {
		result, ok := <-ch
		fmt.Println(result, ok)
	}
}

func sendOnCloseChanel() {
	ch := make(chan int)
	close(ch)
	ch <- 1
}
