package temp

import (
	"fmt"
	"time"
)

func main() {
	//test1()
	test2()

}
func test2() {
	//每秒一次打点
	ticker:=time.NewTicker(time.Second)
	//计时器，4秒后
	stopTick:=time.NewTimer(time.Second*4)
	fmt.Println("开始")
	var i int
	for{
		select{
		case <-ticker.C:
			i++
			fmt.Println("打点",i)
		case <-stopTick.C:
			goto StopTip
		}
	}
	StopTip:
		fmt.Println("stop tip")
}
func test1()  {
	exit:=make(chan int)
	fmt.Println("开始")

	time.AfterFunc(time.Second, func() {
		fmt.Println("print")
		exit<-1
	})
	<-exit
	fmt.Println("结束了")
}