package basic

import (
	"fmt"
	"math"
)

func variableInit1()  {
	var a int
	var b string
	fmt.Printf("%d  %q \n",a,b)

}

func variableInit2() {
	var a,b int=11,33
	var c ,d string="abc" ,"cde"
	fmt.Println(a,b,c,d)
}

func variableInit3() {
	var a,b="aaa",2
	fmt.Println(a,b)
}

func variableInit4()  {
	a,b,c,d :=3,2,true,"abc"
	fmt.Println(a,b,c,d)

}


func triangle() {
	var a,b=3,4
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}


const  filename="abc.txt"

func constsTest1() {
	const(
		filename1="abc.text2"
		d,e=1,2
	)
	const a, b  = 3,4;
	var c int ;
	c=int(math.Sqrt(a*a+b*b));
	fmt.Println(filename,c)
	fmt.Println(filename1,d,e)
}

func enumsTest() {
	const(
		java=0
		android=1
		python=2
		goland=3
		cpp=4
	)
	fmt.Println(java,android,python,goland,cpp)
}
func enumsIota() {
	const(
		java=iota
		android
		typton
		goland
		cpp
	)
	fmt.Println(java,android,typton,goland,cpp)
}

func iotaTest() {
	const(
		b=1<<(10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func main() {
	fmt.Println("Helloworld")
	variableInit1()
	variableInit2()
	variableInit3()
	variableInit4()
	triangle()
	constsTest1()
	enumsTest()
	enumsIota()
	iotaTest()
}
