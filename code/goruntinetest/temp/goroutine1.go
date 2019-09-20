package temp

import "time"

func main() {
	ch:=make(chan int)
	go func() {
		println("start ")
		time.Sleep(time.Microsecond*100)
		println("end")
		ch<-0
	}()
	println("wait")
	<-ch
	println("done")
}
