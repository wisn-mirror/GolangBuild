package main

import (
	"io/ioutil"
	"fmt"
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
		panic("not support operator:"+op)
	}
	return result
}


func main() {
	//iftest()
	result:=switchTest(9,4,"00")
	fmt.Println(result)

}