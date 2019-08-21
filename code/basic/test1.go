package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
	 done:for j := 0; j < 10; j++ {
			fmt.Println(i,j)
			if i*j > 20 {
				continue done
			}
		}
	}
	fmt.Println("done")
}

func gotofun() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i*j > 200 {
				goto done
			}
		}
	}
done:
	fmt.Println("done")
	issucess := false
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i*j > 200 {
				issucess = true
				break
			}
		}
		if issucess {
			break
		}
	}
	fmt.Println("done")
}

func switchfun() {
	var t interface{}
	t = "string"
	switch t := t.(type) {
	default:
		fmt.Printf("type:%T", t)
	case int:
		fmt.Printf("int type:%T", t)
	case bool:
		fmt.Printf("bool type:%T", t)
	case string:
		fmt.Printf("string type:%T", t)
	case *int:
		fmt.Printf("*int type:%T", t)
	case *bool:
		fmt.Printf("*bool type:%T", t)
	}
	a := 3
	switch {
	case a > 2 && a < 30:
		fmt.Println("no")
		fallthrough
	case a < 15:
		fmt.Println("小于15")
	}
	result := "aa"
	switch result {
	//多种可能
	case "stR", "aa":
		fmt.Println("AA")
	case "str":
		fmt.Println("ok")
	case "Str":
		fmt.Println("BB")
	default:
		fmt.Println("CC")
	}
}

func rangeChanel() {
	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(time.Second)
		ch <- 2
		time.Sleep(time.Second)
		ch <- 3
		close(ch)
	}()
	for chresult := range ch {
		fmt.Println(chresult)
	}
}

func u99乘法表() {
	maptest:=map[string]int{
		"key1":100,
		"key2":200,
	}
	for key,value:=range maptest{
		fmt.Println(key,value)
	}
	//fmt.Println(maptest)
	var testStr="abc表"
	for key,value:=range testStr{
		fmt.Println(key,value)
	}
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%2d ", i, j, i*j)
		}
		fmt.Println()
	}
}

func fortest() {
	/*i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i:=0;i<10;i++{
		fmt.Println(i)
	}*/
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}
	/*for {
		fmt.Println("无限循环")
	}*/
}

func testifelse() {
	count := 3
	//变量在其实现了变量的功能后，作用范围越小，所造成的问题可能性越小
	if result := count / 2; result > 2 {
		fmt.Println(result)
	}
	if count > 2 {
		fmt.Println("success")
	} else if count < 5 {
		fmt.Println("AAA")
	} else {
		fmt.Println("BBB")
	}
}

func arraysInit() {
	//声明
	var array1 [2][2]int
	//初始化
	array2 := [2][2]int{{3, 2}, {1, 1}}
	array1 = array2
	fmt.Println(array1)
	//[[3 2] [1 1]]
	var array5 [3][3]int
	//编译错误
	//array5 = array1
	fmt.Println(array5)
	//指定索引初始化
	array3 := [2][2]int{1: {3, 2}}
	//指定索引内部的索引初始化
	array4 := [2][2]int{1: {1: 3}}
	fmt.Println(array1)
	//[[0 0] [0 0]]
	fmt.Println(array2)
	//[[3 2] [1 1]]
	fmt.Println(array3)
	//[[0 0] [3 2]]
	fmt.Println(array4)
	//[[0 0] [0 3]]
}

func compare() {
	a := [4]int{1, 3, 2, 5}
	b := [...]int{1, 3, 2, 5}
	c := [4]int{1, 3, 2, 6}
	fmt.Println(a == b, a == c, b == c)
	//true false false
	//d := [2]int{2, 3}
	//fmt.Println(a == d)
	//编译错误，只有大小相等的数据才能比较
}

func array() {
	name := "ddd"
	i, s := test3(&name)
	fmt.Println(i, s)
	array2 := []int{2, 3}
	//索引，元素
	for i, v := range array2 {
		fmt.Printf("%d %d\r\n", i, v)
	}
	fmt.Println(array2)
	array1 := []int{2, 3}
	fmt.Println(array1)
	fmt.Println(array1 == nil)
}
func test3(name *string) (int, string) {
	return 1, *name
}

func test2() {
	a, b := TestReturnValue()
	fmt.Println(a, b)
	c, _ := TestReturnValue()
	fmt.Println(c)
}
func TestReturnValue() (int, int) {
	return 10, 20
}

func test1() {
	name1, name2 := "evey", "wisn"
	fmt.Println(name1, name2)
	var a = 2
	var b = 3
	a, b = b, a
	fmt.Println(a, b)
}
