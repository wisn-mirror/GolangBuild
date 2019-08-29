package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	second=1
	minute=60*second
	hour=minute*60
	day=hour*24
)

func timeE(seconds int)(secondr int,  minuter int, hourr int) {
	return seconds,secondr/minute ,secondr/hour
}

func  feibona(n int) (result int ) {
	if n<=1{
		return 1
	}else{
		return feibona(n-1)+feibona(n-2)
	}
}

func  ex(ch chan int, i int)  {
	fmt.Println(i, feibona(i))
	ch<-i
}

func main() {
	ch:=make(chan int,50)
	starttime:=time.Now().UnixNano()
	for i:=0;i<50;i++{
		go ex(ch,i)
	}
	for i:=0;i<50;i++{
		fmt.Println(<-ch)
	}
	endtime:=time.Now().UnixNano()
	timesum:=(endtime-starttime)/1e6
	fmt.Println("用时毫秒：",timesum)
}

type testinterface interface {
	testfun(data interface{}) error
}
type TestI struct {
}

func (testinterface *TestI) testfun(data interface{}) error {
	fmt.Println("testfun",data)
	return nil
}
/**
接口测试
*/
func testInterface() {
	test := new(TestI)
	var testinter testinterface
	testinter = test
	testinter.testfun("str")
}

/**
函数封装
 */
func funcMap() {
	skill := map[string]func(string){
		"one": func(str string) {
			fmt.Println(str + "one")
		},
		"two": func(str string) {
			fmt.Println(str + "two")
		},
	}
	if result, ok := skill["one"]; ok {
		result("heheh")
	} else {
		fmt.Println("nil")
	}
}


func OnClickListener(list []string,fu func(string))  {
	for _,value:=range list{
		fu(value+"callback")
	}
}
//函数回调
func funcCallBack() {
	list := []string{
		"str2",
		"str3",
		"str4",
	}
	OnClickListener(list, func(result string) {
		fmt.Println(result)
	})
}
/*
匿名函数赋值
 */
func funcNoNameVar() {
	f := func(str string) string {
		return str + "1"
	}
	result := f("test")
	fmt.Println(result)
}
/*
函数链式调用
 */
func funcLinkExcute() {
	taskStr := []string{
		"gor testLinkFunc a ",
		" gor testLinkFunc B",
		"gor testLinkFunc c ",
		" gor testLinkFunc d",
	}
	funcList := []func(string) string{
		strings.TrimSpace,
		removeStr,
		strings.TrimSpace,
		strings.ToUpper,
	}
	excute(taskStr, funcList)
	for _, str := range taskStr {
		fmt.Println(str)
	}
}
func excute(list []string,listfunc []func(string)string)  {
	for index,_ :=range list{
		for _,funcStr:=range listfunc{
			list[index]=funcStr(list[index])
		}
	}
}
func removeStr(str string) string {
	return strings.TrimPrefix(str,"gor")
}

/**
函数变量
 */
func funcVar() {
	var fu func()
	fu = fire
	fu()
}
func fire()  {
	fmt.Println("test fire")
}


type Data struct {
	complex []int
	instance InstanceData
	pstr *InstanceData
}
type InstanceData struct {
	a int
}

func passByValue( data Data)  Data{
	return data
}

func passByValueP( data *Data)  Data{
	return *data
}

//测试值传递
func testValue() {
	in := Data{
		complex: []int{3, 2, 2},
		instance: InstanceData{
			44,
		},
		pstr: &InstanceData{3},
	}
	fmt.Printf("value  %+v\n", in)
	fmt.Printf("value %p \n", &in)
	//值传递
	result := passByValue(in)
	fmt.Printf("value0  %+v\n", result)
	fmt.Printf("value0 %p \n", &result)
	//值传递
	result1 := passByValueP(&in)
	result1.complex[0] = 33
	fmt.Printf("value1  %+v\n", result1)
	fmt.Printf("value1 %p \n", &result1)
	fmt.Printf("value  %+v\n", in)
	fmt.Printf("value %p \n", &in)
	fmt.Printf("value0  %+v\n", result)
	fmt.Printf("value0 %p \n", &result)
}