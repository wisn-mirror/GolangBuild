package main

import "fmt"

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

func main() {
	//fmt.Println(addTest(2,2))
	in:=Data{
		complex:[]int{3,2,2},
		instance:InstanceData{
			44,
		},
		pstr:&InstanceData{3},
	}
	fmt.Printf("value  %+v\n",in)
	fmt.Printf("value %p \n",&in)
	//值传递
	result:=passByValue(in)
	fmt.Printf("value0  %+v\n",result)
	fmt.Printf("value0 %p \n",&result)
	//值传递
	result1:=passByValueP(&in)
	result1.complex[0]=33
	fmt.Printf("value1  %+v\n",result1)
	fmt.Printf("value1 %p \n",&result1)

	fmt.Printf("value  %+v\n",in)
	fmt.Printf("value %p \n",&in)

	fmt.Printf("value0  %+v\n",result)
	fmt.Printf("value0 %p \n",&result)

	
}
func addTest(a,b int) int{
	return a+b
}