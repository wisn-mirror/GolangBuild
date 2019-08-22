package main
import (
	"fmt"
	"time"
)

func main() {
}

func maneySlice() {
	slice := [][]int{{20}, {10, 20}}
	fmt.Println(slice)
	//[[20] [10 20]]
	slice[0] = append(slice[0], 33)
	fmt.Println(slice)
	//[[20 33] [10 20]]
}

func deleteSlice() {
	deleteSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(deleteSlice)
	fmt.Println(deleteSlice[2:])
	//删除前两个
	fmt.Println(deleteSlice[:len(deleteSlice)-2])
	//删除后两个
	index := 2
	fmt.Println(append(deleteSlice[:index], deleteSlice[index+2:]...))
	//删除以index开始两个元素
}

func copySlice() {
	count := 100
	resData := make([]int, count)
	for i := 0; i < count; i++ {
		resData[i] = i
	}
	refData := resData
	fmt.Printf("resData %p  refData %p\n", resData, refData)
	//指向的内存地址相同
	fmt.Println("resData ", resData)
	fmt.Println("refData", refData)
	copyData := make([]int, count)
	fmt.Println("copyData", copyData)
	copy(copyData, resData)
	fmt.Println("copyDatad", copyData)
	resData[0] = -9
	fmt.Println("refData", refData)
	fmt.Println("copyData", copyData)
	for j := 0; j < 3; j++ {
		copyData[j] = j + 100
	}
	fmt.Println("copyData", copyData)
	fmt.Println("copyData", copyData[1:3])
	//拷贝的copyData[1:3]，从0开始对应到resData中后续不动
	copy(resData, copyData[1:3])
	fmt.Println("refData", refData)
	fmt.Println("copyData", copyData)
}

func appendSlice() {
	/*appendtemp:=[]int{1,2}
	a:=[]int{3,4}
	fmt.Println(appendtemp[:1])
	fmt.Println(append([]int{5,6},appendtemp[:1]...))
	a=append(a[:1],append([]int{5,6},appendtemp[:1]...)...)
	fmt.Println(a)*/

	appendtemp:=[]int{2,4,6}
	var a []int
	for i:=0;i<10;i++{
		//添加到切片头部
		a=append(appendtemp,a...)
		fmt.Printf("%d,len:%d，cap:%d,%p\n",a,len(a),cap(a),a)
	}
	var slice []int
	fmt.Println(slice)
	//追加一个元素
	slice = append(slice, 1)
	fmt.Println(slice)
	//追加一个切片, 手动解包
	slice = append(slice, 2, 3, 4)
	fmt.Println(slice)
	//追加一个切片, 切片需要解包
	slice = append(slice, []int{3, 2}...)
	fmt.Println(slice)
}

func sliceInit() {
	var slicestrlist []string
	var sliceintlist []int
	var emptyslice = []int{}
	fmt.Println(slicestrlist, sliceintlist, emptyslice)
	//[] [] []
	fmt.Println(slicestrlist == nil, sliceintlist == nil, emptyslice == nil)
	//true true false
	a := make([] int, 3, 5)
	fmt.Println(a, len(a), cap(a))
	//[0 0 0] 3 5
}

func sliceTest1() {
/*	var a = [3]int{1, 2, 3}
	b:=a[0:0]
	fmt.Println("清空切片", b)*/
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])
	var test [30]int
	for i := 0; i < 30; i++ {
		test[i] = i + 1
	}
	fmt.Println(test)
	//[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30]
	//取区间
	fmt.Println(test[10:20])
	//[11 12 13 14 15 16 17 18 19 20]
	//开始到指定下标
	fmt.Println(test[:15])
	//[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
	//中间到结束
	fmt.Println(test[15:])
	//[16 17 18 19 20 21 22 23 24 25 26 27 28 29 30]
}

func continue测试() {
	for i := 0; i < 10; i++ {
	done:
		for j := 0; j < 10; j++ {
			fmt.Println(i, j)
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

func 分支结构() {
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

func array多维数组Init() {
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

func array比较() {
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

func init2() {
	a, b := TestReturnValue()
	fmt.Println(a, b)
	c, _ := TestReturnValue()
	fmt.Println(c)
}
func TestReturnValue() (int, int) {
	return 10, 20
}

func init1() {
	name1, name2 := "evey", "wisn"
	fmt.Println(name1, name2)
	var a = 2
	var b = 3
	a, b = b, a
	fmt.Println(a, b)
}
