package basic

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

func operation(op string, a, b int) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		//return a / b;
		result ,_:=div(a,b)
		return result;
	default:
		panic("not support operation:" + op)
	}
}


func operationNoPanic(op string, a, b int) (int ,error) {
	switch op {
	case "+":
		return a + b,nil
	case "-":
		return a - b,nil
	case "*":
		return a * b,nil
	case "/":
		//return a / b;
		result ,_:=div(a,b)
		return result,nil
	default:
		//panic("not support operation:" + op)
		return 0,fmt.Errorf("not support operation:%s",op)

	}
}
func div(a,b int )(q,r int)  {
/*	q=a/b;
	r=a%b;
	return*/
	return a/b,a%b
}

func apply(op  func(int ,int ) string , a ,b int ) string{
	p:=reflect.ValueOf(op).Pointer()
	opName:=runtime.FuncForPC(p).Name()
	fmt.Printf("call func:%s with args (%d,%d)\n",opName,a,b)
	return op(a,b)
}

//可变参数
func changeArgs(numbers ...int ) int  {
	result :=0
	for i:=range numbers{
		result=numbers[i]+result;
	}
	return result

}
func swap(a,b *int)  {
	*a,*b=*b,*a
}

func swap2(a,b int)(int ,int )  {
	return b,a
}
func main() {
	q,r:=div(6,2)
	fmt.Println(q,r)
	fmt.Println(operation("*",2,9))
	fmt.Println(operation("/",19,9))
	result ,err:=operationNoPanic("/",2,2)
	if err==nil{
		fmt.Println(result)
	}else{
		fmt.Println(err)
	}
	funresult:=apply(func(i int, i2 int) string {
		return strconv.Itoa(i)+":"+strconv.Itoa(i2)
	},3,5)
	fmt.Println(funresult)
	fmt.Println(changeArgs(3,2,2))
	a,b:=4,2
	swap(&a,&b)
	fmt.Println(a,b)
	a,b=swap2(a,b)
	fmt.Println(a,b)

 }
