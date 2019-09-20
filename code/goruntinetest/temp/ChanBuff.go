package temp

import "fmt"

func main() {
	ch:=make(chan int,4)
	fmt.Println(len(ch))
	ch<-1
	ch<-2
	ch<-3
	ch<-4
	fmt.Println(len(ch))
}
