// Size
package test

import (
	"fmt"
)

//表示内存单位
const (
	B float64 = 1 << (iota * 10) //等于1
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func Msize() {
	fmt.Println("Msize")
	fmt.Println(Msize)
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(YB)
}
func Zhizheng() {
	ii := 33
	var p *int = &ii
	fmt.Println(p)
	fmt.Println(*p)
}
