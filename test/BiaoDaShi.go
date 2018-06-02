// BiaoDaShi
package test

import (
	"fmt"
	"strconv"
)

func IfTest() {
	a := 44
	//要用分号隔开
	if a, b := 33, 55; a > 2 {
		fmt.Println(a)
		fmt.Println(b)
	}
	fmt.Println(a)
}

/*
只有for循环  3种形式

*/
func ForTest() {
	fmt.Println("**********ForTest")
	//1  死循环
	//	for {

	//	}
	//第一种形式
	a := 1
	for {
		fmt.Println("for1 a:" + strconv.Itoa(a))
		a++
		if a > 5 {
			break
		}
	}
	fmt.Println("over 1")

	a = 0
	//第二种形式
	for a < 5 {
		a++
		fmt.Println("for 2 a:" + strconv.Itoa(a))
	}
	fmt.Println("over 2")

	b := "hellowisn"
	//第3种形式
	length := len(b)
	for i := 0; i < length; i++ {
		fmt.Println("for 3 a:" + strconv.Itoa(a))
	}

	fmt.Println("over 3")
	fmt.Println(length)
}

/**

不需要break  如果想继续执行可以添加fallthrough
*/
func SwitchTest() {
	fmt.Println("**********SwitchTest")

	//第1种形式
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
	default:
		fmt.Println("default")
	}
	fmt.Println(a)
}
