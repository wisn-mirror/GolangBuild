package main

import "fmt"

func createMap() {
	m1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
		"key5": "value5",
	}
	fmt.Println(m1)
	m2 := make(map[string]int)
	fmt.Println(m2)
	var m3 map[string]int
	fmt.Println(m3)
	for  k,v:=range m1{
		fmt.Println(k,v)
	}
	delete(m1,"key1")
	for  k,_:=range m1{
		fmt.Println(k)
	}
	for  _,v:=range m1{
		fmt.Println(v)
	}
	valuetest:=m1["key3"]
	fmt.Println(valuetest)
	valuetest=m1["key3r"]
	fmt.Println(valuetest)
	valuetest1,ok:=m1["key3"]
	fmt.Println(valuetest1,ok)
	if valuetest2,ok2:=m1["key3"];ok2{
		fmt.Println(valuetest2)
	}else{
		fmt.Println("noteok ",ok2)
	}


}
func main() {
	createMap()
}
