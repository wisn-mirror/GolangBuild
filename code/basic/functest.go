package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func Adder() func(int )int  {
	var x int
	return func(detail int) int {
		x=x+detail
		return  x
	}
}

func main() {
	for i:=1;i<54;i++{
		for j:=1;j<54;j++{
			if i+j+i*j==54{
				fmt.Println(i,j,i+j)
			}
		}
	}
}

func funcName() {
	var add = Adder()
	fmt.Println(add(3))
	fmt.Println(add(4))
	fmt.Println(add(5))
	//3
	//7
	//12
}

func protectedRun( entry func() )  {
	defer func() {
		error:=recover()
		switch error.(type) {
		case runtime.Error:
			fmt.Println(error)
		default:
			fmt.Println("nil")
		}
	}()
	entry()
}
func PanicExcute()  {
	protectedRun(func() {
		fmt.Println("PanicExcute1")
		panic("A")
	})
	protectedRun(func() {
		fmt.Println("PanicExcute2")
		var b *int
		*b=2

	})
	fmt.Println("PanicExcute ")
}

func testTimeSum() {
	start := time.Now()
	// todo
	end := time.Now()
	timesum := end.Sub(start)
	fmt.Printf("%s\n", timesum)
}

func testpainc() {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	fmt.Println("panictest start")
	panic("panic")
	fmt.Println("panictest end")
	/*panictest start
	defer2
	defer1
	panic: panic	*/
}


type PassError struct {
	filename string
	line int
}

func (e *PassError)Error()string  {
	return fmt.Sprintf("%s:%d ",e.filename,e.line)
}
func newPassError(filename string,line int)error  {
	return &PassError{filename,line}
}

func customerError() {
	e := newPassError("hello.go", 2)
	fmt.Println(e)
}


func closeFuction()  {
	str:="helloworld"
	testfunc:=func(){
		str="hello"
	}
	fmt.Println(str)//helloworld
	testfunc()
	fmt.Println(str)//hello
}
//用于生成器
func playGen(name string)func()(string,int)  {
	hello:=222
	return func() (s string, i int) {
		return name,hello
	}
}
func testGen() {
	playGenerator := playGen("wisn")
	name, age := playGenerator()
	fmt.Println(name, age)
}

func testMannyArg() {
	//mannyArg1(3,6,4,0,3,2)
	//mannyArg2(3,"hello",3.3,)
	printtypeValue(3, "hello", 3.3, true)
	manyArg3(3, "hello", 3.3, true)
}
//可变参数的多重传递
func manyArg3(arg ...interface{})  {
	mannyArg2(arg...)
}
//内部实现就是一个切片
func mannyArg1(arg ...int ){
	for _,value:=range arg{
		fmt.Println(value)
	}
}
//任意参数
func mannyArg2(arg ...interface{} ){
	for _,value:=range arg{
		fmt.Println(value)
	}
}
//打印字段的类型
func printtypeValue(list ...interface{})  {
	for index,value:=range list{
		var typeStr string
		switch value.(type){
		case bool:
			typeStr="bool"
		case int:
			typeStr="int"
		case string:
			typeStr="string"
		}
		fmt.Println(index,value,typeStr)
	}
}




func testDefer() (result int ) {
	defer fmt.Println(" test defer 1")
	defer fmt.Println(" test defer 2")
	defer fmt.Println(" test defer 3")
	defer fmt.Println(" test defer 4")
	fmt.Println("excute")
	return 1
}
//excute
//test defer 4
//test defer 3
//test defer 2
//test defer 1

func getFileSize(filename string)int64  {
	f,error:=os.Open(filename)
	if error!=nil{
		return 0
	}
	fileinfo ,infoerror:=f.Stat()
	if infoerror!=nil{
		f.Close()
		return 0
	}
	size :=fileinfo.Size()
	f.Close()
	return size
}

func getFileSizeDefer(filename string)int64  {
	file,error:=os.Open(filename)
	if error!=nil{
		return 0
	}
	defer file.Close()
	info ,infoerror:=file.Stat()
	if infoerror!=nil{
		return 0
	}
	return info.Size()
}



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

func testFibonacci() {
	ch := make(chan int, 50)
	starttime := time.Now().UnixNano()
	for i := 0; i < 50; i++ {
		go ex(ch, i)
	}
	for i := 0; i < 50; i++ {
		fmt.Println(<-ch)
	}
	endtime := time.Now().UnixNano()
	timesum := (endtime - starttime) / 1e6
	fmt.Println("用时毫秒：", timesum)
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