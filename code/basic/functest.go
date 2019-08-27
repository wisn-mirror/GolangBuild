package main

import (
	"fmt"
	"strings"
)


func main() {
	taskStr:=[]string{
		"gor testLinkFunc a ",
		" gor testLinkFunc B",
		"gor testLinkFunc c ",
		" gor testLinkFunc d",
	}
	funcList:=[]func( string)string{
		removeStr,
		strings.TrimSpace,
		strings.ToUpper,
	}
	excute(taskStr,funcList)
	for _,str:=range taskStr{
		fmt.Println(str)
	}
}
func excute(list []string,listfunc []func(string)string)  {
	for index,value :=range list{
		for _,funcStr:=range listfunc{
			list[index]=funcStr(value)
		}
	}
}
func removeStr(str string) string {
	return str
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