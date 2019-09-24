package main

import (
	"fmt"
	"time"
)

type testvalue int

func (test *testvalue) increment(step int) {
	*test++
}

func (test *testvalue) get() int {
	return int(*test)
}

/*
go run -race atomic_ex.go
有数据访问冲突
*/
func main() {
	var test testvalue
	test.increment(2)
	go func() {
		test.increment(3)
	}()
	time.Sleep(time.Second)
	fmt.Println(test.get())
}
