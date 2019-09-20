package temp

import (
	"fmt"
	"time"
)

func testgoroutine()  {
	var times int
	for{
		times++
		fmt.Println("time",times)
		time.Sleep(time.Second)
	}


}
func main() {
	go testgoroutine()
	var input string
	var result, _ = fmt.Scan(&input)
	fmt.Println(result)
}