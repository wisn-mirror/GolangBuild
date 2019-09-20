package temp

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		var times int
		for{
			times++
			fmt.Println("time",times)
			time.Sleep(time.Second)
		}
	}()
	var input string

	  result, error:= fmt.Scan(&input)
	  fmt.Println(result)
	  fmt.Println("result",input)
	  if error!=nil{
	  	fmt.Println(error)
	  }
}
