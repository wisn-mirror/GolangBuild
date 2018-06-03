package basic

import (
	"io/ioutil"
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func iftest()  {
	const filename  ="note.md"
	content,err:= ioutil.ReadFile(filename)
	if err!=nil{
		fmt.Println(err)
	} else{
		fmt.Printf("%s\n",content)
	}
	if content,err:=ioutil.ReadFile(filename);err==nil{
		fmt.Println(string(content))
	}else {
		fmt.Println("error:",err)
	}

}

func switchTest(a,b int,op string) int  {
	var result int
	switch op{
	case "+":
		result=a+b
	case "-":
		result=a-b
	case "*":
		result=a*b
	case "/":
		result=a/b
	default:
		//中断程序执行
		panic("not support operator:"+op)
	}
	return result
}


func switchTest1(source int) string  {
	var result string
	//可以没有表达式
	switch{
	case  source>100||source<10:
		result="ok"
	default:
		//中断程序执行
		panic("hahah")
	}
	return result
}

func forTest1()  {
	sum:=0
	for i:=1;i<=100;i++{
		sum+=i
	}
	fmt.Println(sum)
}

func convertToBin(value int)string {
	if value==0 {return "0"}
	result :=""
	for ;value>0;value=value/2{
		temp:=value%2
		result=strconv.Itoa(temp)+result
	}
	return result
}

func printFile() {
	file,err:=os.Open("note.md")
	if err!=nil{
		panic("error hahah ")
	}
	scanner :=bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}
func forever() {
	for{
		fmt.Println("forever hahah ")
	}
}

func main() {
	//iftest()
	result:=switchTest(9,4,"/")
	fmt.Println(result)
	fmt.Println(switchTest1(3),switchTest1(101))
	forTest1()

	fmt.Println(convertToBin(0))
	printFile()
	//forever();
}