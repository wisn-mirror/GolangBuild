###goroutine

Don‘t communicate by sharing memory ;share memory by communicating .

不要通过共享内存来通信，通过通信来共享内存



非抢占式的

go run -race goruntine1.go可以查看数据抢占点

可能的切换点(可能存在的点)：

- i/o  select
- 函数调用
- channel
- runtime.Gosched()
- 等待锁

如果goruntine 超过cpu核心数，一般为最多开的线程为核心数

```go
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
 // 等待输出完成
	time.Sleep(time.Second * 2)
}
```

生产者消费者

```go
func main() {
	var chanlist [10]chan int
	for i := 0; i < len(chanlist); i++ {
		chanlist[i] = make(chan int)
		go func(chantemp chan int, id int) {
			for {
				fmt.Printf("id :%d,result:%c \n", id, <-chantemp)
			}
		}(chanlist[i], i)
	}
	for i := 0; i < len(chanlist); i++ {
		chanlist[i] <- 'a' + i
	}
	for i := 0; i < len(chanlist); i++ {
		chanlist[i] <- 'A' + i
	}
	time.Sleep(time.Second)
}
```

bufferchan 指定缓冲区大小

```go
func bufferchanel()  {
   chanbuf:=make(chan int ,4)
   go worker(chanbuf,9)
   for i:=0;i<4;i++{
      chanbuf<-'a'+i
   }
   time.Sleep(time.Second)
}
```

close chanel

```go
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
```

不要通过共享内存来通信，通过通信来共享内存 

通过通信来去掉上面的sleep

```go
type community struct {
   send chan int
   done chan int
}
func workerc( community community, id int) {
   for {
      for result := range community.send {
         fmt.Printf("id :%d,result:%c \n", id, result)
         community.done<-1
      }
   }
}
func goruntinec() {
   var chanlist [10]community
   for i := 0; i < len(chanlist); i++ {
      chanlist[i] =community{
         send:make(chan int),
         done:make(chan int),
      }
      go workerc(chanlist[i], i)
   }
   for i,com:=range chanlist{
      com.send <- 'a' + i
   }
   for _,com:=range chanlist{
      <-com.done
   }
   for i,com:=range  chanlist{
      com.send <- 'A' + i
   }
   for _,com:=range chanlist{
      <-com.done
   }
}
```











