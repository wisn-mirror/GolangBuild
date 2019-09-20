package temp

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	//testSelect1()
	resquest:=make(chan string,2)
	go server(resquest)
	response, error := client(resquest, "request str")
	if error==nil{
		fmt.Println("response:",response)
	}else{
		fmt.Println(error)
	}

}
func server(request chan string)  {
	for{
		requestStr:=<-request
		fmt.Println(" request:",requestStr)
		time.Sleep(time.Second*2)
		request<-" responseStr"
	}
}

func client(responsechan chan string, res string) (string, error) {
	responsechan <-res
	select {
	case response := <-responsechan:
		return response, nil
	case <-time.After(time.Second):
		return "", errors.New("time out")
	}
}

func testSelect1() {
	selectch1 := make(chan int, 1)
	for {
		select {
		case selectch1 <- 1:
		case selectch1 <- 2:

		}
		i := <-selectch1
		fmt.Println("result:", i)
	}
}
