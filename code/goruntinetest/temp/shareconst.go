package temp

import (
	"fmt"
	"runtime"
	"sync"
)

var countss int

func countAdd( lock *sync.Mutex)  {
	lock.Lock()
	countss++
	lock.Unlock()

}
func main() {
	lock:= &sync.Mutex{}
	for i:=0;i<10;i++{
		go countAdd(lock)
		//time.Sleep(time.Second)
	}
	for{
		lock.Lock()
		c:=countss
		fmt.Println("result",c)
		lock.Unlock()
		runtime.Gosched()
		if c>10{
			break
		}
		//time.Sleep(time.Second)
	}
}